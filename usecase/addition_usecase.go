package usecase

import (
	"fmt"

	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/repository"
)

type AdditionUsecase struct {
	additionRepository *repository.AdditionRepository
}

func NewAdditionUsecase(additionRepository *repository.AdditionRepository) *AdditionUsecase {
	return &AdditionUsecase{
		additionRepository: additionRepository,
	}
}

func (tu *AdditionUsecase) UpsertFavo(favo model.Addition) error {
	err := tu.additionRepository.UpsertFavo(favo)
	if err != nil {
		fmt.Println("update to Favo :", err)
	}
	return err
}
