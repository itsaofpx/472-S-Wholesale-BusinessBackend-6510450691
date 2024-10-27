package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type OrderLinePostgresRepository struct {
	db *gorm.DB
}

func InitiateOrderLinePostgresRepository(db *gorm.DB) repositories.OrderLineRepository {
	return &OrderLinePostgresRepository{
		db: db,
	}
}

func (ops *OrderLinePostgresRepository) CreateOrderLine(ol entities.OrderLine) (entities.OrderLine, error) {
    query := "INSERT INTO public.order_lines(order_id, product_id, price, quantity) VALUES ($1, $2, (SELECT p_price FROM public.products WHERE id = $2) * $3, $3) RETURNING id, order_id, product_id, quantity;"
    var orderLine entities.OrderLine
    result := ops.db.Raw(query, ol.OrderID, ol.ProductID, ol.Quantity).Scan(&orderLine)
    if result.Error != nil {
        return entities.OrderLine{}, result.Error
    }
    return orderLine, nil
}

func (ops *OrderLinePostgresRepository) UpdateOrderLine(id int, ol entities.OrderLine) (entities.OrderLine, error) {
	query := "UPDATE public.order_lines SET order_id=$2, product_id=$3, price=$4, quantity=$5 WHERE id = $1 RETURNING id, order_id, product_id, price, quantity;"

	var updateOrderLine entities.OrderLine

	result := ops.db.Raw(query, id, ol.OrderID, ol.ProductID, ol.Price, ol.Quantity).Scan(&updateOrderLine)

	if result.Error != nil {
		return entities.OrderLine{}, result.Error
	}

	return updateOrderLine, nil
}

func (ops *OrderLinePostgresRepository) GetOrderLineByID(id int) (entities.OrderLine, error) {
	query := "SELECT id, order_id, product_id, price, quantity FROM public.order_lines WHERE id = $1;"

	var orderLine entities.OrderLine

	result := ops.db.Raw(query, id).Scan(&orderLine)

	if result.Error != nil {
		return entities.OrderLine{}, result.Error
	}

	return orderLine, nil
}

func (ops *OrderLinePostgresRepository) GetOrderLinesByOrderID(id int) ([]entities.OrderLine, error) {
	query := "SELECT id, order_id, product_id, price, quantity FROM public.order_lines WHERE order_id = $1;"

	var orderLines []entities.OrderLine

	result := ops.db.Raw(query, id).Scan(&orderLines)

	if result.Error != nil {
		return nil, result.Error
	}

	return orderLines, nil
}

func (ops *OrderLinePostgresRepository) GetAllOrderLines() ([]entities.OrderLine, error) {
	query := "SELECT id, order_id, product_id, price, quantity FROM public.order_lines;"

	var orderLines []entities.OrderLine

	result := ops.db.Raw(query).Scan(&orderLines)

	if result.Error != nil {
		return nil, result.Error
	}

	return orderLines, nil
}

func (ops *OrderLinePostgresRepository) DeleteOrderLine(id int) error {
	query := "DELETE FROM public.order_lines WHERE id = $1;"

	result := ops.db.Exec(query, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ops *OrderLinePostgresRepository) GetOrderLineByOrderIDAndProductID(orderID int, productID int) (entities.OrderLine, error) {
	query := "SELECT id, order_id, product_id, price, quantity FROM public.order_lines WHERE order_id = $1 AND product_id = $2;"

	var orderLine entities.OrderLine

	result := ops.db.Raw(query, orderID, productID).Scan(&orderLine)

	if result.Error != nil {
		return entities.OrderLine{}, result.Error
	}

	return orderLine, nil
}
