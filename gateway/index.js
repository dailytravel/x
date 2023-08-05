const express = require("express");
const {
  ApolloGateway,
  IntrospectAndCompose,
  RemoteGraphQLDataSource,
} = require("@apollo/gateway");
const { ApolloServer } = require("@apollo/server");
const { expressMiddleware } = require("@apollo/server/express4");
const { expressjwt } = require("express-jwt");
const { json } = require("body-parser");
const jwks = require("jwks-rsa");
const cors = require("cors");

const app = express();
const port = 4000;

app.get("/.well-known/jwks.json", (req, res) => {
  res.sendFile("./.well-known/jwks.json", { root: __dirname });
});

const authenticate = expressjwt({
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
  willSendRequest({ request, context }) {
    request.http.headers.set("auth", context.auth);
    request.http.headers.set("x-api-key", context.apiKey);
  }
}

const gateway = new ApolloGateway({
  buildService({ name, url }) {
    return new AuthenticatedDataSource({ url });
  },
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
  }),
});

async function startApolloServer() {
  const server = new ApolloServer({
    gateway,
  });

  await server.start();

  // Use the Apollo server middleware along with authentication and CORS
  app.use(
    cors({
      origin: ["http://localhost:4000", "https://api.trip.express"],
    }),
    json(),
    authenticate,
    expressMiddleware(server, {
      context: async ({ req }) => {
        return {
          apiKey: req.headers["x-api-key"],
          auth: req.headers["auth"],
        };
      },
    })
  );

  // Start the server
  app.listen({ port }, () => {
    console.log(`🚀 Server ready at http://localhost:${port}/graphql`);
  });
}

startApolloServer().catch((err) => {
  console.error("Error starting the server:", err);
});
