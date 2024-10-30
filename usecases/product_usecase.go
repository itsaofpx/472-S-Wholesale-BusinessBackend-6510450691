package usecases

import (
	"errors"

	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/domain/request"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type ProductUsecase interface {
	CreateProduct(p entities.Product) (entities.Product, error)
	CreateProducts(products []entities.Product) ([]entities.Product, error)
	GetProductByID(id int) (entities.Product, error)
	GetProductByFilter(name string, minprice float64, maxprice float64) ([]entities.Product, error)
	GetAllProducts() ([]entities.Product, error)
	UpdateProduct(id int, p entities.Product) (entities.Product, error)
	BuyProduct(req *request.BuyProductRequest) (entities.Product, error)
	BuyProducts(req []request.BuyProductRequest) ([]entities.Product, error)
}

func (ps *ProductService) BuyProduct(req *request.BuyProductRequest) (entities.Product, error) {
	exist, err := ps.GetProductByID(req.ProductID)

	if err != nil {
		return entities.Product{}, err
	}

	if exist.P_amount < req.Quantity {
		return entities.Product{}, errors.New("product is not available")
	}

	exist.P_amount -= req.Quantity
	_, err = ps.UpdateProduct(req.ProductID, exist)
	if err != nil {
		return entities.Product{}, err
	}

	return entities.Product{}, nil
}

type ProductService struct {
	repo repositories.ProductRepository
}

func (ps *ProductService) BuyProducts(req []request.BuyProductRequest) ([]entities.Product, error) {
	var productLS []entities.Product
	for _, p := range req {
		product, err := ps.BuyProduct(&p)
		if err != nil {
			return nil, err
		}
		productLS = append(productLS, product)
	}
	return productLS, nil
}

func InitiateProductsService(repo repositories.ProductRepository) ProductUsecase {
	return &ProductService{
		repo: repo,
	}
}

func (ps *ProductService) CreateProduct(p entities.Product) (entities.Product, error) {
	if p.Image_url_1 == "" {
		p.Image_url_1 = "https://img5.pic.in.th/file/secure-sv1/Your-paragraph-message.png"
	}
	if p.Image_url_2 == "" {
		p.Image_url_2 = "https://img5.pic.in.th/file/secure-sv1/Your-paragraph-message.png"
	}
	if p.Image_url_3 == "" {
		p.Image_url_3 = "https://img5.pic.in.th/file/secure-sv1/Your-paragraph-message.png"
	}

	product, err := ps.repo.CreateProduct(p)
	if err != nil {
		return entities.Product{}, errors.New(err.Error())
	}
	return product, nil
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
		product, err := ps.CreateProduct(p)
		if err != nil {
			return nil, err
		}
		productLS = append(productLS, product)
	}
	return productLS, nil
}
