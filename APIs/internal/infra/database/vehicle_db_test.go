package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateVehicle(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	db.AutoMigrate(&entity.Vehicle{})

	value := 1000000.99
	brand := "Porche"
	model := "911"
	modelYear := "2021"
	fuel := "Gasolina"
	fipeCode := "123456789"
	referenceMonth := "Jan 2021"
	vehicleType := "Car"

	vehicle, err := entity.NewVehicle(value, brand, model, fuel, fipeCode, referenceMonth, vehicleType, modelYear)
	assert.NoError(t, err)

	vehicleDB := NewVehicle(db)
	err = vehicleDB.Create(vehicle)
	assert.Nil(t, err)

	vehicleFound, err := vehicleDB.FindByFipeCode(fipeCode)
	assert.Nil(t, err)
	assert.Equal(t, vehicleFound.ID, vehicle.ID)
	assert.Equal(t, vehicleFound.Value, vehicle.Value)
	assert.Equal(t, vehicleFound.Brand, vehicle.Brand)
	assert.Equal(t, vehicleFound.Model, vehicle.Model)
	assert.Equal(t, vehicleFound.Fuel, vehicle.Fuel)
	assert.Equal(t, vehicleFound.FipeCode, vehicle.FipeCode)
	assert.Equal(t, vehicleFound.ReferenceMonth, vehicle.ReferenceMonth)
	assert.Equal(t, vehicleFound.VehicleType, vehicle.VehicleType)

	// dateFound := vehicleFound.ModelYear.Format(time.RFC3339)
	// dateExpected := vehicle.ModelYear.Format(time.RFC3339)
	// assert.Equal(t, dateFound, dateExpected)

}
