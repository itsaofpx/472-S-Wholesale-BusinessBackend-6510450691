package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type TransactionHandler struct {
	TransactionUsecase usecases.TransactionUsecase
}

func InitiateTransactionHandler(transactionUsecase usecases.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		TransactionUsecase: transactionUsecase,
	}
}

func (th *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	var newTransaction entities.Transaction

	if err := c.BodyParser(&newTransaction); err != nil {
		return errors.New(err.Error())
	}

	transaction, err := th.TransactionUsecase.CreateTransaction(newTransaction)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(transaction)
}

func (th *TransactionHandler) UpdateTransaction(c *fiber.Ctx) error {
	transactionID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	t := new(entities.Transaction)

	if err := c.BodyParser(t); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	transaction, err := th.TransactionUsecase.UpdateTransaction(transactionID, *t)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(transaction)
}

func (th *TransactionHandler) GetTransactionById(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	transaction, err := th.TransactionUsecase.GetTransactionById(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(transaction)

}

func (th *TransactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	transactions, err := th.TransactionUsecase.GetAllTransactions()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(transactions)
}

func (th *TransactionHandler) GetTransactionByOrderId(c *fiber.Ctx) error {
    orderId := c.Params("order_id")
    id, err := strconv.Atoi(orderId)
    if err != nil {
        return err
    }
    transaction, err := th.TransactionUsecase.GetTransactionByOrderId(id)
    if err != nil {
        return err
    }
    return c.JSON(transaction)
}

func (th *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
	transactionID, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = th.TransactionUsecase.DeleteTransaction(transactionID)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)	
	}	

	return c.SendStatus(fiber.StatusOK)
	}
