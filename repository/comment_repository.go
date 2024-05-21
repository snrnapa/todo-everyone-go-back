package repository

import (
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	Database *gorm.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		Database: db.GetDbInstantce(),
	}
}

// func (todoRepo *CommentRepository) GetTodos() ([]model.Todo, error) {
// 	var todos []model.Todo
// 	result := todoRepo.Database.Find(&todos).Order("limit")
// 	return todos, result.Error
// }

func (todoRepo *CommentRepository) InsertComment(comment model.Comment) (model.Comment, error) {
	result := todoRepo.Database.Save(&comment)
	return comment, result.Error
}

// func (todoRepo *CommentRepository) DeleteTodo(id uint) error {
// 	var todo model.Todo
// 	result := todoRepo.Database.Unscoped().Delete(&todo, id)
// 	return result.Error
// }
