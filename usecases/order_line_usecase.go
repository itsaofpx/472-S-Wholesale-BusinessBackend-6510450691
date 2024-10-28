package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"github.com/ppwlsw/sa-project-backend/domain/response"
)

type OrderLineUsecase interface {
	CreateOrderLine(ol entities.OrderLine) (entities.OrderLine, error)
	CreateOrderLines(ols []entities.OrderLine) ([]entities.OrderLine, error)
	UpdateOrderLine(id int, ol entities.OrderLine) (entities.OrderLine, error)
	GetOrderLineByID(id int) (entities.OrderLine, error)
	GetOrderLinesByOrderID(id int) ([]response.OrderLineResponse, error)
	GetOrderLineByOrderIDAndProductID(orderID int, productID int) (entities.OrderLine, error)
	GetAllOrderLines() ([]entities.OrderLine, error)
	DeleteOrderLine(id int) error
}

type OrderLineService struct {
	repo repositories.OrderLineRepository
}

func InitiateOrderLineService(repo repositories.OrderLineRepository) OrderLineUsecase {
	return &OrderLineService{
		repo: repo,
	}
}

func (os *OrderLineService) CreateOrderLine(ol entities.OrderLine) (entities.OrderLine, error) {
	createdOrderLine, err := os.repo.CreateOrderLine(ol)

	if err != nil {
		return entities.OrderLine{}, nil
	}

	return createdOrderLine, nil
}

func (os *OrderLineService) UpdateOrderLine(id int, ol entities.OrderLine) (entities.OrderLine, error) {
	updateOrderLine, err := os.repo.UpdateOrderLine(id, ol)

	if err != nil {
		return entities.OrderLine{}, err
	}

	return updateOrderLine, nil
}

func (os *OrderLineService) GetOrderLineByID(id int) (entities.OrderLine, error) {
	orderLine, err := os.repo.GetOrderLineByID(id)

	if err != nil {
		return entities.OrderLine{}, err
	}

	return orderLine, nil
}

func (os *OrderLineService) GetOrderLinesByOrderID(id int) ([]response.OrderLineResponse, error) {
	orderLines, err := os.repo.GetOrderLinesByOrderID(id)

	if err != nil {
		return nil, err
	}

	return orderLines, nil
}

func (os *OrderLineService) GetAllOrderLines() ([]entities.OrderLine, error) {
	orderLines, err := os.repo.GetAllOrderLines()

	if err != nil {
		return nil, err
	}

	return orderLines, nil
}

func (os *OrderLineService) DeleteOrderLine(id int) error {
	err := os.repo.DeleteOrderLine(id)

	if err != nil {
		return err
	}

	return nil
}

func (os *OrderLineService) CreateOrderLines(ols []entities.OrderLine) ([]entities.OrderLine, error) {
    var createdOrderLines []entities.OrderLine

    for _, ol := range ols {
        createdOrderLine, err := os.repo.CreateOrderLine(ol)
        if err != nil {
            return createdOrderLines, err // Return the partially created slice and the error
        }
        createdOrderLines = append(createdOrderLines, createdOrderLine)
    }

    return createdOrderLines, nil
}

func (os *OrderLineService) GetOrderLineByOrderIDAndProductID(orderID int, productID int) (entities.OrderLine, error) {
	orderLine, err := os.repo.GetOrderLineByOrderIDAndProductID(orderID, productID)

	if err != nil {
		return entities.OrderLine{}, err
	}

	return orderLine, nil
}