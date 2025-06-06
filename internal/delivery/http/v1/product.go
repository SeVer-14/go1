package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go1/internal/entity"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type (
	createProductInput struct {
		ProductID int     `json:"ProductID" binding:"required"`
		Title     string  `json:"title" binding:"required"`
		Price     float64 `json:"price" binding:"required"`
	}

	updateProductInput struct {
		ProductID int     `json:"ProductID" binding:"required"`
		Title     string  `json:"title" binding:"required"`
		Price     float64 `json:"price" binding:"required"`
	}
)

func (h *Handler) initProductRoutes(api *gin.RouterGroup) {
	products := api.Group("/products")
	{
		products.GET("/", h.getAllProducts)
		products.POST("/", h.createProduct)
		products.PUT("/:id", h.updateProduct)
		products.DELETE("/:id", h.deleteProduct)
	}
}
func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.services.Product.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) createProduct(c *gin.Context) {
	var input createProductInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := entity.Product{
		ProductID: input.ProductID,
		Title:     input.Title,
		Price:     input.Price,
	}

	createdProduct, err := h.services.Product.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}
func (h *Handler) updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var input updateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := entity.Product{
		ProductID: input.ProductID,
		Title:     input.Title,
		Price:     input.Price,
	}

	updatedProduct, err := h.services.Product.UpdateProduct(uint(id), product)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (h *Handler) deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	if err := h.services.Product.DeleteProduct(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}
