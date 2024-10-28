package repositories

import ("github.com/ppwlsw/sa-project-backend/domain/entities"; "github.com/ppwlsw/sa-project-backend/domain/response";)

type OrderLineRepository interface {
	CreateOrderLine(ol entities.OrderLine) (entities.OrderLine, error)
	UpdateOrderLine(id int, ol entities.OrderLine) (entities.OrderLine, error)
	GetOrderLineByID(id int) (entities.OrderLine, error)
	GetOrderLinesByOrderID(id int) ([]response.OrderLineResponse, error)
	GetAllOrderLines() ([]entities.OrderLine, error)
	GetOrderLineByOrderIDAndProductID(orderID int, productID int) (entities.OrderLine, error)
	DeleteOrderLine(id int) error
	
}
