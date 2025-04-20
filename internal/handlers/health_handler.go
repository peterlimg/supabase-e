package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peterlimg/supabase-e/pkg/database"
	"github.com/peterlimg/supabase-e/pkg/utils"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	db *database.Client
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(db *database.Client) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

// Check handles checking the health of the API
func (h *HealthHandler) Check(c *gin.Context) {
	// Check database connection
	err := h.db.Health()
	if err != nil {
		utils.ErrorResponse(c, http.StatusServiceUnavailable, "Database connection failed", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "API is healthy", map[string]string{
		"status": "up",
	})
}
