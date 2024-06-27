package handler

import (
	"fmt"
	"log"
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
		errMsg := fmt.Sprintf("ユーザー登録処理中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"validation Error": errMsg})
		return
	}

	if err := uh.userUsecase.Register(registerInput.UserId); err != nil {
		errMsg := fmt.Sprintf("ユーザー登録処理中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uh *AuthHandler) Login(c *gin.Context) {

	var registerInput RegisterInput
	if err := c.BindJSON(&registerInput); err != nil {
		errMsg := fmt.Sprintf("ログイン処理中にサーバーでエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}
	uh.userUsecase.GetUserById(registerInput.UserId)
}
