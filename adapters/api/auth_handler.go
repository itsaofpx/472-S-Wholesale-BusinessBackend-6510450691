package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type AuthHandler struct {
	AuthUsecase usecases.AuthUsecase
}

func InitiateAuthHandler(usecase usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		AuthUsecase: usecase,
	}
}

func (ah *AuthHandler) Register(c *fiber.Ctx) error {
	var user entities.User
	var req request.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can not Parse Body",
		})
	}

	if err := copier.Copy(&user, &req); err != nil {
		fmt.Println("Error copying data:", err)
	}

	getUser, err := ah.AuthUsecase.Register(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Create Successfully",
		"id":      getUser.ID,
	})

}

func (ah *AuthHandler) Login(c *fiber.Ctx) error {
	var req request.LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	user, err := ah.AuthUsecase.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "incorrect email or password",
		})
	}

	return c.JSON(fiber.Map{
		"message": "login successful",
		"user":    user,
	})
}

func (ah *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	var req request.ChangePasswordRequest

	// รับข้อมูลจาก body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	fmt.Println("eiei")
	fmt.Println(req)

	// ส่ง pointer ของ req ไปยัง ChangePassword
	err := ah.AuthUsecase.ChangePassword(&req)  // ส่ง &req แทน req
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "password changed successfully",
	})
}
