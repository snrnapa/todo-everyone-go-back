package repository

import (
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		Database: db.GetDbInstantce(),
	}
}

func (ur *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	result := ur.Database.Find(&users).Order("id desc")
	return users, result.Error
}

func (ur *UserRepository) GetUser(email string) (model.User, error) {
	var user model.User
	result := ur.Database.Where("email = ?", email).Find(&user)
	return user, result.Error
}

func (ur *UserRepository) Register(userCredential model.User) (model.User, error) {
	err := ur.Database.Save(&userCredential).Error
	if err != nil {
		return userCredential, err
	}
	return userCredential, err
}
