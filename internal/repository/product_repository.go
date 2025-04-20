package repository

import (
	"fmt"

	"github.com/peterlimg/supabase-e/internal/models"
	"github.com/peterlimg/supabase-e/pkg/database"
)

// ProductRepository handles product data operations
type ProductRepository struct {
	db *database.Client
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *database.Client) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Create creates a new product
func (r *ProductRepository) Create(product models.Product) (*models.Product, error) {
	var result []models.Product
	err := r.db.ServiceClient.DB.From("products").Insert(product).Execute(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no product returned after insert")
	}

	return &result[0], nil
}

// GetByID retrieves a product by ID
func (r *ProductRepository) GetByID(id string) (*models.Product, error) {
	var products []models.Product
	err := r.db.ServiceClient.DB.From("products").Select("*").Eq("id", id).Execute(&products)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return &products[0], nil
}

// Update updates a product
func (r *ProductRepository) Update(id string, product models.UpdateProductRequest) (*models.Product, error) {
	var result []models.Product
	err := r.db.ServiceClient.DB.From("products").Update(product).Eq("id", id).Execute(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("product not found")
	}

	return &result[0], nil
}

// Delete deletes a product
func (r *ProductRepository) Delete(id string) error {
	err := r.db.ServiceClient.DB.From("products").Delete().Eq("id", id).Execute(nil)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// List lists all products with pagination and optional filtering
func (r *ProductRepository) List(page, pageSize int, category string) ([]models.Product, error) {
	var products []models.Product
	
	// Calculate offset
	offset := (page - 1) * pageSize
	limit := pageSize
	
	// Build query
	query := r.db.ServiceClient.DB.From("products").Select("*")
	
	// Add category filter if provided
	if category != "" {
		// Need to handle the filter differently
		filterQuery := query.Eq("category", category)
		err := filterQuery.Execute(&products)
		if err != nil {
			return nil, fmt.Errorf("failed to filter products: %w", err)
		}
	} else {
		// If no filter, just execute the base query with limit
		query = query.Limit(limit)
		err := query.Execute(&products)
		if err != nil {
			return nil, fmt.Errorf("failed to list products: %w", err)
		}
	}
	
	// Apply offset and limit manually
	if len(products) > offset {
		end := offset + limit
		if end > len(products) {
			end = len(products)
		}
		products = products[offset:end]
	} else {
		products = []models.Product{}
	}
	
	// Sort manually by created_at in descending order
	// Note: In a real application, you might want to use a proper sorting function
	
	return products, nil
}

// GetProductWithUser retrieves a product with its creator's information
func (r *ProductRepository) GetProductWithUser(id string) (*models.ProductResponse, error) {
	// First get the product
	product, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	// Then get the user who created it
	var users []models.User
	err = r.db.ServiceClient.DB.From("users").Select("*").Eq("id", product.CreatedBy).Execute(&users)
	if err != nil || len(users) == 0 {
		// If we can't get the user, just return the product without user info
		return &models.ProductResponse{Product: *product}, nil
	}
	
	// Return both
	return &models.ProductResponse{
		Product:       *product,
		CreatedByUser: users[0],
	}, nil
}
