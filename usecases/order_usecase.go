package usecases

import (
	"time"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/response"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type OrderUsecase interface {
	CreateOrder(o entities.Order) (*response.OrderResponse, error)
	UpdateOrder(id int, o entities.Order) (entities.Order, error)
	UpdateOrderStatus(id int, status string) (entities.Order, error)
	GetOrderByID(id int) (entities.Order, error)
	GetOrderByUserID(userId int) ([]entities.Order, error)
	GetOrderAndUserByID(id int) (response.OrderUserResponse, error)
	GetAllOrders() ([]entities.Order, error)
}

type OrderService struct {
	repo repositories.OrderRepository
}

func InitiateOrderService(repo repositories.OrderRepository) OrderUsecase {
	return &OrderService{
		repo: repo,
	}
}

func (os *OrderService) CreateOrder(o entities.Order) (*response.OrderResponse, error) {
	o.O_status = "P"
	o.O_timestamp = time.Now()
	createdOrder, err := os.repo.CreateOrder(o)

	if err != nil {
		return nil, err
	}

	return createdOrder, nil
}

func (os *OrderService) UpdateOrder(id int, o entities.Order) (entities.Order, error) {
	updateOrder, err := os.repo.UpdateOrder(id, o)

	if err != nil {
		return entities.Order{}, err
	}

	return updateOrder, nil
}

func (os *OrderService) GetOrderByID(id int) (entities.Order, error) {
	order, err := os.repo.GetOrderByID(id)

	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}

func (os *OrderService) GetAllOrders() ([]entities.Order, error) {
	orders, err := os.repo.GetAllOrders()

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (os *OrderService) GetOrderByUserID(userId int) ([]entities.Order, error) {
	order, err := os.repo.GetOrderByUserID(userId)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (os *OrderService) GetOrderAndUserByID(id int) (response.OrderUserResponse, error) {
	order, err := os.repo.GetOrderAndUserByID(id)

	if err != nil {
		return response.OrderUserResponse{}, err
	}

	return order, nil
}

func (os *OrderService) UpdateOrderStatus(id int, status string) (entities.Order, error) {
	order, err := os.GetOrderByID(id)
	if err != nil {
		return entities.Order{}, err
	}
	order.O_status = status
	_, err = os.UpdateOrder(id, order)

	if err != nil {
		return entities.Order{}, err
	}

	return order, nil
	}