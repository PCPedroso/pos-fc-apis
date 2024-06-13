package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jose Joarez", "jj@gmail.com", "654641")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jose Joarez", user.Name)
	assert.Equal(t, "jj@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jose Joarez", "jj@gmail.com", "654641")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("654641"))
	assert.False(t, user.ValidatePassword("8974561"))
	assert.NotEqual(t, "654641", user.Password)
}
