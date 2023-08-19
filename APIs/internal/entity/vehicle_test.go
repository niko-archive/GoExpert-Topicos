package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVehicle(t *testing.T) {
	value := 1000000.99
	brand := "Porche"
	model := "911"
	modelYear := "2021"
	fuel := "Gasolina"
	fipeCode := "123456789"
	referenceMonth := "Jan 2021"
	vehicleType := "Car"

	user, err := NewVehicle(value, brand, model, fuel, fipeCode, referenceMonth, vehicleType, modelYear)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, value, user.Value)
	assert.Equal(t, brand, user.Brand)
	assert.Equal(t, model, user.Model)
	assert.Equal(t, modelYear, user.ModelYear)
	assert.Equal(t, fuel, user.Fuel)
	assert.Equal(t, fipeCode, user.FipeCode)
	assert.Equal(t, referenceMonth, user.ReferenceMonth)
	assert.Equal(t, vehicleType, user.VehicleType)

}
