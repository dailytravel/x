package auth

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dailytravel/x/community/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
)

// Define a custom context key type
type contextKey string

// Define the key you want to use to store the gin.Context in the request context
const GinContextKey contextKey = "GinContextKey"
const AuthContextKey contextKey = "AuthContextKey"
const APIKeyContextKey contextKey = "APIKeyContextKey"
const LocaleContextKey contextKey = "LocaleContextKey"
const ClientIPContextKey contextKey = "ClientIPContextKey"
const UserAgentContextKey contextKey = "UserAgentContextKey"

func Middleware(client *grpc.ClientConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		ctx = context.WithValue(ctx, GinContextKey, c)
		ctx = context.WithValue(ctx, LocaleContextKey, c.GetHeader("x-locale"))
		ctx = context.WithValue(ctx, APIKeyContextKey, c.GetHeader("x-api-key"))
		ctx = context.WithValue(ctx, ClientIPContextKey, c.ClientIP())
		ctx = context.WithValue(ctx, UserAgentContextKey, c.GetHeader("user-agent"))

		authHeader := c.GetHeader("auth")
		if authHeader != "" {
			var authMap jwt.MapClaims
			if err := json.Unmarshal([]byte(authHeader), &authMap); err == nil {
				if jti, ok := authMap["jti"].(string); ok {
					if err == nil {
						if token, err := database.Redis.Get(ctx, jti).Result(); err == nil && token == "authenticated" {
							ctx = context.WithValue(ctx, AuthContextKey, authMap)
						} else {
							log.Println(err)
							c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
							return
						}
					}
				}
			}
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func Auth(ctx context.Context) jwt.MapClaims {
	raw, _ := ctx.Value(AuthContextKey).(jwt.MapClaims)
	return raw
}

func Locale(ctx context.Context) *string {
	raw, _ := ctx.Value(LocaleContextKey).(string)
	return &raw
}

func APIKey(ctx context.Context) *string {
	raw, _ := ctx.Value(APIKeyContextKey).(string) // Change this line
	return &raw
}

func ClientIP(ctx context.Context) *string {
	raw, _ := ctx.Value(ClientIPContextKey).(string) // Change this line
	return &raw
}

func UserAgent(ctx context.Context) *string {
	raw, _ := ctx.Value(UserAgentContextKey).(string) // Change this line
	return &raw
}
