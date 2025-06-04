package entity

type Product struct {
	ID        uint `gorm:"primaryKey"`
	ProductID int
	Title     string `json:"title"`
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
}
