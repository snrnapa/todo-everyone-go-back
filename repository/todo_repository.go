package repository

import (
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	Database *gorm.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		Database: db.GetDbInstantce(),
	}
}

func (todoRepo *TodoRepository) GetTodos() ([]model.Todo, error) {
	var todos []model.Todo
	result := todoRepo.Database.Find(&todos).Order("limit")
	return todos, result.Error
}

func (todoRepo *TodoRepository) InsertTodo(todo model.Todo) (model.Todo, error) {
	result := todoRepo.Database.Save(&todo)
	return todo, result.Error
}
