package database

import "github.com/PCPedroso/pos-fc-apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User)
	FindByEmail(email string) (*entity.User, error)
}
