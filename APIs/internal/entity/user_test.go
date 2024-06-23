package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("any name", "any email", "any pass")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "any name", user.Name)
	assert.Equal(t, "any email", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("any-name", "any-email", "any-pass")
	assert.Nil(t, err)
	assert.True(t, user.ValidadePassword("any-pass"))
	assert.False(t, user.ValidadePassword("foo"))
	assert.NotEqual(t, "any-pass", user.Password)
}
