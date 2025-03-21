package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type TierListUsecase interface {
	GetDiscountPercentByUserID(id int) (*entities.TierList, error)
	InitialTierList() error
	CreateTireList(tierList entities.TierList) (entities.TierList, error)
	GetAllTierList() ([]entities.TierList, error)
}

type TierListService struct {
	repo repositories.TierListRepository
}

func InitiateTierListService(repo repositories.TierListRepository) *TierListService {
	return &TierListService{repo: repo}
}

func (tls *TierListService) GetDiscountPercentByUserID(id int) (*entities.TierList, error) {

	discount, err := tls.repo.GetDiscountPercentByUserID(id)
	if err != nil {
		return &entities.TierList{}, err
	}

	return discount, nil

}
func (tls *TierListService) InitialTierList() error {

	discountArr := [5]int{0, 10, 15, 20, 30}

	for i := 1; i <= 5; i++ {
		tier := i
		discount := discountArr[i-1]

		var tierList = entities.TierList{
			Tier:            tier,
			DiscountPercent: float64(discount),
		}
		_, err := tls.repo.InitialTierList(tierList.Tier, tierList.DiscountPercent)

		if err != nil {
			return err
		}

	}

	return nil
}

func (tls *TierListService) CreateTireList(tierList entities.TierList) (entities.TierList, error) {

	tierList, err := tls.repo.CreateTireList(tierList)

	if err != nil {
		return entities.TierList{}, err
	}

	return tierList, nil
}

func (tls *TierListService) GetAllTierList() ([]entities.TierList, error) {

	tierList, err := tls.repo.GetAllTierList()
	if err != nil {
		return []entities.TierList{}, err
	}

	return tierList, nil

}
