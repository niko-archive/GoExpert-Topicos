package database

import (
	"github.dev/nicolasmmb/GoExpert-Topicos/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindById(id string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) FindAll(page int, limit int, sort string) ([]*entity.User, error) {
	var user []*entity.User
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = u.DB.Limit(limit).Offset((page - 1) * limit).Order("email " + sort).Find(&user).Error
	} else {
		err = u.DB.Find(&entity.Product{}).Error
	}
	return user, err
}

func (u *User) Update(user *entity.User) error {
	return u.DB.Save(user).Error
}

func (u *User) Delete(id string) error {
	return u.DB.Delete(&entity.User{}, id).Error
}
