package handler

import (
	"fmt"
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
	todos, err := th.todoUsecase.GetTodos()
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

	fmt.Println(todo)

	todo, err := th.todoUsecase.InsertTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, todo)
}
