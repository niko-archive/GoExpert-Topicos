package entity

import (
	"errors"

	"github.dev/nicolasmmb/GoExpert-Topicos/pkg/entity"
)

var (
	ErrInvalidEntity = errors.New("invalid entity")
)

type Vehicle struct {
	ID             entity.ID `gorm:"type:uuid;primaryKey;not null;uniqueIndex"`
	Value          float64   `gorm:"not null; index:idx_value;"`
	Brand          string    `gorm:"not null; index:idx_brand;"`
	Model          string    `gorm:"not null; index:idx_model;"`
	ModelYear      string    `gorm:"not null; index:idx_model_year;"`
	Fuel           string    `gorm:"not null; index:idx_fuel;"`
	FipeCode       string    `gorm:"not null; index:idx_fipe_code;"`
	ReferenceMonth string    `gorm:"not null; index:idx_reference_month;"`
	VehicleType    string    `gorm:"not null; index:idx_vehicle_type;"`
}

func NewVehicle(value float64, brand, model, fuel, fipeCode, referenceMonth, vehicleType, modelYear string) (*Vehicle, error) {
	if brand == "" || model == "" || modelYear == "" || fuel == "" || fipeCode == "" || referenceMonth == "" || vehicleType == "" {
		return nil, ErrInvalidEntity
	}
	return &Vehicle{
		ID:             entity.NewID(),
		Value:          value,
		Brand:          brand,
		Model:          model,
		ModelYear:      modelYear,
		Fuel:           fuel,
		FipeCode:       fipeCode,
		ReferenceMonth: referenceMonth,
		VehicleType:    vehicleType,
	}, nil
}
