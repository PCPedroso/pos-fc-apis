package entity

import (
	"errors"
	"fmt"
	"time"

	"github.com/PCPedroso/pos-fc-apis/pkg/entity"
)

var (
	ErrIDIsRiquired    = errors.New("id is required")
	ErrIDIsInvalid     = errors.New("invalid id")
	ErrNameIsRiquired  = errors.New("name is required")
	ErrPriceIsRiquired = errors.New("price is required")
	ErrPriceIsInvalid  = errors.New("invalid id")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (product *Product) TableName() string {
	return "products"
}

func (product *Product) NameSuffix(suffix string) string {
	return fmt.Sprintf("%v %v", product.Name, suffix)
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRiquired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}
	if p.Name == "" {
		return ErrNameIsRiquired
	}
	if p.Price == 0 {
		return ErrPriceIsRiquired
	}
	if p.Price < 0 {
		return ErrPriceIsInvalid
	}
	return nil
}
