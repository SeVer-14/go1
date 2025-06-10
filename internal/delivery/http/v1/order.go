package v1

import (
	"github.com/gin-gonic/gin"
	dto "go1/internal/DTO"
	"net/http"
	"strconv"
)

func (h *Handler) initOrderRoutes(api *gin.RouterGroup) {
	orders := api.Group("/orders")
	{
		orders.POST("/:cartId", h.createOrder)
		orders.GET("/:cartId", h.getOrders)
		orders.PUT("/status", h.updateOrderStatus)
	}
}

func (h *Handler) createOrder(c *gin.Context) {
	cartID, _ := strconv.Atoi(c.Param("cartId"))

	cart, err := h.services.Cart.GetCart(uint(cartID))
	if err != nil || len(cart.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty or invalid cart"})
		return
	}

	order, err := h.services.Order.CreateOrder(cart.CartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

func (h *Handler) getOrders(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cart id"})
		return
	}

	orders, err := h.services.Order.GetOrders(uint(cartID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) updateOrderStatus(c *gin.Context) {

	var input dto.UpdateOrderStatusDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Order.UpdateOrderStatus(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
