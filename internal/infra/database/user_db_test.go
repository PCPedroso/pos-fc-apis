package database

import (
	"testing"

	"github.com/PCPedroso/pos-fc-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _user = entity.User{
	Name:     "Jose Joarez",
	Email:    "jj@gmail.com",
	Password: "654641",
}

func ConectaDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{})

	return db, nil
}

func TestCreateUser(t *testing.T) {
	db, err := ConectaDB()
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser(_user.Name, _user.Email, _user.Password)
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
}

func TestFindByEmail(t *testing.T) {
	db, err := ConectaDB()
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser(_user.Name, _user.Email, _user.Password)
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
