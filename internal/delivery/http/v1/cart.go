package v1

import (
	"github.com/gin-gonic/gin"
	dto "go1/internal/DTO"
	"net/http"
	"strconv"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {
	cart := api.Group("/cart")
	{
		cart.GET("/:cartId", h.getCart)
		cart.POST("/:cartId/items", h.addToCart)
		cart.DELETE("/:cartId/items/:productId", h.removeFromCart)
	}
}
func (h *Handler) getCart(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("cartId"))

	cart, err := h.services.Cart.GetCart(uint(cartID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *Handler) addToCart(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("cartId"))

	var input dto.AddToCartDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.services.Cart.AddToCart(uint(cartID), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
		"item":   item,
	})
}

func (h *Handler) removeFromCart(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("cartId"))
	productID, _ := strconv.Atoi(c.Param("productId"))

	if err := h.services.Cart.RemoveFromCart(uint(cartID), uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
