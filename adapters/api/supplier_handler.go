package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type SupplierHandler struct {
	SupplierUsecase usecases.SupplierUsecase
}

func InitiateSupplierHandler(supplierUsecase usecases.SupplierUsecase) *SupplierHandler {
	return &SupplierHandler{
		SupplierUsecase: supplierUsecase,
	}
}

func (sh *SupplierHandler) CreateSupplier(c *fiber.Ctx) error {
	var newSupplier entities.Supplier

	if err := c.BodyParser(&newSupplier); err != nil {
		return errors.New(err.Error())
	}

	createdSupplier, err := sh.SupplierUsecase.CreateSupplier(newSupplier)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(createdSupplier)
}

func (sh *SupplierHandler) UpdateSupplier(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	updateSupplier := new(entities.Supplier)

	if err := c.BodyParser(updateSupplier); err != nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	verifiedSupplier, err := sh.SupplierUsecase.UpdateSupplier(id, *updateSupplier)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(verifiedSupplier)
}

func (sh *SupplierHandler) GetSupplierByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	supplier, err := sh.SupplierUsecase.GetSupplierByID(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(supplier)
}

func (sh *SupplierHandler) GetAllSuppliers(c *fiber.Ctx) error {
	suppliers, err := sh.SupplierUsecase.GetAllSuppliers()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(suppliers)
}
