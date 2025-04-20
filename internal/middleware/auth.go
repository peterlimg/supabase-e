package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/peterlimg/supabase-e/config"
	"github.com/peterlimg/supabase-e/pkg/utils"
)

// AuthMiddleware creates a middleware for JWT authentication
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(c)
			c.Abort()
			return
		}

		// Check if the Authorization header is in the correct format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.UnauthorizedResponse(c)
			c.Abort()
			return
		}

		// Validate the token
		tokenString := parts[1]
		claims, err := utils.ValidateJWT(tokenString, cfg.JWTSecret)
		if err != nil {
			utils.UnauthorizedResponse(c)
			c.Abort()
			return
		}

		// Set the user ID and role in the context
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// RoleMiddleware creates a middleware for role-based authorization
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user role from the context
		role, exists := c.Get("role")
		if !exists {
			utils.UnauthorizedResponse(c)
			c.Abort()
			return
		}

		// Check if the user has one of the required roles
		userRole := role.(string)
		for _, r := range roles {
			if r == userRole {
				c.Next()
				return
			}
		}

		// If the user doesn't have any of the required roles, return a forbidden response
		utils.ForbiddenResponse(c)
		c.Abort()
	}
}
