package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
)

type CommentHandler struct {
	commentUsecase *usecase.CommentUsecase
}

func NewCommentHandler(commentUsecase *usecase.CommentUsecase) *CommentHandler {
	return &CommentHandler{
		commentUsecase: commentUsecase,
	}
}

// func (th *CommentHandler) GetTodos(c *gin.Context) {
// 	todos, err := th.commentUsecase.GetTodos()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	c.JSON(http.StatusOK, todos)
// }

func (ch *CommentHandler) InsertComment(c *gin.Context) {

	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		errMsg := fmt.Sprintf("サーバーでコメント追加の解析中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	todo, err := ch.commentUsecase.InsertComment(comment)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでコメント追加中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, todo)
}

// func (th *CommentHandler) DeleteTodo(c *gin.Context) {

// 	var todo model.Todo
// 	if err := c.BindJSON(&todo); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	err := th.commentUsecase.DeleteTodo(todo.ID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 	}
// 	c.JSON(http.StatusOK, todo)
// }
