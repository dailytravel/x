const { ApolloGateway, RemoteGraphQLDataSource } = require("@apollo/gateway");
const { ApolloServer } = require("apollo-server-express");
const { expressjwt: jwt } = require("express-jwt");
const express = require("express");
const jwks = require("jwks-rsa");

const app = express();
const port = 4000; // Replace with your desired port number

// Define authentication middleware
const authenticate = jwt({
  secret: jwks.expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: "https://YOUR_AUTH0_DOMAIN/.well-known/jwks.json", // Replace with your Auth0 domain
  }),
  audience: "YOUR_AUTH0_AUDIENCE", // Replace with your Auth0 audience
  issuer: "https://YOUR_AUTH0_DOMAIN/", // Replace with your Auth0 domain
  algorithms: ["RS256"],
});

// Create a remote GraphQL data source
class AuthenticatedDataSource extends RemoteGraphQLDataSource {
  willSendRequest({ request, context }) {
    request.http.headers.set("issuer", context.issuer);
    request.http.headers.set("authorization", context.authorization);
    request.http.headers.set("client-id", context.clientId);
    request.http.headers.set("api-key", context.apiKey);
    request.http.headers.set(
      "auth",
      context.auth ? JSON.stringify(context.auth) : null
    );
  }
}

// Create an Apollo Gateway with IntrospectAndCompose
const gateway = new ApolloGateway({
  serviceList: [
    // Add your GraphQL services here
    { name: "4001", url: "http://localhost:4001/query" },
    { name: "4002", url: "http://localhost:4002/query" },
    { name: "4003", url: "http://localhost:4003/query" },
    { name: "4004", url: "http://localhost:4004/query" },
    { name: "4005", url: "http://localhost:4005/query" },
    { name: "4006", url: "http://localhost:4006/query" },
    { name: "4007", url: "http://localhost:4007/query" },
    { name: "4008", url: "http://localhost:4008/query" },
    { name: "4009", url: "http://localhost:4009/query" },
    { name: "4010", url: "http://localhost:4010/query" },
  ],
  buildService: ({ name, url }) => {
    return new AuthenticatedDataSource({ url });
  },
});

// Apply the authentication middleware to the Express app
app.use(authenticate);

(async () => {
  var server = new ApolloServer({
    gateway,
    subscriptions: false,
    introspection: true,
    playground: true,
    context: ({ req }) => ({
      req, // Make sure to include the req object in the context
    }),
  });

  await server.start();

  server.applyMiddleware({ app, cors: false });

  app.listen({ port }, () =>
    console.log(`Server ready at http://localhost:${port}${server.graphqlPath}`)
  );
})();
