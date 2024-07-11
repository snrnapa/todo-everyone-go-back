package usecase

import (
	"log"
	"time"

	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/repository"
	"gorm.io/gorm"
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
	var response []repository.TodoWithAdditions
	err := tu.todoRepository.Database.Transaction(func(tx *gorm.DB) error {
		var err error
		response, err = tu.todoRepository.GetTodos(tx, userId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("failed to GetTodos :", err)
		return nil, err
	}
	return response, err
}

func (tu *TodoUsecase) GetSummary(userId string) ([]repository.Summary, error) {
	today := time.Now()
	oneWeekLater := today.AddDate(0, 0, 7)
	dateFormat := "2006-01-02"

	todayString := today.Format(dateFormat)
	oneWeekLaterString := oneWeekLater.Format(dateFormat)
	var response []repository.Summary
	err := tu.todoRepository.Database.Transaction(func(tx *gorm.DB) error {
		var err error
		response, err = tu.todoRepository.GetSummary(tx, userId, todayString, oneWeekLaterString)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("failed to GetSummary :", err)
		return nil, err
	}
	return response, err
}

func (tu *TodoUsecase) GetTodoById(todoId string, userId string) (TodoInfoResponse, error) {
	todo, err := tu.todoRepository.GetTodoById(todoId, userId)
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

func (tu *TodoUsecase) InsertTodo(todo model.Todo) ([]repository.TodoWithAdditions, error) {
	var response []repository.TodoWithAdditions
	userId := todo.UserId

	err := tu.todoRepository.Database.Transaction(func(tx *gorm.DB) error {
		var err error
		response, err = tu.todoRepository.InsertTodo(tx, todo)
		if err != nil {
			return err
		}
		response, err = tu.todoRepository.GetTodos(tx, userId)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("failed to GetSummary :", err)
		return response, err
	}
	return response, err
}

func (tu *TodoUsecase) DeleteTodo(id uint) error {
	err := tu.todoRepository.DeleteTodo(id)
	if err != nil {
		log.Println("failed to InsertTodo :", err)
	}
	return err
}

func (tu *TodoUsecase) UpdateTodo(todo model.Todo) error {
	err := tu.todoRepository.UpdateTodo(todo)
	if err != nil {
		log.Println("update to Todo :", err)
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
