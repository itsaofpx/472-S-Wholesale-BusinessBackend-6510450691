package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type SupplierOrderListHandler struct {
	SupplierOrderListUsecase usecases.SupplierOrderListUsecase
}

func InitiateSupplierOrderListHandler(supplierOrderListUsecase usecases.SupplierOrderListUsecase) *SupplierOrderListHandler {
	return &SupplierOrderListHandler{
		SupplierOrderListUsecase: supplierOrderListUsecase,
	}
}

func (sos *SupplierOrderListHandler) CreateSupplierOrderList(c *fiber.Ctx) error {
	var newSupplierOrderList entities.SupplierOrderList

	if err := c.BodyParser(&newSupplierOrderList); err != nil {
		return errors.New(err.Error())
	}

	createdSupplierOrderList, err := sos.SupplierOrderListUsecase.CreateSupplierOrderList(newSupplierOrderList)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(createdSupplierOrderList)
}

func (sos *SupplierOrderListHandler) UpdateSupplierOrderList(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updateSupplierOrderList := new(entities.SupplierOrderList)

	if err := c.BodyParser(updateSupplierOrderList); err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	verifiedSupplierOrderList, err := sos.SupplierOrderListUsecase.UpdateSupplierOrderList(id, *updateSupplierOrderList)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(verifiedSupplierOrderList)
}

func (sos *SupplierOrderListHandler) GetSupplierOrderListsBySupplierID(c *fiber.Ctx) error {
	supplierID, err := strconv.Atoi(c.Params("supplierID"))

	if err != nil {
		return errors.New(err.Error())
	}

	supplierOrderLists, err := sos.SupplierOrderListUsecase.GetSupplierOrderListsBySupplierID(supplierID)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(supplierOrderLists)

}

func (sos *SupplierOrderListHandler) GetSupplierOrderListByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	supplierOrderList, err := sos.SupplierOrderListUsecase.GetSupplierOrderListByID(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(supplierOrderList)
}

func (sos *SupplierOrderListHandler) GetAllSupplierOrderLists(c *fiber.Ctx) error {
	supplierOrderLists, err := sos.SupplierOrderListUsecase.GetAllSupplierOrderLists()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(supplierOrderLists)
}
