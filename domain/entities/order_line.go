package entities

type OrderLine struct {
	Id        int     `json:"id" gorm:"primaryKey"`
	OrderID   *int    `json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	ProductID *int    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
