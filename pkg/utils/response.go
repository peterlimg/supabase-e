package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse returns a success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse returns an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   errMsg,
	})
}

// BadRequestResponse returns a 400 Bad Request response
func BadRequestResponse(c *gin.Context, message string, err error) {
	ErrorResponse(c, http.StatusBadRequest, message, err)
}

// UnauthorizedResponse returns a 401 Unauthorized response
func UnauthorizedResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
}

// ForbiddenResponse returns a 403 Forbidden response
func ForbiddenResponse(c *gin.Context) {
	ErrorResponse(c, http.StatusForbidden, "Forbidden", nil)
}

// NotFoundResponse returns a 404 Not Found response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// InternalServerErrorResponse returns a 500 Internal Server Error response
func InternalServerErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err)
}
