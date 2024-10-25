package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type ShipmentHandler struct {
	ShipmentUsecase usecases.ShipmentUsecase
}

func InitiateShipmentHandler(shipmentUsecase usecases.ShipmentUsecase) *ShipmentHandler {
	return &ShipmentHandler{
		ShipmentUsecase: shipmentUsecase,
	}
}

func (sh *ShipmentHandler) CreateShipment(c *fiber.Ctx) error {
	var newShipment entities.Shipment

	if err := c.BodyParser(&newShipment); err != nil {
		return errors.New(err.Error())
	}

	verifiedShipment, err := sh.ShipmentUsecase.CreateShipment(newShipment)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(verifiedShipment)
}

func (sh *ShipmentHandler) UpdateShipment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updateShipment := new(entities.Shipment)

	if err := c.BodyParser(updateShipment); err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	verifiedShipment, err := sh.ShipmentUsecase.UpdateShipment(id, *updateShipment)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(verifiedShipment)
}

func (sh *ShipmentHandler) GetShipmentByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	shipment, err := sh.ShipmentUsecase.GetShipmentByID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(shipment)
}

func (sh *ShipmentHandler) GetAllShipments(c *fiber.Ctx) error {
	shipments, err := sh.ShipmentUsecase.GetAllShipments()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(shipments)
}

func (sh *ShipmentHandler) GetShipmentByOrderID(c *fiber.Ctx) error {	
	idParams, err := strconv.Atoi(c.Params("order_id"))	
	if err != nil {
		return errors.New(err.Error())
	}

	shipment, err := sh.ShipmentUsecase.GetShipmentByOrderID(idParams)
	if err != nil {
		return errors.New(err.Error())
	}
	return c.JSON(shipment)
}
