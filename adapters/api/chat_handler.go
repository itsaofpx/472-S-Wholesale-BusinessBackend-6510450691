package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type ChatHandler struct {
	Chatusecase usecases.ChatUsecase
}

func InitiateChatHandler(chatUsecase usecases.ChatUsecase) *ChatHandler {
	return &ChatHandler{
		Chatusecase: chatUsecase,
	}
}

func (ch ChatHandler) CreateChat(c *fiber.Ctx) error {
	var newChat entities.Chat

	if err := c.BodyParser(&newChat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if newChat.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "UserID is required"})
	}

	chat, err := ch.Chatusecase.CreateChat(newChat)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(chat)
}


func (ch ChatHandler) GetAllChats(c *fiber.Ctx) error{
	chats, err := ch.Chatusecase.GetAllChats()

	if err != nil {
		return errors.New(err.Error())
	}
	return c.JSON(chats)
}

func (ch ChatHandler) GetChatByUserID(c *fiber.Ctx) error {
	id := c.Params("id")

	chat, err := ch.Chatusecase.GetChatByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(chat)
}