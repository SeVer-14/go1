package controller

import (
	"github.com/gin-gonic/gin"
	"go1/entity"
	"go1/service"
	"strconv"
)

type ProductController interface {
	Show() []entity.Product
	Add(ctx *gin.Context) entity.Product
	Delete(ctx *gin.Context)
	AddToCart(ctx *gin.Context)
	GetCart(ctx *gin.Context)
	RemoveFromCart(ctx *gin.Context)
	UpdateCartItem(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	GetOrders(ctx *gin.Context)
	UpdateOrderStatus(ctx *gin.Context)
}

type controller struct {
	service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &controller{service: service}
}
func (c *controller) Show() []entity.Product {
	return c.service.Show()
}
func (c *controller) Add(ctx *gin.Context) entity.Product {
	var product entity.Product
	ctx.BindJSON(&product)
	c.service.Add(product)
	return product
}

func (c *controller) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	success := c.service.Delete(uint(id))

	if success {
		ctx.JSON(200, gin.H{"message": "Product deleted successfully"})
	} else {
		ctx.JSON(404, gin.H{"error": "Product not found"})
	}
}

func (c *controller) AddToCart(ctx *gin.Context) {
	var request struct {
		ProductID uint `json:"product_id"`
		UserID    uint `json:"user_id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.AddToCart(request.ProductID, request.UserID); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Product added to cart"})
}

func (c *controller) GetCart(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	cart, err := c.service.GetCart(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, cart)
}

func (c *controller) RemoveFromCart(ctx *gin.Context) {
	cartIDStr := ctx.Param("cart_id")
	cartID, err := strconv.ParseUint(cartIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid cart ID"})
		return
	}

	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.service.RemoveFromCart(uint(cartID), uint(userID)); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Item removed from cart"})
}

func (c *controller) UpdateCartItem(ctx *gin.Context) {
	cartIDStr := ctx.Param("cart_id")
	cartID, err := strconv.ParseUint(cartIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid cart ID"})
		return
	}

	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		Quantity int `json:"quantity"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateCartItem(uint(cartID), uint(userID), request.Quantity); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Cart item updated"})
}

func (c *controller) CreateOrder(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	order, err := c.service.CreateOrder(uint(userID))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

func (c *controller) GetOrders(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := c.service.GetOrders(uint(userID))
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)
}

func (c *controller) UpdateOrderStatus(ctx *gin.Context) {
	orderIDStr := ctx.Param("order_id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	var request struct {
		Status string `json:"status"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateOrderStatus(uint(orderID), request.Status); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Order status updated"})
}
