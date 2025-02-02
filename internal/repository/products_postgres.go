package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pelicanch1k/ProductGatewayAPI/structs"
)

type ProductsPostgres struct {
	db *sqlx.DB
}

func NewProductsPostgres(db *sqlx.DB) *ProductsPostgres {
	return &ProductsPostgres{db: db}
}

func (p *ProductsPostgres) Create(product structs.Product) (int, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}

	var productId int
	createProduct := "INSERT INTO products (name, price) values ($1, $2) RETURNING id"

	row := tx.QueryRow(createProduct, product.Name, product.Price)
	err = row.Scan(&productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return productId, tx.Commit()
}

func (p *ProductsPostgres) Update(product structs.Product, id int) error {
	if err := p.db.Ping(); err != nil {
		return err
	}

	query := "update products set name = $1, price = $2 where id = $3"

	_, err := p.db.Exec(query, product.Name, product.Price, id)
	return err
}

func (p *ProductsPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM products WHERE id = $1")

	_, err := p.db.Exec(query, id)
	return err
}

func (p *ProductsPostgres) Get(id int) (structs.Product, error) {
	var product structs.Product
	query := "SELECT * FROM products WHERE id = $1"

	if err := p.db.Get(&product, query, id); err != nil {
		return product, err
	}

	return product, nil
}

func (p *ProductsPostgres) GetAll() ([]structs.Product, error) {
	var products []structs.Product

	query := "SELECT * FROM products"
	if err := p.db.Select(&products, query); err != nil {
		return products, err
	}

	return products, nil
}
