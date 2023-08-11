package database

import (
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindById(id string) (*entity.User, error)
	FindAll(page int, limit int, sort string) ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}
