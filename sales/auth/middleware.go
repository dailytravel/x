package auth

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Define a custom context key type
type contextKey string

// Define the key you want to use to store the gin.Context in the request context
const GinContextKey contextKey = "GinContextKey"
const AuthContextKey contextKey = "AuthContextKey"
const APIKeyContextKey contextKey = "APIKeyContextKey"
const LocaleContextKey contextKey = "LocaleContextKey"

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		ctx = context.WithValue(ctx, LocaleContextKey, c.GetHeader("x-locale"))
		ctx = context.WithValue(ctx, APIKeyContextKey, c.GetHeader("x-api-key"))

		// Decode the "auth" header into a map
		authHeader := c.GetHeader("auth")

		var authMap jwt.MapClaims
		if authHeader != "" {
			if err := json.Unmarshal([]byte(authHeader), &authMap); err != nil {
				fmt.Println("Error decoding auth header:", err)
			}
		}

		ctx = context.WithValue(ctx, AuthContextKey, authMap)

		// Add the gin.Context to the request context using the custom context key
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
