package entities

import "time"

type Package struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	ShipmentID *int      `json:"shipmentID"`
	Shipment   Shipment  `gorm:"foreignKey:ShipmentID"`
	ProductID  *int      `json:"productID"`
	Product    Product   `gorm:"foreignKey:ProductID"`
	Quantity   int       `json:"quantity"`
	Time_stamp time.Time `json:"time_stamp"`
	TrackingNo string    `json:"tracking_no"`
}
