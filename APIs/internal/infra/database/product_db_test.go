package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.Equal(t, product.CreatedAt.Day(), productFound.CreatedAt.Day())
	assert.Equal(t, product.CreatedAt.Hour(), productFound.CreatedAt.Hour())
	assert.Equal(t, product.CreatedAt.Minute(), productFound.CreatedAt.Minute())
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	for i := 0; i < 30; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product: %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))

	products, err = productDB.FindAll(2, 5, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 5, len(products))

}

func TestProductFindById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)

	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
	assert.Equal(t, product.CreatedAt.Day(), productFound.CreatedAt.Day())
	assert.Equal(t, product.CreatedAt.Hour(), productFound.CreatedAt.Hour())
	assert.Equal(t, product.CreatedAt.Minute(), productFound.CreatedAt.Minute())

}

func TestProductUpdate(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)

	productFound.Name = "Product 2"
	productFound.Price = 20

	err = productDB.Update(productFound)
	assert.NoError(t, err)

	productUpdated, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)

	assert.Equal(t, productFound.ID, productUpdated.ID)
	assert.Equal(t, productFound.Name, productUpdated.Name)
	assert.Equal(t, productFound.Price, productUpdated.Price)
	assert.Equal(t, productFound.CreatedAt.Day(), productUpdated.CreatedAt.Day())
	assert.Equal(t, productFound.CreatedAt.Hour(), productUpdated.CreatedAt.Hour())
	assert.Equal(t, productFound.CreatedAt.Minute(), productUpdated.CreatedAt.Minute())

}

func TestProductUpdateNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFound, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)

	productFound.Name = "Product 2"
	productFound.Price = 20
	productFound.ID = uuid.UUID{}

	err = productDB.Update(productFound)
	assert.Error(t, err)
	assert.Equal(t, "product not found", err.Error())
}

func TestProductDelete(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Equal(t, "record not found", err.Error())
	assert.Empty(t, productFound.Name)
	assert.Empty(t, productFound.Price)
}
