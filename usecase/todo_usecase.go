package usecase

import (
	"fmt"
	"time"

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

func (tu *TodoUsecase) GetTodos(userId string) ([]repository.TodoWithAdditions, error) {

	response, err := tu.todoRepository.GetTodos(userId)
	if err != nil {
		fmt.Println("failed to GetTodos :", err)
	}
	return response, err
}

func (tu *TodoUsecase) GetSummary(userId string) ([]repository.Summary, error) {
	today := time.Now()
	oneWeekLater := today.AddDate(0, 0, 7)
	dateFormat := "2006-01-02"

	todayString := today.Format(dateFormat)
	oneWeekLaterString := oneWeekLater.Format(dateFormat)

	response, err := tu.todoRepository.GetSummary(userId, todayString, oneWeekLaterString)
	if err != nil {
		fmt.Println("failed to GetSummary :", err)
	}
	return response, err
}

func (tu *TodoUsecase) GetTodoById(todoId string) (TodoInfoResponse, error) {
	todo, err := tu.todoRepository.GetTodoById(todoId)
	if err != nil {
		return TodoInfoResponse{}, err
	}
	comment, err := tu.todoRepository.GetCommentByTodoId(todoId)
	if err != nil {
		return TodoInfoResponse{}, err
	}

	result := TodoInfoResponse{
		ID:            todo.ID,
		UserID:        todo.UserID,
		Title:         todo.Title,
		Detail:        todo.Detail,
		Deadline:      todo.Deadline,
		Completed:     todo.Completed,
		FavoriteCount: todo.FavoriteCount,
		BookedCount:   todo.BookedCount,
		CheeredCount:  todo.CheeredCount,
		CommentCount:  len(comment),
		IsFavoriteMe:  todo.IsFavoriteMe,
		IsBookedMe:    todo.IsBookedMe,
		IsCheeredMe:   todo.IsCheeredMe,
		Comments:      comment,
	}

	return result, err
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

type TodoInfoResponse struct {
	ID            int64           `json:"id"`
	UserID        string          `json:"user_id"`
	Title         string          `json:"title"`
	Detail        string          `json:"detail"`
	Deadline      string          `json:"deadline"`
	Completed     bool            `json:"completed"`
	FavoriteCount int             `json:"favorite_count"`
	BookedCount   int             `json:"booked_count"`
	CheeredCount  int             `json:"cheered_count"`
	CommentCount  int             `json:"comment_count"`
	IsFavoriteMe  bool            `json:"is_favorite_me"`
	IsBookedMe    bool            `json:"is_booked_me"`
	IsCheeredMe   bool            `json:"is_cheered_me"`
	Comments      []model.Comment `json:"comments"`
}
