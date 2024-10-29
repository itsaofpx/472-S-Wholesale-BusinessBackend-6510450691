package usecases

import (
	"errors"
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type ProductUsecase interface {
	CreateProduct(p entities.Product) error
	CreateProducts(products []entities.Product) ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	GetProductByFilter(name string, minprice float64, maxprice float64) ([]entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
	UpdateProduct(id int, p entities.Product) (entities.Product, error)
}

type ProductService struct {
	repo repositories.ProductRepository
}

func InitiateProductsService(repo repositories.ProductRepository) ProductUsecase {
	return &ProductService{
		repo: repo,
	}
}

func (ps *ProductService) CreateProduct(p entities.Product) error {
	if p.Image_url_1 == "" {
		p.Image_url_1 = "https://img5.pic.in.th/file/secure-sv1/Your-paragraph-message.png"
	}
	if p.Image_url_2 == "" {
		p.Image_url_2 = "https://img5.pic.in.th/file/secure-sv1/Your-paragraph-message.png"
	}
	if p.Image_url_3 == "" {
		p.Image_url_3 = "https://img5.pic.in.th/file/secure-sv1/Your-paragraph-message.png"
	}
	

	if err := ps.repo.CreateProduct(p); err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (ps *ProductService) GetProductByID(id int) (entities.Product, error) {
	p, err := ps.repo.GetProductByID(id)

	if err != nil {
		return entities.Product{}, err
	}

	return p, nil
}

func (ps *ProductService) GetAllProducts() ([]entities.Product, error) {
	p_list, err := ps.repo.GetAllProducts()

	if err != nil {
		return []entities.Product{}, err
	}

	return p_list, nil
}

func (ps *ProductService) GetProductByFilter(name string, minprice float64, maxprice float64) ([]entities.Product, error) {
	if name == "" {
		name = "%"
	} else {
		name = "%" + name + "%"
	}
	if minprice <= 0 {
		minprice = 0
	}
	if maxprice <= 0 {
		maxprice = 999999
	}
	p_list, err := ps.repo.GetProductByFilter(name, minprice, maxprice)

	if err != nil {
		return []entities.Product{}, err
	}

	return p_list, nil
}

func (ps *ProductService) UpdateProduct(id int, p entities.Product) (entities.Product, error) {
	up, err := ps.repo.UpdateProduct(id, p)

	if err != nil {
		return entities.Product{}, err
	}

	return up, nil
}

func (ps *ProductService) CreateProducts(products []entities.Product) ([]entities.Product, error) {
	var productLS []entities.Product
	for _, p := range products {
		err := ps.CreateProduct(p)
		if err != nil {
			return nil, err
		}
		productLS = append(productLS, p)
	}
	return productLS, nil
}

