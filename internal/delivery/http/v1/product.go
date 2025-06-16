package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	dto "go1/internal/DTO"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

// @Summary Получить все продукты
// @Description Возвращает список всех продуктов
// @Tags products
// @Produce json
// @Success 200 {array} dto.ProductDTO "Успешный ответ"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/products [get]
func (h *Handler) getAllProducts(c *gin.Context) {
	products, err := h.services.Product.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// @Summary Создать продукт
// @Description Создаёт новый продукт
// @Tags products
// @Accept json
// @Produce json
// @Param input body dto.ProductDTO true "Данные продукта"
// @Success 201 {object} dto.ProductDTO "Продукт создан"
// @Failure 400 {object} map[string]string "Неверный запрос"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/products [post]
func (h *Handler) createProduct(c *gin.Context) {
	var input dto.ProductDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.services.Product.CreateProduct(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// @Summary Обновить продукт
// @Description Обновляет существующий продукт
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Param input body dto.ProductDTO true "Новые данные продукта"
// @Success 200 {object} dto.ProductDTO "Продукт обновлён"
// @Failure 400 {object} map[string]string "Неверный ID или данные"
// @Failure 404 {object} map[string]string "Продукт не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/products/{id} [put]
func (h *Handler) updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var input dto.ProductDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.services.Product.UpdateProduct(uint(id), input)
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

// @Summary Удалить продукт
// @Description Удаляет продукт по ID
// @Tags products
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]string "Продукт удалён"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 404 {object} map[string]string "Продукт не найден"
// @Failure 500 {object} map[string]string "Ошибка сервера"
// @Router /api/products/{id} [delete]
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
