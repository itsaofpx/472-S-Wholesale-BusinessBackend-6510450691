package api

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type MessageHandler struct {
	MessageUsecase usecases.MessageUsecase
}

func InitiateMessageHandler(messageUsecase usecases.MessageUsecase) *MessageHandler {
	return &MessageHandler{
		MessageUsecase: messageUsecase,
	}
}

func (ch MessageHandler) CreateMessage(c *fiber.Ctx) error {
	var messageCreate entities.Message
	
	chatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid ChatID",
			})
	}
	
	if err := c.BodyParser(&messageCreate); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid request body",
			})
	}
	
	if messageCreate.UserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "UserID is required",
			})
	}
	
	if messageCreate.Body == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Message body is required",
			})
	}
	
	newMessage := entities.Message{
			Timestamp: time.Now(),
			Body:      messageCreate.Body,
			UserID:    messageCreate.UserID,
			ChatID:    chatID,
	}
	
	message, err := ch.MessageUsecase.CreateMessage(newMessage)
	if err != nil {
			// Check if it's a validation error (not found)
			if strings.Contains(err.Error(), "not found") {
					return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
							"error": err.Error(),
					})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
			})
	}
	
	return c.Status(fiber.StatusCreated).JSON(message)
}

func (ch MessageHandler) CreateMessageByChatID(c *fiber.Ctx) error {
	var messageCreate entities.Message

	chatID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ChatID",
		})
	}

	if err := c.BodyParser(&messageCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if messageCreate.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserID is required",
		})
	}

	if messageCreate.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Message body is required",
		})
	}

	newMessage := entities.Message{
		Timestamp: time.Now(),
		Body:      messageCreate.Body,
		UserID:    messageCreate.UserID,
		ChatID:    chatID,
	}

	message, err := ch.MessageUsecase.CreateMessageByChatID(newMessage)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(message)
}

