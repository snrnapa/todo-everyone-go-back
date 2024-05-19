package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/handler"
	"github.com/snrnapa/todo-everyone-go-back/repository"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file", err)

	}

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
	r.GET("/user", userHandler.GetUser)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// r.POST("/login", userHandler.Login)

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
