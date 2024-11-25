package service

import (
	"github.com/pelicanch1k/ProductGatewayAPI/internal/repository"
	"github.com/pelicanch1k/ProductGatewayAPI/structs"
)

type Auth interface {
	CreateUser(user structs.User) (int, error)
	GenerateJWT(user structs.User) (string, error)
	ParseJWT(tokenString string) (int, error)
}

type Products interface {
	Create(product structs.Product) (int, error)
	Update(product structs.Product, id int) error
	Delete(id int) error
	Get(id int) (structs.Product, error)
	GetAll() ([]structs.Product, error)
}

type Service struct {
	Products
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Products: NewProductsService(repo),
		Auth:     NewAuthService(repo),
	}
}
