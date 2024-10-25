package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type PackageUsecase interface {
	CreatePackage(p entities.Package) (entities.Package, error)
	GetPackageByID(id int) (entities.Package, error)
	GetAllPackagesByShipmentID(id int) ([]entities.Package, error)
	GetAllPackages() ([]entities.Package, error)
}

type PackageService struct {
	repo repositories.PackageRepository
}

func InitiatePackageService(repo repositories.PackageRepository) PackageUsecase {
	return &PackageService{
		repo: repo,
	}
}

func (ps *PackageService) CreatePackage(p entities.Package) (entities.Package, error) {
	createdPackage, err := ps.repo.CreatePackage(p)

	if err != nil {
		return entities.Package{}, err
	}

	return createdPackage, nil
}

func (ps *PackageService) GetPackageByID(id int) (entities.Package, error) {
	getPackage, err := ps.repo.GetPackageByID(id)

	if err != nil {
		return entities.Package{}, err
	}

	return getPackage, nil
}

func (ps *PackageService) GetAllPackagesByShipmentID(id int) ([]entities.Package, error) {
	packages, err := ps.repo.GetAllPackagesByShipmentID(id)

	if err != nil {
		return nil, err
	}

	return packages, nil
}

func (ps *PackageService) GetAllPackages() ([]entities.Package, error) {
	packages, err := ps.repo.GetAllPackages()

	if err != nil {
		return nil, err
	}

	return packages, nil
}
