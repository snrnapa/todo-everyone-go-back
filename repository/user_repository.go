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

func (ur *UserRepository) GetUsers() ([]model.MstUser, error) {
	var users []model.MstUser
	result := ur.Database.Find(&users).Order("id desc")
	return users, result.Error
}

func (ur *UserRepository) GetUser(id string) (model.MstUser, error) {
	var user model.MstUser
	result := ur.Database.Where("id = ?", id).Find(&user)
	return user, result.Error
}
