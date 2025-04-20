package services

import (
	"fmt"

	"github.com/peterlimg/supabase-e/internal/models"
	"github.com/peterlimg/supabase-e/internal/repository"
)

// ProductService handles product operations
type ProductService struct {
	productRepo *repository.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(req models.CreateProductRequest, userID string) (*models.Product, error) {
	// Create a new product model
	product := models.NewProduct(req, userID)

	// Save to the database
	result, err := s.productRepo.Create(product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return result, nil
}

// GetProductByID gets a product by ID
func (s *ProductService) GetProductByID(id string) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}

// GetProductWithUser gets a product with its creator's information
func (s *ProductService) GetProductWithUser(id string) (*models.ProductResponse, error) {
	return s.productRepo.GetProductWithUser(id)
}

// UpdateProduct updates a product
func (s *ProductService) UpdateProduct(id string, req models.UpdateProductRequest) (*models.Product, error) {
	return s.productRepo.Update(id, req)
}

// DeleteProduct deletes a product
func (s *ProductService) DeleteProduct(id string) error {
	return s.productRepo.Delete(id)
}

// ListProducts lists all products with pagination and optional filtering
func (s *ProductService) ListProducts(page, pageSize int, category string) ([]models.Product, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	return s.productRepo.List(page, pageSize, category)
}
