package usecase

import (
	"fmt"

	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/repository"
)

type CommentUsecase struct {
	todoRepository *repository.CommentRepository
}

func NewCommentUsecase(todoRepository *repository.CommentRepository) *CommentUsecase {
	return &CommentUsecase{
		todoRepository: todoRepository,
	}
}

// func (tu *CommentUsecase) GetComments() ([]model.Comment, error) {
// 	response, err := tu.todoRepository.GetComments()
// 	if err != nil {
// 		fmt.Println("failed to GetComments :", err)
// 	}
// 	return response, err
// }

func (tu *CommentUsecase) InsertComment(comment model.Comment) (model.Comment, error) {
	response, err := tu.todoRepository.InsertComment(comment)
	if err != nil {
		fmt.Println("failed to InsertComment :", err)
	}
	return response, err
}

// func (tu *CommentUsecase) DeleteComment(id uint) error {
// 	err := tu.todoRepository.DeleteComment(id)
// 	if err != nil {
// 		fmt.Println("failed to InsertComment :", err)
// 	}
// 	return err
// }

// func (uc *CommentUsecase) GetUser(email string) (model.User, error) {
// 	response, err := uc.userRepository.GetUser(email)
// 	if err != nil {
// 		fmt.Println("failed to GetUser :", err)
// 	}
// 	return response, err
// }
