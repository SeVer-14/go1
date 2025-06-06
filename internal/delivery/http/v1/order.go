package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) initOrderRoutes(api *gin.RouterGroup) {
	orders := api.Group("/orders")
	{
		orders.POST("/:user_id", h.createOrder)
		orders.GET("/:user_id", h.getOrders)
		orders.PUT("/:order_id/status", h.updateOrderStatus)
	}
}
func (h *Handler) createOrder(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("user_id"))

	cart, err := h.services.Cart.GetCart(uint(userID))
	if err != nil || len(cart.CartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty or invalid cart"})
		return
	}

	order, err := h.services.Order.CreateOrder(cart)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

func (h *Handler) getOrders(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	orders, err := h.services.Order.GetOrders(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) updateOrderStatus(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid order id"})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required,oneof=pending processing completed cancelled"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.services.Order.UpdateOrderStatus(uint(orderID), input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
