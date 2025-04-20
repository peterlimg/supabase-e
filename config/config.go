package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Config holds all configuration for the application
type Config struct {
	Port               int
	Environment        string
	LogLevel           string
	SupabaseURL        string
	SupabaseKey        string
	SupabaseServiceKey string
	JWTSecret          string
	JWTExpiry          time.Duration
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("No .env file found, using environment variables")
	}

	// Default values
	port := 8080
	env := "development"
	logLevel := "info"
	jwtExpiry := 24 * time.Hour

	// Parse port
	if os.Getenv("PORT") != "" {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err == nil {
			port = p
		}
	}

	// Parse environment
	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	}

	// Parse log level
	if os.Getenv("LOG_LEVEL") != "" {
		logLevel = os.Getenv("LOG_LEVEL")
	}

	// Parse JWT expiry
	if os.Getenv("JWT_EXPIRY") != "" {
		duration, err := time.ParseDuration(os.Getenv("JWT_EXPIRY"))
		if err == nil {
			jwtExpiry = duration
		}
	}

	// Required values
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabaseServiceKey := os.Getenv("SUPABASE_SERVICE_KEY")
	jwtSecret := os.Getenv("JWT_SECRET")

	// Validate required values
	if supabaseURL == "" || supabaseKey == "" || supabaseServiceKey == "" || jwtSecret == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &Config{
		Port:               port,
		Environment:        env,
		LogLevel:           logLevel,
		SupabaseURL:        supabaseURL,
		SupabaseKey:        supabaseKey,
		SupabaseServiceKey: supabaseServiceKey,
		JWTSecret:          jwtSecret,
		JWTExpiry:          jwtExpiry,
	}, nil
}
