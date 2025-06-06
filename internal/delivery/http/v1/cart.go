package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initCartRoutes(api *gin.RouterGroup) {
	cart := api.Group("/cart")
	{
		cart.GET("/:user_id", h.getCart)
		cart.POST("/:user_id/items", h.addToCart)
		cart.DELETE("/:user_id/items/:product_id", h.removeFromCart)
	}
}
func (h *Handler) getCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))

	cart, err := h.services.Cart.GetCart(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cart)
}

func (h *Handler) addToCart(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))

	var input struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"min=1"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.services.Cart.AddToCart(uint(userID), input.ProductID, input.Quantity)
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
	userID, _ := strconv.Atoi(c.Param("user_id"))
	productID, _ := strconv.Atoi(c.Param("product_id"))

	if err := h.services.Cart.RemoveFromCart(uint(userID), uint(productID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
