package api

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type ProductHandler struct {
	ProductUsecase usecases.ProductUsecase
}

func InitiateProductHandler(productUsecase usecases.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (ph *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var newProduct entities.Product

	if err := c.BodyParser(&newProduct); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := ph.ProductUsecase.CreateProduct(newProduct)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func (ph *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	product, err := ph.ProductUsecase.GetProductByID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(product)
}

func (ph *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := ph.ProductUsecase.GetAllProducts()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(products)
}

func (ph *ProductHandler) GetProductByFilter(c *fiber.Ctx) error {
	var req request.FilterProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}
	products, err := ph.ProductUsecase.GetProductByFilter(req.Name, req.MinPrice, req.MaxPrice)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(products)
}

func (ph *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(entities.Product)

	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := ph.ProductUsecase.UpdateProduct(productID, *p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func (ph *ProductHandler) CreateProducts(c *fiber.Ctx) error {
	var products []entities.Product

	if err := c.BodyParser(&products); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p, err := ph.ProductUsecase.CreateProducts(products)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(p)
}

func (ph *ProductHandler) BuyProduct(c *fiber.Ctx) error {
	var req request.BuyProductRequest

	if err := c.BodyParser(&req); err != nil {
		return c.JSON(fiber.Map{"MESSAGE": "Cannot parse request body"})
	}

	product, err := ph.ProductUsecase.BuyProduct(&req)

	if err != nil {
		return c.JSON(fiber.Map{"MESSAGE": "Can't buy product"})
	}

	return c.JSON(product)
}

func (ph *ProductHandler) BuyProducts(c *fiber.Ctx) error {
	var reqs []request.BuyProductRequest

	if err := c.BodyParser(&reqs); err != nil {
		return c.JSON(fiber.Map{"MESSAGE": fmt.Sprintf("Error: %v", err)})
	}

	products, err := ph.ProductUsecase.BuyProducts(reqs)
	if err != nil {
		return c.JSON(fiber.Map{"MESSAGE": fmt.Sprintf("Error: %v", err)})
	}

	return c.JSON(products)
}
