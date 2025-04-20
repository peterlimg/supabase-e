package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// LoggerMiddleware logs request and response details
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get client IP
		clientIP := c.ClientIP()

		// Get status code
		statusCode := c.Writer.Status()

		// Get request method
		method := c.Request.Method

		// Construct the full path
		if raw != "" {
			path = path + "?" + raw
		}

		// Log the request
		logger := log.With().
			Str("method", method).
			Str("path", path).
			Int("status", statusCode).
			Str("ip", clientIP).
			Dur("latency", latency).
			Logger()

		switch {
		case statusCode >= 500:
			logger.Error().Msg("Server error")
		case statusCode >= 400:
			logger.Warn().Msg("Client error")
		default:
			logger.Info().Msg("Request processed")
		}
	}
}
