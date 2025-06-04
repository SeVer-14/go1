package entity

type Product struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity"`
}
