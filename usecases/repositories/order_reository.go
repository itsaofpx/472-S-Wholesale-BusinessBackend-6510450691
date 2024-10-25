package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type OrderRepository interface {
	CreateOrder(o entities.Order) (entities.Order, error)
	UpdateOrder(id int, o entities.Order) (entities.Order, error)
	GetOrderByID(id int) (entities.Order, error)
	GetOrderByUserID(userId int) ([]entities.Order, error)
	GetAllOrders() ([]entities.Order, error)
}
