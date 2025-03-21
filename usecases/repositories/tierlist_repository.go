package repositories

import "github.com/ppwlsw/sa-project-backend/domain/entities"

type TierListRepository interface {
	GetDiscountPercentByUserID(id int) (*entities.TierList, error)
	InitialTierList(tier int, discount float64) (*entities.TierList, error)
	CreateTireList(tierList entities.TierList) (entities.TierList, error)
	GetAllTierList() ([]entities.TierList, error)
}
