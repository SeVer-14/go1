package dto

type AddToCartDTO struct {
	ProductID uint `json:"productId" binding:"required"`
	Quantity  int  `json:"quantity" binding:"min=1"`
}

//type UpdateOrderStatusDTO struct {
//	OrderID int    `json:"orderId" binding:"required"`
//	Status  string `json:"status" binding:"required,oneof=pending processing completed cancelled"`
//}

type ProductDTO struct {
	ID        uint    `json:"id"`
	ProductID int     `json:"productId" binding:"required"`
	Title     string  `json:"title" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}

type CartItemDTO struct {
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Title     string  `json:"title"`
}

type CartDTO struct {
	CartID uint          `json:"cartId"`
	Items  []CartItemDTO `json:"items"`
	Total  float64       `json:"total"`
}

type OrderItemDTO struct {
	ProductID uint    `json:"productId"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderDTO struct {
	ID     uint           `json:"id"`
	Status string         `json:"status"`
	Total  float64        `json:"total"`
	Items  []OrderItemDTO `json:"items"`
}
type UpdateOrderStatusDTO struct {
	OrderID uint   `json:"orderId" validate:"required"`
	Status  string `json:"status" validate:"required,oneof=pending processing completed cancelled"`
}
