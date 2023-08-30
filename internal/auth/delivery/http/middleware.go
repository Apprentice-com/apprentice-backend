package http

import (
	"net/http"
	"strings"

	"github.com/KadirbekSharau/apprentice-backend/internal/models"
	token "github.com/KadirbekSharau/apprentice-backend/pkg/auth"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(m token.TokenManager) gin.HandlerFunc {
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
		claims, err := m.ParseToken(ctx.Request.Context(), tokenString)
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

func GetUserIdFromContext(c *gin.Context) int {
    user, exists := c.Get("user")
    if !exists {
        return 0 // or any default value indicating no user
    }

    // Type assert user to the actual user type
    actualUser, ok := user.(*models.User) // Replace with your user type
    if !ok {
        return 0 // or any default value indicating invalid user type
    }

    // Access the ID field of the user
    userID := actualUser.ID

    return userID
}