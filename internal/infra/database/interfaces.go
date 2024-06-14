package database

import (
	"github.com/PCPedroso/pos-fc-apis/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User)
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product)
	AddNameSuffix(name string)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product)
	Delete(id string) error
}
