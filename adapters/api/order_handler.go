package api

import (
	"errors"
	"strconv"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type OrderHandler struct {
	OrderUsecase usecases.OrderUsecase
}

func InitiateOrderHandler(orderUsecase usecases.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		OrderUsecase: orderUsecase,
	}
}

func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var newOrder entities.Order

	if err := c.BodyParser(&newOrder); err != nil {
		return errors.New(err.Error())
	}

	createdOrder, err := oh.OrderUsecase.CreateOrder(newOrder)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(createdOrder)
}

func (oh *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	orderToUpdate := new(entities.Order)

	if err := c.BodyParser(orderToUpdate); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updatedOrder, err := oh.OrderUsecase.UpdateOrder(orderID, *orderToUpdate)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(updatedOrder)
}

func (oh *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	order, err := oh.OrderUsecase.GetOrderByID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(order)
}

func (oh *OrderHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := oh.OrderUsecase.GetAllOrders()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(orders)
}

func (oh *OrderHandler) GetOrderByUserID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	order, err := oh.OrderUsecase.GetOrderByUserID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(order)
}

func (oh *OrderHandler) GetOrderAndUserByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	order, err := oh.OrderUsecase.GetOrderAndUserByID(idParams)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(order)
	}

func (oh *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	
	var req request.UpdateOrderStatusRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	_, err := oh.OrderUsecase.UpdateOrderStatus(req.ID, req.Status)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return nil
}