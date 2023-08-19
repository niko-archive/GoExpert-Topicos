package database

import (
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"
)

type Vehicle struct {
	DB *gorm.DB
}

func NewVehicle(db *gorm.DB) *Vehicle {
	return &Vehicle{DB: db}
}

func (v *Vehicle) Create(vehicle *entity.Vehicle) error {
	return v.DB.Create(vehicle).Error
}

func (v *Vehicle) FindByFipeCode(fipeCode string) (*entity.Vehicle, error) {
	var vehicle entity.Vehicle
	err := v.DB.Where("fipe_code = ?", fipeCode).First(&vehicle).Error
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (v *Vehicle) FindAll(page, limit int, sort string) ([]*entity.Vehicle, error) {
	var vehicles []*entity.Vehicle
	offset := (page - 1) * limit
	err := v.DB.Offset(offset).Limit(limit).Order("value " + sort).Find(&vehicles).Error
	if err != nil {
		return nil, err
	}
	return vehicles, nil
}
