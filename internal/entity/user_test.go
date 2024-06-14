package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _user = User{
	Name:     "Jose Joarez",
	Email:    "jj@gmail.com",
	Password: "654641",
}

func TestNewUser(t *testing.T) {
	user, err := NewUser(_user.Name, _user.Email, _user.Password)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, _user.Name, user.Name)
	assert.Equal(t, _user.Email, user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser(_user.Name, _user.Email, _user.Password)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(_user.Password))
	assert.False(t, user.ValidatePassword("8974561"))
	assert.NotEqual(t, _user.Password, user.Password)
}
