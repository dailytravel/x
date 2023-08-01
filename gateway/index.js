const { ApolloGateway, RemoteGraphQLDataSource } = require("@apollo/gateway");
const { ApolloServer } = require("@apollo/server");
const express = require("express");
const jwt = require("express-jwt");
const jwks = require("jwks-rsa");

const app = express();

// Define authentication middleware
const authenticate = jwt({
  secret: jwks.expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: "http://localhost:4000/.well-known/jwks.json", // Replace with your Auth0 domain
  }),
  audience: "YOUR_AUTH0_AUDIENCE", // Replace with your Auth0 audience
  issuer: "https://YOUR_AUTH0_DOMAIN/", // Replace with your Auth0 domain
  algorithms: ["RS256"],
});

// Create a remote GraphQL data source
class MyRemoteDataSource extends RemoteGraphQLDataSource {
  willSendRequest({ request, context }) {
    if (context.authorization) {
      request.http.headers.set("Authorization", context.authorization);
    }
  }
}

// Create an Apollo Gateway
const gateway = new ApolloGateway({
  serviceList: [
    // Add your GraphQL services here
    { name: "account", url: "http://localhost:4001/query" },
  ],
  buildService({ url }) {
    return new MyRemoteDataSource({ url });
  },
});

// Create an Apollo Server with gateway
const server = new ApolloServer({
  gateway,
  subscriptions: false,
  context: ({ req }) => ({
    authorization: req.headers.authorization,
  }),
});

// Apply the authentication middleware to the Express app
app.use(authenticate);

// Apply the Apollo middleware to the Express app
server.applyMiddleware({ app });

// Start the server
const PORT = process.env.PORT || 4000;
app.listen(PORT, () => {
  console.log(`Server listening on port ${PORT}`);
});
