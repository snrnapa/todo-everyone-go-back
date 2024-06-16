package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
)

type TodoHandler struct {
	todoUsecase *usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase *usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{
		todoUsecase: todoUsecase,
	}
}

func (th *TodoHandler) GetTodos(c *gin.Context) {
	userId := c.Param("user_id")
	todos, err := th.todoUsecase.GetTodos(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todos)
}

func (th *TodoHandler) GetSummary(c *gin.Context) {

	userId := c.Param("user_id")
	summary, err := th.todoUsecase.GetSummary(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, summary)
}

func (th *TodoHandler) GetTodoById(c *gin.Context) {
	id := c.Param("id")
	todos, err := th.todoUsecase.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todos)
}

func (th *TodoHandler) InsertTodo(c *gin.Context) {

	var todo model.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := th.todoUsecase.InsertTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todo)
}

func (th *TodoHandler) DeleteTodo(c *gin.Context) {

	var todo model.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := th.todoUsecase.DeleteTodo(todo.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todo)
}

func (th *TodoHandler) UpdateTodo(c *gin.Context) {

	var todo model.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := th.todoUsecase.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todo)
}
