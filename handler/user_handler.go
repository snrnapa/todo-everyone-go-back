package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (uh *UserHandler) GetUsers(c *gin.Context) {
	users, err := uh.userUsecase.GetUsers()
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでユーザー情報の取得中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {

	userId := c.Query("user_id")
	user, err := uh.userUsecase.GetUserById(userId)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーで、ユーザー情報をIDで検索中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) InsertContact(c *gin.Context) {

	var contactInfo model.Contact
	if err := c.BindJSON(&contactInfo); err != nil {
		errMsg := fmt.Sprintf("サーバーで、問い合わせ内容の入力内容を解析中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
	err := uh.userUsecase.InsertContact(contactInfo)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーで、問い合わせ内容の送信中にエラーが発生しました : %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
	c.JSON(http.StatusOK, nil)
}
