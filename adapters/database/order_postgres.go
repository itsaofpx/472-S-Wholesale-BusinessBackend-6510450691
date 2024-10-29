package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/response"
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

func (opr *OrderPostgreRepository) CreateOrder(o entities.Order) (*response.OrderResponse, error) {
	query := "INSERT INTO public.orders(o_status, o_timestamp, o_total_price, user_id, tracking_number) VALUES ($1, $2, $3, $4, $5) RETURNING id, o_status, o_timestamp, o_total_price, user_id, tracking_number;"

	var order entities.Order

	result := opr.db.Raw(query, o.O_status, o.O_timestamp, o.O_total_price, o.UserID, o.TrackingNumber).Scan(&order)

	var response = &response.OrderResponse{
		Id:            order.Id,
		O_status:      order.O_status,
		O_timestamp:   order.O_timestamp,
		O_total_price: order.O_total_price,
		UserID:        order.UserID,
		TrackingNumber: order.TrackingNumber, // Include tracking number in response
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return response, nil
}


func (opr *OrderPostgreRepository) UpdateOrder(id int, o entities.Order) (entities.Order, error) {
	query := "UPDATE public.orders SET o_status=$2, o_timestamp=$3, o_total_price=$4, user_id=$5, tracking_number=$6 WHERE id = $1 RETURNING id, o_status, o_timestamp, o_total_price, user_id, tracking_number;"

	var order entities.Order

	result := opr.db.Raw(query, id, o.O_status, o.O_timestamp, o.O_total_price, o.UserID, o.TrackingNumber).Scan(&order)

	if result.Error != nil {
		return entities.Order{}, result.Error
	}

	return order, nil
}


func (opr *OrderPostgreRepository) GetOrderByID(id int) (entities.Order, error) {
	query := "SELECT id, o_status, o_timestamp, o_total_price, user_id, tracking_number FROM public.orders WHERE id = $1;"

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

func (opr *OrderPostgreRepository) GetOrderByUserID(userId int) ([]entities.Order, error) {
	query := "SELECT id, o_status, o_timestamp, o_total_price, user_id FROM public.orders WHERE user_id = $1;"

	var orders []entities.Order

	result := opr.db.Raw(query, userId).Scan(&orders)

	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (opr *OrderPostgreRepository) GetOrderAndUserByID(id int) (response.OrderUserResponse, error) {
	query := `
		SELECT 
			orders.id, 
			orders.o_status, 
			orders.o_timestamp, 
			orders.o_total_price, 
			orders.user_id,
			orders.tracking_number, 
			users.f_name, 
			users.l_name, 
			users.phone_number, 
			users.email, 
			users.tier_rank, 
			users.address 
		FROM 
			public.orders 
		JOIN 
			public.users ON orders.user_id = users.id 
		WHERE 
			orders.id = $1;
	`

	var orderUserResponse response.OrderUserResponse

	result := opr.db.Raw(query, id).Scan(&orderUserResponse)


	if result.Error != nil {
		return response.OrderUserResponse{}, result.Error
	}
	return orderUserResponse, nil
}
