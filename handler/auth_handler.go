package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
	"github.com/snrnapa/todo-everyone-go-back/util"
)

type AuthHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewAuthHandler(userUsecase *usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{
		userUsecase: userUsecase,
	}
}

type RegisterInput struct {
	UserId string `json:"user_id" validate:"required"`
}

func (uh *AuthHandler) Register(c *gin.Context) {

	var registerInput RegisterInput
	if err := c.BindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := util.ValidationCheck(registerInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation Error": err.Error()})
		return
	}

	if err := uh.userUsecase.Register(registerInput.UserId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uh *AuthHandler) Login(c *gin.Context) {

	var registerInput RegisterInput
	if err := c.BindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uh.userUsecase.GetUserById(registerInput.UserId)
}
