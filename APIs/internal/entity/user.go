package entity

import (
	"github.dev/nicolasmmb/GoExpert-Topicos/pkg/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       entity.ID `json:"id" gorm:"type:uuid;primaryKey;not null;uniqueIndex"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"uniqueIndex;not null"`
	Password string    `json:"-" gorm:"not null"`
	gorm.Model
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) ChangeName(name string) {
	if name == "" {
		return
	}
	u.Name = name
}

func (u *User) ChangeEmail(email string) {
	if email == "" {
		return
	}
	u.Email = email
}

func (u *User) ChangePassword(password string) error {
	if password == "" {
		return nil
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}
