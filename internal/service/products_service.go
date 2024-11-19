package service

import (
	product "github.com/pelicanch1k/rest-api"
	"github.com/pelicanch1k/rest-api/internal/repository"
)

type ProductsService struct {
	repo *repository.Repository
}

func NewProductsService(repo *repository.Repository) *ProductsService {
	return &ProductsService{repo: repo}
}

func (p *ProductsService) Create(product product.Product) (int, error) {
	return p.repo.Create(product)
}

func (p *ProductsService) Update(product product.Product, id int) error {
	return p.repo.Update(product, id)
}

func (p *ProductsService) Delete(id int) error {
	return p.repo.Delete(id)
}

func (p *ProductsService) Get(id int) (product.Product, error) {
	return p.repo.Get(id)
}

func (p *ProductsService) GetAll() []product.Product {
	return p.repo.GetAll()
}
