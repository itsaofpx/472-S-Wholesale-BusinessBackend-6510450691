package database

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
	"gorm.io/gorm"
)

type TierListPostgres struct {
	db *gorm.DB
}

func InitiateTierListPostgres(db *gorm.DB) repositories.TierListRepository {
	return &TierListPostgres{
		db: db,
	}

}

func (t *TierListPostgres) GetDiscountPercentByUserID(id int) (*entities.TierList, error) {
	query := "SELECT t.tier, t.discount_percent FROM users AS u JOIN tier_lists AS t ON u.tier_rank = t.tier WHERE u.id = ?;"

	var tierList *entities.TierList

	result := t.db.Raw(query, id).Scan(&tierList)

	if result.Error != nil {
		return &entities.TierList{}, result.Error
	}

	return tierList, nil
}

func (t *TierListPostgres) InitialTierList(tier int, discount float64) (*entities.TierList, error) {
	query := "INSERT INTO public.tier_lists(tier, discount_percent) VALUES ($1, $2);"

	var tierList *entities.TierList

	result := t.db.Raw(query, tier, discount).Scan(&tierList)

	if result.Error != nil {
		return &entities.TierList{}, result.Error
	}

	return tierList, nil
}

func (t *TierListPostgres) CreateTireList(tierList entities.TierList) (entities.TierList, error) {
	query := "INSERT INTO public.tier_lists(tier, discount_percent) VALUES ($1, $2) RETURNING tier, discount_percent;"

	var newTierList entities.TierList

	result := t.db.Raw(query, tierList.Tier, tierList.DiscountPercent).Scan(&newTierList)

	if result.Error != nil {
		return entities.TierList{}, result.Error
	}

	return newTierList, nil
}

func (t *TierListPostgres) GetAllTierList() ([]entities.TierList, error) {
	query := "SELECT tier, discount_percent FROM public.tier_lists;"

	var tierList []entities.TierList

	result := t.db.Raw(query).Scan(&tierList)

	if result.Error != nil {
		return []entities.TierList{}, result.Error
	}

	return tierList, nil
}
