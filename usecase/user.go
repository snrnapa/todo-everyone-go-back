package usecase

import (
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/repository"
)

type UserUsecase struct {
	userRepository *repository.UserRepository
}

func NewUserUsecase(userRepository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (uc *UserUsecase) GetUsers() []model.MstUser {
	response := uc.userRepository.GetUsers()
	return response
}
