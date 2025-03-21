package entities

type SupplierOrderList struct {
	Id         int      `json:"id" gorm:"primaryKey"`
	SupplierID *int     `json:"supplier_id"`
	Supplier   Supplier `gorm:"foreignKey:SupplierID"`
	ProductID  *int     `json:"product_id"`
	Product    Product  `gorm:"foreignKey:ProductID"`
	Quantity   int      `json:"quantity"`
	Price      float64  `json:"price"`
}
