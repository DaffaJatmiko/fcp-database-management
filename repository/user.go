package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	if result := u.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	result := u.db.Table("users").Where("username = ? AND password = ?", user.Username, user.Password).Find(&user)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil // TODO: replace this
}
