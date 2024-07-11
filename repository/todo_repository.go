package repository

import (
	"log"

	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/gorm"
)

type TodoRepository struct {
	Database *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		Database: db,
	}
}

func (todoRepo *TodoRepository) GetTodos(tx *gorm.DB, userId string) ([]TodoWithAdditions, error) {
	query := `
	with add_list as ( 
		select
			todos.user_id
			, todos.id
			, bool_or(ad.is_booked) as is_booked_me
			, bool_or(ad.is_cheered) as is_cheered_me 
		from
			todos 
			left join trn_additions ad 
				on todos.id = ad.todo_id 
				where ad.user_id = $1
		group by
			todos.user_id
			, todos.id
	) 
	, comment as ( 
		select
			com.todo_id
			, COUNT(com.user_id) AS comment_count 
		from
			comments com 
		group by
			todo_id
	) 
	, add_count as ( 
		SELECT
			a.todo_id
			, COUNT(a.is_cheered) FILTER(WHERE a.is_cheered = TRUE) AS cheered_count
			, COUNT(a.is_booked) FILTER(WHERE a.is_booked = TRUE) AS booked_count 
		FROM
			trn_additions a 
		GROUP BY
			a.todo_id
	) 
	select
		t.id
		, t.user_id
		, t.title
		, t.detail
		, t.deadline
		, t.completed
		, ac.cheered_count
		, ac.booked_count
		, coalesce(al.is_booked_me, false) as is_booked_me
		, coalesce(al.is_cheered_me, false) as is_cheered_me 
		, com.comment_count
	from
		todos t 
		left join comment com
			on t.id = com.todo_id 
		left join add_count ac 
			on t.id = ac.todo_id 
		left join add_list al 
			on t.id = al.id;	
	`

	var todoWithAdditions []TodoWithAdditions
	if err := tx.Raw(query, userId).Scan(&todoWithAdditions); err != nil {
		log.Printf("query execution failed: %v", err)
		return nil, err.Error
	}

	return todoWithAdditions, nil
}

func (todoRepo *TodoRepository) GetSummary(tx *gorm.DB, id string, today string, oneWeekLater string) ([]Summary, error) {

	query := `
	select
		id
		, user_id
		, title
		, deadline
		, completed 
	from
		todos 
	where
		1 = 1 
		and user_id = $1 
		and deadline between $2 and $3;
	`
	var summary []Summary
	if err := tx.Raw(query, id, today, oneWeekLater).Scan(&summary); err != nil {
		log.Printf("query execution failed: %v", err.Error)
		return summary, err.Error

	}
	return summary, nil
}

type Summary struct {
	ID        int64  `json:"id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Deadline  string `json:"deadline"` // Adjust the type if needed
	Completed bool   `json:"completed"`
}

func (todoRepo *TodoRepository) GetTodoById(id string, userId string) (TodoWithAdditions, error) {
	query := `
	select
		td.id
		, td.user_id
		, td.title
		, td.detail
		, td.deadline
		, td.completed
		, COUNT(ad.is_favorite) FILTER(WHERE ad.is_favorite = TRUE) AS favorite_count
		, COUNT(ad.is_booked) FILTER(WHERE ad.is_booked = TRUE) AS booked_count
		, COUNT(ad.is_cheered) FILTER(WHERE ad.is_cheered = TRUE) AS cheered_count
	from
		todos td
		left join trn_additions ad 
			on td.id = ad.todo_id 
		left join comments cm 
			on td.id = cm.id 
	where
		td.id = $1 
	group by
		td.id
		, td.user_id; `

	var todoWithAdditions TodoWithAdditions
	result := todoRepo.Database.Raw(query, id).Scan(&todoWithAdditions)
	if result.Error != nil {
		log.Printf("query execution failed: %v", result.Error)
		return todoWithAdditions, result.Error
	}

	add_query := `
	select
		is_favorite as is_favorite_me
		, is_booked as is_booked_me
		, is_cheered as is_cheered_me 
	from
		trn_additions t 
	where
		1 = 1 
		and todo_id = $1
		and user_id = $2;`

	result = todoRepo.Database.Raw(add_query, id, userId).Scan(&todoWithAdditions)
	if result.Error != nil {
		log.Printf("query execution failed: %v", result.Error)
		return todoWithAdditions, result.Error
	}
	return todoWithAdditions, result.Error
}

func (todoRepo *TodoRepository) GetCommentByTodoId(todoId string) ([]model.Comment, error) {
	query := `
			select
				* 
			from
				comments 
			where
				todo_id = $1 
			order by
				created_at;`

	var comments []model.Comment
	result := todoRepo.Database.Raw(query, todoId).Scan(&comments)
	if result.Error != nil {
		log.Printf("query execution failed: %v", result.Error)
		return comments, result.Error
	}
	return comments, result.Error
}

func (todoRepo *TodoRepository) InsertTodo(tx *gorm.DB, todo model.Todo) ([]TodoWithAdditions, error) {
	var dummyResponse []TodoWithAdditions
	if err := tx.Save(&todo); err != nil {
		log.Printf("query execution failed: %v", err.Error)
		return dummyResponse, err.Error
	}
	return dummyResponse, nil
}

func (todoRepo *TodoRepository) DeleteTodo(id uint) error {
	var todo model.Todo
	var comment model.Comment
	var additions model.Addition

	tx := todoRepo.Database.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// todosの削除
	if err := tx.Unscoped().Delete(&todo, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commentsの削除
	if err := tx.Unscoped().Where("todo_id = ?", id).Delete(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	// trn_additionsの削除
	if err := tx.Unscoped().Where("todo_id = ?", id).Delete(&additions).Error; err != nil {
		tx.Rollback()
		return err
	}

	// コミットを行う。エラーが出ればロールバック
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
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
	Deadline      string `json:"deadline"`
	Completed     bool   `json:"completed"`
	FavoriteCount int    `json:"favorite_count"`
	BookedCount   int    `json:"booked_count"`
	CheeredCount  int    `json:"cheered_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavoriteMe  bool   `json:"is_favorite_me"`
	IsBookedMe    bool   `json:"is_booked_me"`
	IsCheeredMe   bool   `json:"is_cheered_me"`
}
