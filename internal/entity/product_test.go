package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _p = Product{
	Name:  "Produto 1",
	Price: 10.0,
}

func TestNewProduct(t *testing.T) {
	p, err := NewProduct(_p.Name, _p.Price)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, _p.Name, p.Name)
	assert.Equal(t, _p.Price, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", _p.Price)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRiquired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct(_p.Name, 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRiquired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct(_p.Name, -1)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsInvalid, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct(_p.Name, _p.Price)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
