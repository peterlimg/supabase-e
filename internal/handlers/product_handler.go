package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/peterlimg/supabase-e/internal/models"
	"github.com/peterlimg/supabase-e/internal/services"
	"github.com/peterlimg/supabase-e/pkg/utils"
)

// ProductHandler handles product requests
type ProductHandler struct {
	productService *services.ProductService
}

// NewProductHandler creates a new product handler
func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

// CreateProduct handles creating a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	// Get the user ID from the context (set by the auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		utils.UnauthorizedResponse(c)
		return
	}

	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	product, err := h.productService.CreateProduct(req, userID.(string))
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Product created successfully", product)
}

// GetProduct handles getting a product by ID
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequestResponse(c, "Product ID is required", nil)
		return
	}

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		utils.NotFoundResponse(c, "Product not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}

// GetProductWithUser handles getting a product with its creator's information
func (h *ProductHandler) GetProductWithUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequestResponse(c, "Product ID is required", nil)
		return
	}

	product, err := h.productService.GetProductWithUser(id)
	if err != nil {
		utils.NotFoundResponse(c, "Product not found")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}

// UpdateProduct handles updating a product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequestResponse(c, "Product ID is required", nil)
		return
	}

	var req models.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request body", err)
		return
	}

	product, err := h.productService.UpdateProduct(id, req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update product", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Product updated successfully", product)
}

// DeleteProduct handles deleting a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequestResponse(c, "Product ID is required", nil)
		return
	}

	err := h.productService.DeleteProduct(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete product", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Product deleted successfully", nil)
}

// ListProducts handles listing all products with pagination and optional filtering
func (h *ProductHandler) ListProducts(c *gin.Context) {
	// Parse pagination parameters
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	category := c.Query("category")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	products, err := h.productService.ListProducts(page, pageSize, category)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to list products", err)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Products retrieved successfully", products)
}
