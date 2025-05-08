package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gabrielagui373/obiwanapp-api/internal/models"
	"github.com/gabrielagui373/obiwanapp-api/internal/services"
	"github.com/gin-gonic/gin"
)

// verifies JWT token and sets user in context
func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "bearer token required"})
			return
		}

		token, err := authService.ValidateJWT(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}

		// Check if token is access token
		if claims["type"] != "access" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token type"})
			return
		}

		//Get user from DB
		var user models.User
		userID := claims["user_id"].(string) // Certifique-se de que o claim é do tipo string
		if err := authService.DB.Preload("Roles.Permissions").First(&user, "id = ?", userID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			return
		}

		// Set user in context
		ctx.Set("user", &user)
		ctx.Next()
	}
}

func RBACMiddleware(permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exists := ctx.Get("user")
		if !exists {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		authenticatedUser := user.(models.User)
		hasPermission := false

		for _, role := range authenticatedUser.Roles {
			for _, perm := range role.Permissions {
				if perm.Name == permission {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}

		if !hasPermission {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
			return
		}

		ctx.Next()
	}
}
