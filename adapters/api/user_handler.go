package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/domain/response"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type UserHandler struct {
	UserUsecase usecases.UserUseCase
}

func InitiateUserHandler(userUsecase usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUsecase: userUsecase,
	}
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return errors.New(err.Error())
	}

	user, err := uh.UserUsecase.GetUserByID(idParams)

	if user == nil {
		return c.JSON(fiber.Map{
			"message": "can't find user",
		})
	}

	if err != nil {
		return errors.New(err.Error())
	}

	var res response.UserResponse

	if err := copier.Copy(&res, &user); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(res)

}

func (uh *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := uh.UserUsecase.GetAllUsers()

	if users == nil {
		return c.JSON(fiber.Map{
			"message": "can't find user",
		})
	}

	if err != nil {
		return errors.New(err.Error())
	}

	var userResponses []response.GetUserResponse
	var res response.GetUserResponse

	for _, user := range *users {
		copier.Copy(&res, &user)
		userResponses = append(userResponses, res)
	}

	return c.JSON(userResponses)
}

func (uh *UserHandler) UpdateUserByID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}

	// อ่านค่าจาก body
	req := new(request.UpdateUserByIDRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	// เรียกใช้ usecase เพื่ออัปเดตข้อมูล
	updatedUser, err := uh.UserUsecase.UpdateUserByID(idParams, req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
		})
	}

	// คืนค่าข้อมูลที่อัปเดต
	var res response.UserResponse
	if err := copier.Copy(&res, &updatedUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to map response data",
		})
	}

	return c.JSON(res)
}


func (uh *UserHandler) UpdateTierByUserID(c *fiber.Ctx) error {

	req := new(request.UpdateTierByUserIDRequest)

	if err := c.BodyParser(&req); err != nil {
		return errors.New(err.Error())
	}

	user, err := uh.UserUsecase.UpdateTierByUserID(req)

	if user == nil {
		return c.JSON(fiber.Map{
			"message": "can't find user",
		})
	}

	if err != nil {
		return errors.New(err.Error())
	}

	var res response.UserResponse

	if err := copier.Copy(&res, &user); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	return c.JSON(res)

}
