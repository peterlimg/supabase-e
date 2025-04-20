package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/peterlimg/supabase-e/internal/models"
	"github.com/peterlimg/supabase-e/internal/services"
	"github.com/peterlimg/supabase-e/pkg/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	user, err := h.authService.Register(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user", err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User registered successfully", user)
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	resp, err := h.authService.Login(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Authentication failed", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login successful", resp)
}

// GetProfile handles getting the current user's profile
func (h *AuthHandler) GetProfile(c *gin.Context) {
	// Get the user ID from the context (set by the auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		utils.UnauthorizedResponse(c)
		return
	}

	user, err := h.authService.GetUserByID(userID.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get user profile", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile retrieved successfully", user)
}

// UpdateProfile handles updating the current user's profile
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	// Get the user ID from the context (set by the auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		utils.UnauthorizedResponse(c)
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	user, err := h.authService.UpdateUser(userID.(string), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile updated successfully", user)
}
