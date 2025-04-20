package models

import (
	"time"

	"github.com/google/uuid"
)

// Product represents a product in the system
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateProductRequest represents the request to create a new product
type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Category    string  `json:"category" binding:"required"`
	ImageURL    string  `json:"image_url,omitempty"`
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty" binding:"omitempty,gt=0"`
	Category    string  `json:"category,omitempty"`
	ImageURL    string  `json:"image_url,omitempty"`
}

// ProductResponse represents a product response with additional data
type ProductResponse struct {
	Product
	CreatedByUser User `json:"created_by_user,omitempty"`
}

// NewProduct creates a new product with default values
func NewProduct(req CreateProductRequest, userID string) Product {
	now := time.Now()
	return Product{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
		CreatedBy:   userID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
