package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type OrderLineHandler struct {
	OrderLineUsecase usecases.OrderLineUsecase
}

func InitiateOrderLineHandler(orderLineUsecase usecases.OrderLineUsecase) *OrderLineHandler {
	return &OrderLineHandler{
		OrderLineUsecase: orderLineUsecase,
	}
}

// CreateOrderLine creates a new order line entry
func (olh *OrderLineHandler) CreateOrderLine(c *fiber.Ctx) error {
	var newOrderLine entities.OrderLine

	if err := c.BodyParser(&newOrderLine); err != nil {
		return errors.New(err.Error())
	}

	createdOrderLine, err := olh.OrderLineUsecase.CreateOrderLine(newOrderLine)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(createdOrderLine)
}

// GetOrderLineByID retrieves an order line by its ID
func (olh *OrderLineHandler) GetOrderLineByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	orderLine, err := olh.OrderLineUsecase.GetOrderLineByID(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(orderLine)
}

// GetOrderLinesByOrderID retrieves all order lines associated with a specific order
func (olh *OrderLineHandler) GetOrderLinesByOrderID(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("orderID"))

	if err != nil {
		return errors.New(err.Error())
	}

	orderLines, err := olh.OrderLineUsecase.GetOrderLinesByOrderID(orderID)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(orderLines)
}

// GetAllOrderLines retrieves all order lines in the system
func (olh *OrderLineHandler) GetAllOrderLines(c *fiber.Ctx) error {
	allOrderLines, err := olh.OrderLineUsecase.GetAllOrderLines()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(allOrderLines)
}

// UpdateOrderLine updates the details of a specific order line by its ID
func (olh *OrderLineHandler) UpdateOrderLine(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updateOrderLine := new(entities.OrderLine)

	if err := c.BodyParser(updateOrderLine); err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	updatedOrderLine, err := olh.OrderLineUsecase.UpdateOrderLine(id, *updateOrderLine)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(updatedOrderLine)
}

// DeleteOrderLine deletes an order line by its ID
func (olh *OrderLineHandler) DeleteOrderLine(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = olh.OrderLineUsecase.DeleteOrderLine(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (olh *OrderLineHandler) CreateOrderLines(c *fiber.Ctx) error {
	var newOrderLines []entities.OrderLine // Change to a slice to accept multiple order lines

	if err := c.BodyParser(&newOrderLines); err != nil {
		return errors.New(err.Error())
	}

	// Assuming your CreateOrderLine function can handle multiple order lines
	createdOrderLines, err := olh.OrderLineUsecase.CreateOrderLines(newOrderLines) // Adjusted to handle slice

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(createdOrderLines)
}

func (olh *OrderLineHandler) GetOrderLineByOrderIDAndProductID(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("orderID"))
	if err != nil {
		return errors.New(err.Error())
	}
	productID, err := strconv.Atoi(c.Params("productID"))
	if err != nil {
		return errors.New(err.Error())
	}

	orderLine, err := olh.OrderLineUsecase.GetOrderLineByOrderIDAndProductID(orderID, productID)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(orderLine)
	}