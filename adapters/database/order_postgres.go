package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type OrderPostgreRepository struct {
	db *gorm.DB
}

func InitiateOrderPostgresRepository(db *gorm.DB) repositories.OrderRepository {
	return &OrderPostgreRepository{
		db: db,
	}
}

func (opr *OrderPostgreRepository) CreateOrder(o entities.Order) (entities.Order, error) {
	query := "INSERT INTO public.orders(o_status, o_timestamp, o_total_price, user_id) VALUES ($1, $2, $3, $4) RETURNING id, o_status, o_timestamp, o_total_price, user_id;"

	var order entities.Order

	result := opr.db.Raw(query, o.O_status, o.O_timestamp, o.O_total_price, o.UserID).Scan(&order)

	if result.Error != nil {
		return entities.Order{}, result.Error
	}

	return order, nil
}

func (opr *OrderPostgreRepository) UpdateOrder(id int, o entities.Order) (entities.Order, error) {
	query := "UPDATE public.orders SET o_status=$2, o_timestamp=$3, o_total_price=$4, user_id=$5 WHERE id = $1 RETURNING id, o_status, o_timestamp, o_total_price, user_id;"

	var order entities.Order

	result := opr.db.Raw(query, id, o.O_status, o.O_timestamp, o.O_total_price, o.UserID).Scan(&order)

	if result.Error != nil {
		return entities.Order{}, result.Error
	}

	return order, nil
}

func (opr *OrderPostgreRepository) GetOrderByID(id int) (entities.Order, error) {
	query := "SELECT id, o_status, o_timestamp, o_total_price, user_id FROM public.orders WHERE id = $1;"

	var order entities.Order

	result := opr.db.Raw(query, id).Scan(&order)

	if result.Error != nil {
		return entities.Order{}, result.Error
	}

	return order, nil
}

func (opr *OrderPostgreRepository) GetAllOrders() ([]entities.Order, error) {
	query := "SELECT id, o_status, o_timestamp, o_total_price, user_id FROM public.orders;"

	var orders []entities.Order

	result := opr.db.Raw(query).Scan(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}