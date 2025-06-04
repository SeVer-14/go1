package main

import (
	"github.com/gin-gonic/gin"
	"go1/controller"
	"go1/database"
	"go1/service"
	"log"
)

var (
	productService    service.ProductService       = service.NewProductService()
	productController controller.ProductController = controller.NewProductController(productService)
)

func main() {

	// БД и миграции
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.Show())
	})
	server.POST("/products", func(ctx *gin.Context) {
		ctx.JSON(200, productController.Add(ctx))
	})
	server.DELETE("/products/:id", productController.Delete)

	server.POST("/cart", productController.AddToCart)
	server.GET("/cart/:user_id", productController.GetCart)
	server.DELETE("/cart/:user_id/:cart_id", productController.RemoveFromCart)
	server.PUT("/cart/:user_id/:cart_id", productController.UpdateCartItem)

	server.POST("/orders/:user_id", productController.CreateOrder)
	server.GET("/orders/:user_id", productController.GetOrders)
	server.PUT("/orders/:order_id/status", productController.UpdateOrderStatus)

	server.Run(":8080")
}
