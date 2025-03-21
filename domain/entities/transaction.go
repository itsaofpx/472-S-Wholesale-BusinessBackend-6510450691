package entities

import "time"

type Transaction struct {
	Id           int       `json:"id" gorm:"primaryKey"`
	T_time_stamp time.Time `json:"t_time_stamp"`
	T_net_price  float64   `json:"t_net_price"`
	T_image_url  string    `json:"t_image_url"`
	OrderId      int       `gorm:"not null;unique" json:"order_id"`
	Order        Order     `json:"foreignKey:OrderId"`
}
