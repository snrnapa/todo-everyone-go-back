package handler

import (
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
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {

	userId := c.Query("user_id")
	user, err := uh.userUsecase.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) InsertContact(c *gin.Context) {

	var contactInfo model.Contact
	if err := c.BindJSON(&contactInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uh.userUsecase.InsertContact(contactInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)

}
