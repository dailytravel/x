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
    buildService({ url }) {
      return new RemoteGraphQLDataSource({
        url,
        willSendRequest({ request, context }) {
          Object.keys(context.headers).forEach((key) => {
            if (context.headers[key]) {
              request.http.headers.set(key, context.headers[key]);
            }
          });
        },
      });
    },
  }),
});

async function startApolloServer() {
  const server = new ApolloServer({
    gateway,
    subscriptions: false,
  });

  await server.start();

  app.use(
    cors({
      origin: "*", // Be sure to switch to your production domain
      exposedHeaders: true, // Expose all headers
      credentials: true,
    })
  );

  app.use(bodyParser.json());
  app.use(authenticate);
  app.use(
    expressMiddleware(server, {
      context: async ({ req }) => ({ req }),
    })
  );

  app.listen({ port }, () => {
    console.log(`🚀 Server ready at http://localhost:${port}/graphql`);
  });
}

startApolloServer().catch((err) => {
  console.error("Error starting the server:", err);
});
