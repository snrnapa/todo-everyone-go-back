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

func (ur *UserRepository) GetUserById(userId string) (model.User, error) {
	var user model.User
	result := ur.Database.Where("user_id = ?", userId).Find(&user)
	return user, result.Error
}

func (ur *UserRepository) Register(userId string) (string, error) {
	targetUser := model.User{
		UserId: userId,
	}

	err := ur.Database.Save(&targetUser).Error
	if err != nil {
		return userId, err
	}
	return userId, err
}

func (ur *UserRepository) InsertContact(contactInfo model.Contact) error {
	err := ur.Database.Save(&contactInfo).Error
	if err != nil {
		return err
	}
	return nil
}
