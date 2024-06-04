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

func (th *AdditionHandler) UpsertFavo(c *gin.Context) {

	var favo model.Addition
	if err := c.BindJSON(&favo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := util.ValidationCheck(favo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation Error": err.Error()})
		return
	}

	err := th.additionUsecase.UpsertFavo(favo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, err)
}
