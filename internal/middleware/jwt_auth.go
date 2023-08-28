package middleware

import (
	"net/http"
	"strings"

	"github.com/KadirbekSharau/apprentice-backend/internal/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(a auth.UseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the JWT token from the Authorization header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			ctx.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Parse and verify the JWT token
		claims, err := a.ParseToken(ctx.Request.Context(), tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// Attach the claims to the context for use in handlers
		ctx.Set("user", claims)
		ctx.Next()
	}
}