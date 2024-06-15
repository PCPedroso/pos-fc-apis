package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/PCPedroso/pos-fc-apis/internal/entity"
	"github.com/stretchr/testify/assert"
)

type _p struct {
	Name  string
	Price float64
}

var _product = _p{
	Name:  "Produto",
	Price: 10.0,
}

func (p _p) NameSuffix(s string) string {
	return fmt.Sprintf("%v %v", p.Name, s)
}

func TestCreateProduct(t *testing.T) {
	db, err := ConectaDB(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct(_product.Name, _product.Price)
	assert.Nil(t, err)

	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAll(t *testing.T) {
	db, err := ConectaDB(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("%v %d", _product.Name, i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, _product.NameSuffix("1"), products[0].Name)
	assert.Equal(t, _product.NameSuffix("10"), products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, _product.NameSuffix("11"), products[0].Name)
	assert.Equal(t, _product.NameSuffix("20"), products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, _product.NameSuffix("21"), products[0].Name)
	assert.Equal(t, _product.NameSuffix("23"), products[2].Name)
}

func TestFindByID(t *testing.T) {
	db, err := ConectaDB(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct(_product.Name, _product.Price)
	assert.Nil(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	product, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, _product.Name, product.Name)
}

func TestUpdate(t *testing.T) {
	db, err := ConectaDB(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct(_product.Name, _product.Price)
	assert.Nil(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	product.Name = fmt.Sprintf("%v %v", product.Name, "1")
	err = productDB.Update(product)
	assert.Nil(t, err)

	finded, err := productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.Name, finded.Name)
}

func TestDelete(t *testing.T) {
	db, err := ConectaDB(&entity.Product{})
	if err != nil {
		t.Error(err)
	}

	product, err := entity.NewProduct(_product.Name, _product.Price)
	assert.Nil(t, err)

	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)

	_, err = productDB.FindByID(product.ID.String())
	assert.NotNil(t, err)
}
