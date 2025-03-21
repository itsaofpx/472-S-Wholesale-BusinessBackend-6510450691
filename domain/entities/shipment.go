package entities

type Shipment struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Address  string `json:"address"`
	S_status string `json:"s_status"`
	OrderId  int    `gorm:"not null" json:"order_id"`
	Order    Order  `json:"foreignKey:OrderId"`
}
