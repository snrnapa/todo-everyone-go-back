package handler

import (
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := util.ValidationCheck(addition); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation Error": err.Error()})
		return
	}

	err := th.additionUsecase.UpsertAddition(addition)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, err)
}
