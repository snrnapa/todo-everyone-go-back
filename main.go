package main

import (
	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/handler"
	"github.com/snrnapa/todo-everyone-go-back/repository"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
)

func main() {

	dsn := "host=localhost user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db.Init(dsn)
	db.CreateInitData()

	r := gin.Default()

	userHandler := handler.NewUserHandler(
		usecase.NewUserUsecase(
			repository.NewUserRepository(),
		),
	)

	r.GET("/users", userHandler.GetUsers)

	// group化するときの記述
	// v1 := r.Group("/v1")
	// {
	// 	v1.GET("/home", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"message": "hogehogesigeru",
	// 		})
	// 	})
	// }
	r.Run()
}
