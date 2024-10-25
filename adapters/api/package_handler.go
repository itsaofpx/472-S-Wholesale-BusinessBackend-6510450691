package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type PackageHandler struct {
	PackageUsecase usecases.PackageUsecase
}

func InitiatePackageHandler(packageUsecase usecases.PackageUsecase) *PackageHandler {
	return &PackageHandler{
		PackageUsecase: packageUsecase,
	}
}

func (ph *PackageHandler) CreatePackage(c *fiber.Ctx) error {

	var newPackage entities.Package

	if err := c.BodyParser(&newPackage); err != nil {

		return errors.New(err.Error())
	}

	createdPackage, err := ph.PackageUsecase.CreatePackage(newPackage)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(createdPackage)
}

func (ph *PackageHandler) GetPackageByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	packageData, err := ph.PackageUsecase.GetPackageByID(id)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(packageData)
}

func (ph *PackageHandler) GetAllPackagesByShipmentID(c *fiber.Ctx) error {
	orderID, err := strconv.Atoi(c.Params("orderID"))

	if err != nil {
		return errors.New(err.Error())
	}

	packages, err := ph.PackageUsecase.GetAllPackagesByShipmentID(orderID)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(packages)
}

func (ph *PackageHandler) GetAllPackages(c *fiber.Ctx) error {
	allPackages, err := ph.PackageUsecase.GetAllPackages()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(allPackages)
}
