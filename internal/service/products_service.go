package service

import (
	product "github.com/pelicanch1k/rest-api"
	"github.com/pelicanch1k/rest-api/internal/repository"
)

type ProductsService struct {
	repo repository.Repository
}

func NewProductsService(repo repository.Repository) *ProductsService {
	return &ProductsService{repo: repo}
}

func (p *ProductsService) Create(id int, name string, price float64) error {
	return nil
}

func (p *ProductsService) Update(id int, name string, price float64) error {
	return nil
}

func (p *ProductsService) Delete(id int) error {
	return nil
}

func (p *ProductsService) Get(id int) product.Product {
	return product.Product{}
}

func (p *ProductsService) GetAll() []product.Product {
	return []product.Product{}
}
