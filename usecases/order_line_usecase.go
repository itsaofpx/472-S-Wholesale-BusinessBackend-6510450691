package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type OrderLineUsecase interface {
	CreateOrderLine(ol entities.OrderLine) (entities.OrderLine, error)
	UpdateOrderLine(id int, ol entities.OrderLine) (entities.OrderLine, error)
	GetOrderLineByID(id int) (entities.OrderLine, error)
	GetOrderLinesByOrderID(id int) ([]entities.OrderLine, error)
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
	col, err := os.repo.CreateOrderLine(ol)

	if err != nil {
		return entities.OrderLine{}, nil
	}

	verifiedOrderLine, err := os.repo.GetOrderLineByID(col.Id)

	if err != nil {
		return entities.OrderLine{}, err
	}

	return verifiedOrderLine, nil
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

func (os *OrderLineService) GetOrderLinesByOrderID(id int) ([]entities.OrderLine, error) {
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
