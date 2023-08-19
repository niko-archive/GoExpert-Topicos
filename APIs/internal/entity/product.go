package entity

import (
	"errors"

	"github.dev/nicolasmmb/GoExpert-Topicos/pkg/entity"
	"gorm.io/gorm"
)

var (
	ErrInvalidID     = errors.New("invalid id")
	ErrIdRequired    = errors.New("id is required")
	ErrPriceRequired = errors.New("price is required")
	ErrNameRequired  = errors.New("name is required")
	ErrInvalidPrice  = errors.New("invalid price")
)

type Product struct {
	ID    entity.ID `gorm:"type:uuid;primaryKey;not null;uniqueIndex"`
	Name  string    `gorm:"not null; index:idx_name;"`
	Price float64   `gorm:"not null; index:idx_price;"`
	gorm.Model
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:    entity.NewID(),
		Name:  name,
		Price: price,
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIdRequired
	}
	if p.Name == "" {
		return ErrNameRequired
	}
	if p.Price < 0 {
		return ErrPriceRequired
	}
	if p.Price == 0 {
		return ErrInvalidPrice
	}
	return nil
}
