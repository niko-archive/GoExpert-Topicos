package filter

import (
	"log"
	"net/http"
	"testing"

	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// func TestParseParamEqual(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=equal.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "=", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamNotEqual(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=not-equal.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "!=", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamGreaterThan(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=greater.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, ">", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamLessThan(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=less.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "<", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamGreaterThanOrEqual(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=greater-equal.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, ">=", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamLessThanOrEqual(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=less-equal.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "<=", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamLike(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=like.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "like", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamIn(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=in.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "in", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamNotIn(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=not-in.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "not in", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamBetween(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=between.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "between", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamNotBetween(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=not-between.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "not between", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamAnd(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=and.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "and", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamOr(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=or.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "or", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamNot(t *testing.T) {
// 	key, operator, value, err := ParseParam("id=not.1")
// 	assert.Equal(t, nil, err)

// 	assert.Equal(t, "id", key)
// 	assert.Equal(t, "not", operator)
// 	assert.Equal(t, "1", value)
// }

// func TestParseParamInvalid(t *testing.T) {
// 	_, _, _, err := ParseParam("id=invalid.1")
// 	assert.NotEqual(t, nil, err)
// }

// func TestValidateOperator(t *testing.T) {
// 	operator, err := ValidateOperator("equal")
// 	assert.Equal(t, nil, err)
// 	assert.Equal(t, "=", operator)
// }

func TestValidateOperatorInvalid(t *testing.T) {

	ops := []Operator{
		{Type: Equal},
		{Type: NotEqual},
	}

	exp := ExpectedParam{
		Param:          "id",
		ValidOperators: ops,
	}

	request, _ := http.NewRequest("GET", "http://localhost:3000?id=equal.180a725e-0bc3-4159-9f68-dfce90f10ac0", nil)
	// exp.ValidateExpectedParam(Equal)
	valuex, _ := exp.GetFromRequest(request)

	log.Println(valuex)

	db, err := gorm.Open(sqlite.Open("./file.sql"), &gorm.Config{})

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

	db.Create(&vehicle)


	GenerateQueryGORM(entity.Vehicle{}, db, valuex)

}
