package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/model"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
	"github.com/snrnapa/todo-everyone-go-back/util"
)

type AdditionHandler struct {
	additionUsecase *usecase.AdditionUsecase
}

func NewAdditionHandler(additionUsecase *usecase.AdditionUsecase) *AdditionHandler {
	return &AdditionHandler{
		additionUsecase: additionUsecase,
	}
}

func (th *AdditionHandler) UpsertAddition(c *gin.Context) {

	var addition model.Addition
	if err := c.BindJSON(&addition); err != nil {
		errMsg := fmt.Sprintf("サーバーでAdd情報の解析中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	if err := util.ValidationCheck(addition); err != nil {
		errMsg := fmt.Sprintf("サーバーでAdd情報の更新中にヴァリデーションエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, gin.H{"validation Error": errMsg})
		return
	}

	err := th.additionUsecase.UpsertAddition(addition)
	if err != nil {
		errMsg := fmt.Sprintf("サーバーでAdd情報の更新中にエラーが発生しました: %v", err.Error())
		log.Println(errMsg)
		c.JSON(http.StatusBadRequest, errMsg)
	}
	c.JSON(http.StatusOK, err)
}
