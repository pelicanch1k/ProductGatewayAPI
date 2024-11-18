package repository

import (
	"github.com/jmoiron/sqlx"
	product "github.com/pelicanch1k/rest-api"
)

type ProductsPostgres struct {
	db *sqlx.DB
}

func NewProductsPostgres(db *sqlx.DB) *ProductsPostgres {
	return &ProductsPostgres{db: db}
}

func (p *ProductsPostgres) Create(id int, name string, price float64) error {
	return nil
}

func (p *ProductsPostgres) Update(id int, name string, price float64) error {
	return nil
}

func (p *ProductsPostgres) Delete(id int) error {
	return nil
}

func (p *ProductsPostgres) Get(id int) product.Product {
	return product.Product{}
}

func (p *ProductsPostgres) GetAll() []product.Product {
	return []product.Product{}
}
