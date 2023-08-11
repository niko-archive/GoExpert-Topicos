package database

import (
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
)

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page int, limit int, sort string) ([]*entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
