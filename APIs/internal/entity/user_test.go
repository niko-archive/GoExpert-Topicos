package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Nicolas", "niko@mail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Nicolas", user.Name)
	assert.Equal(t, "niko@mail.com", user.Email)
}

func TestComparePassword(t *testing.T) {
	user, err := NewUser("Nicolas", "niko@mail.com", "123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.True(t, user.ComparePassword("123456"))
	assert.False(t, user.ComparePassword("1234567"))

	assert.NotEqual(t, "123456", user.Password)

}
