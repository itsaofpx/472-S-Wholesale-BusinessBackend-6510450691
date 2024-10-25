package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type ShipmentRepository interface {
	CreateShipment(s entities.Shipment) (entities.Shipment, error)
	UpdateShipment(id int, s entities.Shipment) (entities.Shipment, error)
	GetShipmentByID(id int) (entities.Shipment, error)
	GetShipmentByOrderID(orderId int) (entities.Shipment, error)
	GetAllShipments() ([]entities.Shipment, error)
}
