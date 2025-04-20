package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/peterlimg/supabase-e/config"
	"github.com/peterlimg/supabase-e/internal/middleware"
	"github.com/peterlimg/supabase-e/internal/services"
	"github.com/peterlimg/supabase-e/pkg/database"
)

// SetupRouter sets up the API routes
func SetupRouter(
	cfg *config.Config,
	db *database.Client,
	authService *services.AuthService,
	productService *services.ProductService,
) *gin.Engine {
	// Create a new Gin router
	r := gin.New()

	// Use the logger and recovery middleware
	r.Use(middleware.LoggerMiddleware())
	r.Use(gin.Recovery())

	// Create handlers
	authHandler := NewAuthHandler(authService)
	productHandler := NewProductHandler(productService)
	healthHandler := NewHealthHandler(db)

	// Health check route
	r.GET("/health", healthHandler.Check)

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg))
		{
			// User routes
			user := protected.Group("/users")
			{
				user.GET("/me", authHandler.GetProfile)
				user.PUT("/me", authHandler.UpdateProfile)
			}

			// Product routes
			products := protected.Group("/products")
			{
				products.POST("", productHandler.CreateProduct)
				products.GET("", productHandler.ListProducts)
				products.GET("/:id", productHandler.GetProduct)
				products.GET("/:id/with-user", productHandler.GetProductWithUser)
				products.PUT("/:id", productHandler.UpdateProduct)
				products.DELETE("/:id", productHandler.DeleteProduct)
			}
		}
	}

	return r
}
