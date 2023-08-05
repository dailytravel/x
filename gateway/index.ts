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

app.get("/.well-known/jwks.json", (req, res) => {
  res.sendFile("./.well-known/jwks.json", { root: __dirname });
});

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

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
  constructor(url) {
    super({ url });
  }

  willSendRequest({ request, context }) {
    const { req } = context;
    request.http.headers.set("x-api-key", req.headers["x-api-key"]);
    request.http.headers.set(
      "auth",
      req.headers.auth ? JSON.stringify(req.headers.auth) : null
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
    buildService({ name, url }) {
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
      req,
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
