package database

import (
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
)

type VehicleInterface interface {
	Create(user *entity.Vehicle) error
	FindAll(page, limit int, sort string) ([]*entity.Vehicle, error)
}
