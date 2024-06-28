package handler

import (
	"fmt"
	"log"
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
		errMsg := fmt.Sprintf("サーバーでTodoの取得中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, todos)
}

func (th *TodoHandler) GetSummary(c *gin.Context) {

	userId := c.Param("user_id")
	summary, err := th.todoUsecase.GetSummary(userId)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでSummaryの取得中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		fmt.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, summary)
}

func (th *TodoHandler) GetTodoById(c *gin.Context) {
	id := c.Param("id")
	todos, err := th.todoUsecase.GetTodoById(id)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでTodoのIDで検索中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, todos)
}

func (th *TodoHandler) InsertTodo(c *gin.Context) {

	var todo model.Todo
	if err := c.BindJSON(&todo); err != nil {
		errMsg := fmt.Sprintf("サーバーでTodoの入力内容の解析中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	todo, err := th.todoUsecase.InsertTodo(todo)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでTodoの登録中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, todo)
}

func (th *TodoHandler) DeleteTodo(c *gin.Context) {

	var todo model.Todo
	if err := c.BindJSON(&todo); err != nil {
		errMsg := fmt.Sprintf("サーバーで削除対象の解析中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	err := th.todoUsecase.DeleteTodo(todo.ID)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでTodoの削除中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, todo)
}

func (th *TodoHandler) UpdateTodo(c *gin.Context) {

	var todo model.Todo
	if err := c.BindJSON(&todo); err != nil {
		errMsg := fmt.Sprintf("サーバーで更新対象のTodo解析中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	err := th.todoUsecase.UpdateTodo(todo)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでTodoの更新中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, todo)
}
