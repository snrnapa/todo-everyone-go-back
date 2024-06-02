package usecase

import (
	"fmt"

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

func (uc *UserUsecase) GetUsers() ([]model.User, error) {
	response, err := uc.userRepository.GetUsers()
	if err != nil {
		fmt.Println("failed to GetUsers :", err)
	}
	return response, err
}

func (uc *UserUsecase) GetUser(email string) (model.User, error) {
	response, err := uc.userRepository.GetUser(email)
	if err != nil {
		fmt.Println("failed to GetUser :", err)
	}
	return response, err
}

func (uc *UserUsecase) GetUserById(userId string) (model.User, error) {
	response, err := uc.userRepository.GetUserById(userId)
	if err != nil {
		fmt.Println("failed to GetUser :", err)
	}
	return response, err
}

func (uc *UserUsecase) Register(userId string) error {
	_, err := uc.userRepository.Register(userId)
	if err != nil {
		fmt.Println("failed to Register :", err)
	}
	return err
}
