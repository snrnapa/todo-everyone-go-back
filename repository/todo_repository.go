package repository

import (
	"log"

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

func (todoRepo *TodoRepository) GetTodos() ([]TodoWithAdditions, error) {
	query := `
	SELECT
		t.id
		, t.user_id
		, t.title
		, t.detail
		, t.deadline
		, t.completed
		, COUNT(a.is_favorite) FILTER(WHERE a.is_favorite = TRUE) AS favorite_count
		, COUNT(a.is_booked) FILTER(WHERE a.is_booked = TRUE) AS booked_count
		, COUNT(a.is_cheered) FILTER(WHERE a.is_cheered = TRUE) AS cheered_count
		, bool_or(a.is_favorite) as is_favorite_me
		, bool_or(a.is_booked) as is_booked_me
		, bool_or(a.is_cheered) as is_cheered_me 
	FROM
		todos t 
		LEFT JOIN ( 
			SELECT
				a.todo_id
				, a.is_favorite
				, a.is_booked
				, a.is_cheered 
			FROM
				trn_additions a
		) a 
			ON t.id = a.todo_id 
	GROUP BY
		t.id
		, t.user_id 
	order by
		t.deadline;
	`

	var todoWithAdditions []TodoWithAdditions
	result := todoRepo.Database.Raw(query).Scan(&todoWithAdditions)
	if result.Error != nil {
		log.Printf("query execution failed: %v", result.Error)
		return nil, result.Error
	}
	return todoWithAdditions, result.Error
}

func (todoRepo *TodoRepository) InsertTodo(todo model.Todo) (model.Todo, error) {
	result := todoRepo.Database.Save(&todo)
	return todo, result.Error
}

func (todoRepo *TodoRepository) DeleteTodo(id uint) error {
	var todo model.Todo
	result := todoRepo.Database.Unscoped().Delete(&todo, id)
	return result.Error
}

func (todoRepo *TodoRepository) UpdateTodo(todo model.Todo) error {
	if err := todoRepo.Database.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

type TodoWithAdditions struct {
	ID            int64  `json:"id"`
	UserID        string `json:"user_id"`
	Title         string `json:"title"`
	Detail        string `json:"detail"`
	Deadline      string `json:"deadline"` // Adjust the type if needed
	Completed     bool   `json:"completed"`
	FavoriteCount int    `json:"favorite_count"`
	BookedCount   int    `json:"booked_count"`
	CheeredCount  int    `json:"cheered_count"`
	IsFavoriteMe  bool   `json:"is_favorite_me"`
	IsBookedMe    bool   `json:"is_booked_me"`
	IsCheeredMe   bool   `json:"is_cheered_me"`
}
