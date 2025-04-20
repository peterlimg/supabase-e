package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/peterlimg/supabase-e/config"
	"github.com/peterlimg/supabase-e/internal/handlers"
	"github.com/peterlimg/supabase-e/internal/repository"
	"github.com/peterlimg/supabase-e/internal/services"
	"github.com/peterlimg/supabase-e/pkg/database"
	"github.com/peterlimg/supabase-e/pkg/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}

	// Setup logger
	logger.Setup(cfg.LogLevel)
	logger := logger.GetLogger("main")
	logger.Info().Msg("Starting API server")

	// Initialize database connection
	db := database.NewSupabaseClient(cfg)
	logger.Info().Msg("Connected to Supabase")

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, db, cfg)
	productService := services.NewProductService(productRepo)

	// Setup router
	router := handlers.SetupRouter(cfg, db, authService, productService)

	// Create HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.Info().Int("port", cfg.Port).Msg("Server listening")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info().Msg("Shutting down server...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	logger.Info().Msg("Server exited gracefully")
}
