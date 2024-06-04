package repository

import (
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/gorm"
)

type AdditionRepository struct {
	Database *gorm.DB
}

func NewAdditionRepository() *AdditionRepository {
	return &AdditionRepository{
		Database: db.GetDbInstantce(),
	}
}

func (additionRepo *AdditionRepository) UpsertFavo(favo model.Addition) error {
	if err := additionRepo.Database.Save(&favo).Error; err != nil {
		return err
	}
	return nil
}
