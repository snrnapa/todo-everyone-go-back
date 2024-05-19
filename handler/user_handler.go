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
	users := uh.userUsecase.GetUsers()
	c.JSON(http.StatusOK, users)
}
