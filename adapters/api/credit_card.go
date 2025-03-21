package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"strconv"
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

func (cch *CreditCardHandler) GetCreditCardByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	creditCard, err := cch.CreditCardUsecase.GetCreditCardByUserID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(creditCard)
}

func (cch *CreditCardHandler) UpdateCreditCardByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
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

	updatedCard, err := cch.CreditCardUsecase.UpdateCreditCardByUserID(id, creditCard)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(updatedCard)
}

func (cch *CreditCardHandler) DeleteCreditCardByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	err = cch.CreditCardUsecase.DeleteCreditCardByUserID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"message": "Credit card deleted successfully",
	})
}

func (cch *CreditCardHandler) GetCreditCardsByUserID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	creditCards, err := cch.CreditCardUsecase.GetCreditCardsByUserID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(creditCards)
}

func (cch *CreditCardHandler) DeleteByCardNumber(c *fiber.Ctx) error {
	cardNumber := c.Params("card_number")
	if cardNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Card number is required",
		})
	}

	err := cch.CreditCardUsecase.DeleteByCardNumber(cardNumber)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Credit card not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Credit card deleted successfully",
	})
}
