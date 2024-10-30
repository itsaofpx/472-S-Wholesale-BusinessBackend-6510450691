package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases"
)

type TierListHandler struct {
	TierListUsecase usecases.TierListUsecase
}

func InitiateTierListHandler(TierListUsecase usecases.TierListUsecase) *TierListHandler {
	return &TierListHandler{
		TierListUsecase: TierListUsecase,
	}
}

func (tlh *TierListHandler) GetDiscountPercentByUserID(c *fiber.Ctx) error {
	idParams, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return errors.New(err.Error())
	}

	discount, err := tlh.TierListUsecase.GetDiscountPercentByUserID(idParams)

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(discount)
}

func (tlh *TierListHandler) CreateTierList(c *fiber.Ctx) error {
	var newTier entities.TierList

	if err := c.BodyParser(&newTier); err != nil {
		return errors.New(err.Error())
	}

	tierList, err := tlh.TierListUsecase.CreateTireList(newTier)

	if err != nil {
		return errors.New(err.Error())

	}

	return c.JSON(tierList)
}

func (tlh *TierListHandler) GetAllTierList(c *fiber.Ctx) error {
	tierLists, err := tlh.TierListUsecase.GetAllTierList()

	if err != nil {
		return errors.New(err.Error())
	}

	return c.JSON(tierLists)
}
