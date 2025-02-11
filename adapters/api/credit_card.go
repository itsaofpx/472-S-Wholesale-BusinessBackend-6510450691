package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type CreditCardHandler struct {
	CreditCardUsecase usecases.CreditCardUseCase
}

func InitiateCreditCardHandler(creditCardUsecase usecases.CreditCardUseCase) *CreditCardHandler {
	return &CreditCardHandler{
		CreditCardUsecase: creditCardUsecase,
	}
}


func (cch *CreditCardHandler) CreateCreditCard(c *fiber.Ctx) error {
	var newCreditCard entities.CreditCard

	if err := c.BodyParser(&newCreditCard); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	
	creditCard, err := cch.CreditCardUsecase.CreateCreditCard(newCreditCard)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(creditCard)
}


func (cch *CreditCardHandler) GetCreditCardByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email is required",
		})
	}

	creditCard, err := cch.CreditCardUsecase.GetCreditCardByEmail(email)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(creditCard)
}


func (cch *CreditCardHandler) UpdateCreditCardByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email is required",
		})
	}

	var req request.UpdateCreditCardRequest

	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	
	creditCard := entities.CreditCard{
		CardNumber:   req.CardNumber,
		CardHolder:   req.CardHolder,
		Expiration:   req.Expiration,
		SecurityCode: req.SecurityCode,
	}

	updatedCard, err := cch.CreditCardUsecase.UpdateCreditCardByEmail(email, creditCard)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(updatedCard)
}


func (cch *CreditCardHandler) DeleteCreditCardByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	
	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email is required",
		})
	}

	err := cch.CreditCardUsecase.DeleteCreditCardByEmail(email)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Credit card deleted successfully",
	})
}
