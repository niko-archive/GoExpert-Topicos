package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	db.AutoMigrate(&entity.User{})
	user, err := entity.NewUser("Nicolas", "niko@mail.com", "123456")
	assert.NoError(t, err)

	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail("niko@mail.com")
	assert.Nil(t, err)

	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Password, user.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Nicolas", "niko@mail.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)

	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.Equal(t, userFound.Password, user.Password)
}

func TestFindByEmailNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("Nicolas", "niko@mail.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)

	userNotFound, err := userDB.FindByEmail(user.Email + "Error")
	assert.NotNil(t, err)

	assert.Nil(t, userNotFound)
}
