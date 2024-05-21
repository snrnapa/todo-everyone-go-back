package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/token"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
	"golang.org/x/crypto/bcrypt"
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

func (uh *UserHandler) GetUser(c *gin.Context) {

	email := c.Query("email")
	fmt.Println("input email :", email)
	user, err := uh.userUsecase.GetUser(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) Register(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}
	user.Password = string(hashedPassword)
	user.Id = uuid.New().String()

	if err := uh.userUsecase.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uh *UserHandler) Login(c *gin.Context) {

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Usersテーブルから、Emailをキーとしてユーザー情報を取得
	foundUser, err := uh.userUsecase.GetUser(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(foundUser)

	// Usersテーブルから取得した暗号化したパスワードと、入力されたパスワードを比較
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error CompareHashAndPassword": err.Error()})
		return
	}

	// パスワードが正しければ、tokenを作成する
	token, err := token.GenerateToken(foundUser.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (uh *UserHandler) FindCurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	result, err := uh.userUsecase.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})

}
