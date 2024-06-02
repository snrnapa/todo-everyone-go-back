package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	}
	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserById(c *gin.Context) {

	userId := c.Query("user_id")
	user, err := uh.userUsecase.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, user)
}
