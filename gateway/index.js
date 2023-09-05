import {
  ApolloGateway,
  IntrospectAndCompose,
  RemoteGraphQLDataSource,
} from "@apollo/gateway";
import { ApolloServer } from "@apollo/server";
import { expressMiddleware } from "@apollo/server/express4";
import { ApolloServerPluginDrainHttpServer } from "@apollo/server/plugin/drainHttpServer";
import { expressjwt } from "express-jwt";
import { fileURLToPath } from "url";
import { dirname } from "path";
import express from "express";
import http from "http";
import cors from "cors";
import bodyParser from "body-parser";
import jwks from "jwks-rsa";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const app = express();
const httpServer = http.createServer(app);

app.get("/.well-known/jwks.json", (req, res) => {
  res.sendFile("./.well-known/jwks.json", { root: __dirname });
});

const authenticate = expressjwt({
  secret: jwks.expressJwtSecret({
    cache: true,
    rateLimit: true,
    jwksRequestsPerMinute: 5,
    jwksUri: "http://localhost:4000/.well-known/jwks.json",
  }),
  audience: "https://api.trip.express/graphql",
  issuer: "https://api.trip.express",
  algorithms: ["RS256"],
  credentialsRequired: false,
});

class AuthenticatedDataSource extends RemoteGraphQLDataSource {
  willSendRequest({ request, context }) {
    for (const [key, value] of Object.entries(context)) {
      // Set the auth header
      if (key === "auth") {
        request.http.headers.set(key, JSON.stringify(value));
      }
      // Set headers that start with "x-" or are explicitly "x-api-key"
      else if (key.startsWith("x-")) {
        request.http.headers.set(key, value);
      }
    }
  }
}

const gateway = new ApolloGateway({
  buildService({ name, url }) {
    return new AuthenticatedDataSource({ url });
  },
  supergraphSdl: new IntrospectAndCompose({
    subgraphs: [
      { name: "account", url: "http://localhost:4001/query" },
      { name: "base", url: "http://localhost:4002/query" },
      { name: "cms", url: "http://localhost:4003/query" },
      { name: "community", url: "http://localhost:4004/query" },
      { name: "config", url: "http://localhost:4005/query" },
      { name: "finance", url: "http://localhost:4006/query" },
      { name: "hrm", url: "http://localhost:4007/query" },
      { name: "insight", url: "http://localhost:4008/query" },
      { name: "marketing", url: "http://localhost:4009/query" },
      { name: "sales", url: "http://localhost:4010/query" },
    ],
  }),
});

async function startApolloServer() {
  const server = new ApolloServer({
    gateway,
    uploads: false,
    plugins: [ApolloServerPluginDrainHttpServer({ httpServer })],
  });

  await server.start();

  app.use(
    cors({
      origin: [
        "http://localhost:3000",
        "https://api.trip.express",
        "https://app.trip.express",
        "https://trip.express",
        "https://www.trip.express",
      ],
    })
  );
  app.use(bodyParser.json({ limit: "50mb" }));
  app.use(bodyParser.urlencoded({ limit: "50mb", extended: true }));
  app.use(authenticate);
  app.use(
    expressMiddleware(server, {
      context: async ({ req }) => ({
        ...req.headers,
        auth: req.auth,
      }),
    })
  );

  // Modified server startup
  await new Promise((resolve) => httpServer.listen({ port: 4000 }, resolve));

  console.log(`🚀 Server ready at http://localhost:4000/`);
}

startApolloServer().catch((err) => {
  console.error("Error starting the server:", err);
});
