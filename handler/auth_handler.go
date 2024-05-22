package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/token"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
	"github.com/snrnapa/todo-everyone-go-back/util"
	"golang.org/x/crypto/bcrypt"
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
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8"`
	RetypePassword string `json:"retype_password" validate:"required,eqfield=Password"`
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	registerUser := model.User{
		Id:       uuid.New().String(),
		Password: string(hashedPassword),
		Email:    registerInput.Email,
	}

	if err := uh.userUsecase.Register(registerUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while registering user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (uh *AuthHandler) Login(c *gin.Context) {

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

func (uh *AuthHandler) FindCurrentUser(c *gin.Context) {
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
