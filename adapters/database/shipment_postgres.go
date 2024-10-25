package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type ShipmentPostgresRepository struct {
	db *gorm.DB
}

func InitiateShipmentPostgresRepository(db *gorm.DB) repositories.ShipmentRepository {
	return &ShipmentPostgresRepository{
		db: db,
	}
}

func (spr *ShipmentPostgresRepository) CreateShipment(s entities.Shipment) (entities.Shipment, error) {
    query := `
        INSERT INTO public.shipments(s_status, address, order_id) 
        VALUES ($1, 
                (SELECT u.address FROM public.orders o 
                 JOIN public.users u ON o.user_id = u.id 
                 WHERE o.id = $2), 
                $2) 
        RETURNING id, s_status, order_id;
    `

    var verifiedShipment entities.Shipment

    result := spr.db.Raw(query, s.S_status, s.OrderId).Scan(&verifiedShipment)

    if result.Error != nil {
        return entities.Shipment{}, result.Error
    }

    return verifiedShipment, nil
}

func (spr *ShipmentPostgresRepository) UpdateShipment(id int, s entities.Shipment) (entities.Shipment, error) {
    query := "UPDATE public.shipments SET s_status=$2, order_id=$3 WHERE id = $1 RETURNING id, s_status, order_id;"

    var verifiedShipment entities.Shipment

    result := spr.db.Raw(query, id, s.S_status, s.OrderId).Scan(&verifiedShipment)

    if result.Error != nil {
        return entities.Shipment{}, result.Error
    }

    return verifiedShipment, nil
}

func (spr *ShipmentPostgresRepository) GetShipmentByID(id int) (entities.Shipment, error) {
    query := "SELECT id, s_status, order_id FROM public.shipments WHERE id = $1;"

    var shipment entities.Shipment

    result := spr.db.Raw(query, id).Scan(&shipment)

    if result.Error != nil {
        return entities.Shipment{}, result.Error
    }

    return shipment, nil
}

func (spr *ShipmentPostgresRepository) GetAllShipments() ([]entities.Shipment, error) {
    query := "SELECT id, s_status, order_id FROM public.shipments;"

    var shipments []entities.Shipment

    result := spr.db.Raw(query).Scan(&shipments)

    if result.Error != nil {
        return nil, result.Error
    }

    return shipments, nil
}

func (spr *ShipmentPostgresRepository) GetShipmentByOrderID(orderId int) (entities.Shipment, error) {
	query := "SELECT id, s_status, order_id FROM public.shipments WHERE order_id = $1;"

	var shipment entities.Shipment	

	result := spr.db.Raw(query, orderId).Scan(&shipment)	

	if result.Error != nil {
		return entities.Shipment{}, result.Error
	}

	return shipment, nil
}
