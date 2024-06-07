package usecase

import (
	"fmt"

	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/repository"
)

type TodoUsecase struct {
	todoRepository *repository.TodoRepository
}

func NewTodoUsecase(todoRepository *repository.TodoRepository) *TodoUsecase {
	return &TodoUsecase{
		todoRepository: todoRepository,
	}
}

func (tu *TodoUsecase) GetTodos() ([]repository.TodoWithAdditions, error) {
	response, err := tu.todoRepository.GetTodos()
	if err != nil {
		fmt.Println("failed to GetTodos :", err)
	}
	return response, err
}

func (tu *TodoUsecase) GetTodoById(id string) ([]repository.TodoWithAdditions, error) {
	response, err := tu.todoRepository.GetTodoById(id)
	if err != nil {
		fmt.Println("failed to GetTodos :", err)
	}
	return response, err
}

func (tu *TodoUsecase) InsertTodo(todo model.Todo) (model.Todo, error) {
	response, err := tu.todoRepository.InsertTodo(todo)
	if err != nil {
		return response, err
	}
	return response, err
}

func (tu *TodoUsecase) DeleteTodo(id uint) error {
	err := tu.todoRepository.DeleteTodo(id)
	if err != nil {
		fmt.Println("failed to InsertTodo :", err)
	}
	return err
}

func (tu *TodoUsecase) UpdateTodo(todo model.Todo) error {
	err := tu.todoRepository.UpdateTodo(todo)
	if err != nil {
		fmt.Println("update to Todo :", err)
	}
	return err
}

// func (uc *TodoUsecase) GetUser(email string) (model.User, error) {
// 	response, err := uc.userRepository.GetUser(email)
// 	if err != nil {
// 		fmt.Println("failed to GetUser :", err)
// 	}
// 	return response, err
// }
