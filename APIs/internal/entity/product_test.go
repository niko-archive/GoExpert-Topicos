package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product-001", 100)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, "Product-001", p.Name)
	assert.Equal(t, 100.0, p.Price)

}

func TestProductWhenNameIsRequeired(t *testing.T) {
	p, err := NewProduct("", 100)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameRequired, err)
}

func TestProductWhenPriceIsRequeired(t *testing.T) {
	p, err := NewProduct("Product-001", 0)

	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	p, err := NewProduct("Product-001", -1)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceRequired, err)
}

func TestProductWhenIDIsRequeired(t *testing.T) {
	p, err := NewProduct("Product-001", 100)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
