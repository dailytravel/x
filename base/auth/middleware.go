package auth

import (
	"context"

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
		auth := c.Request.Header.Get("auth")
		apiKey := c.Request.Header.Get("x-api-key")
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		ctx = context.WithValue(ctx, AuthContextKey, auth)
		ctx = context.WithValue(ctx, APIKeyContextKey, apiKey)

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
	raw, _ := ctx.Value(LocaleContextKey).(*string)
	return raw
}

func APIKey(ctx context.Context) *string {
	raw, _ := ctx.Value(APIKeyContextKey).(*string)
	return raw
}
