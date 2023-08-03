const express = require("express");
const {
  ApolloGateway,
  IntrospectAndCompose,
  RemoteGraphQLDataSource,
} = require("@apollo/gateway");
const { ApolloServer } = require("@apollo/server");
const { expressMiddleware } = require("@apollo/server/express4");
const { expressjwt: jwt } = require("express-jwt");
const jwks = require("jwks-rsa");
const cors = require("cors");
const bodyParser = require("body-parser");

const app = express();
const port = 4000; // Replace with your desired port number

// Define authentication middleware
const authenticate = jwt({
  secret: jwks.expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: "http://localhost:4000/.well-known/jwks.json", // Replace with your Auth0 domain
  }),
  audience: "api.trip.express",
  issuer: "https://api.trip.express/",
  algorithms: ["RS256"],
  credentialsRequired: false,
});

// Create a remote GraphQL data source
class AuthenticatedDataSource extends RemoteGraphQLDataSource {
  constructor({ url }) {
    super();
    this.url = url;
  }

  willSendRequest({ request, context }) {
    request.http.headers.set("x-api-key", context.apiKey);
    request.http.headers.set(
      "auth",
      context.auth ? JSON.stringify(context.auth) : null
    );
  }
}

// Create an Apollo Gateway with IntrospectAndCompose
const gateway = new ApolloGateway({
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: "account", url: "http://localhost:4001/query" },
      { name: "cms", url: "http://localhost:4002/query" },
      { name: "config", url: "http://localhost:4003/query" },
      { name: "finance", url: "http://localhost:4004/query" },
      { name: "hrm", url: "http://localhost:4005/query" },
      { name: "marketing", url: "http://localhost:4006/query" },
      { name: "reporting", url: "http://localhost:4007/query" },
      { name: "sales", url: "http://localhost:4008/query" },
      { name: "search", url: "http://localhost:4009/query" },
      { name: "service", url: "http://localhost:4010/query" },
    ],
    introspectionHeaders: ({ context }) => {
      return {
        Authorization: context?.authorization || "",
        "x-api-key": context?.apiKey || "",
      };
    },
    buildService: ({ name, url }) => {
      return new AuthenticatedDataSource({ url });
    },
  }),
});

async function startApolloServer() {
  const server = new ApolloServer({
    gateway,
    subscriptions: false,
    introspection: true,
    playground: true,
    context: ({ req }) => ({
      req, // Make sure to include the req object in the context
    }),
  });

  await server.start();

  app.use(bodyParser.json());
  app.use(authenticate);

  app.use(
    "/graphql",
    expressMiddleware(server),
    cors({
      origin: "*",
      credentials: true,
    })
  );

  app.listen({ port }, () => {
    console.log(`🚀 Server ready at http://localhost:${port}/graphql`);
  });
}

startApolloServer().catch((err) => {
  console.error("Error starting the server:", err);
});
