package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/handler"
	"github.com/snrnapa/todo-everyone-go-back/middlewares"
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

	todoHandler := handler.NewTodoHandler(
		usecase.NewTodoUsecase(
			repository.NewTodoRepository(),
		),
	)

	commentHandler := handler.NewCommentHandler(
		usecase.NewCommentUsecase(
			repository.NewCommentRepository(),
		),
	)

	r.GET("/users", userHandler.GetUsers)
	// r.GET("/user", userHandler.GetUser)
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	protected := r.Group("/v1")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", userHandler.GetUser)
	protected.GET("/current-user", userHandler.FindCurrentUser)

	// todo information
	protected.GET("/todos", todoHandler.GetTodos)
	protected.POST("/todo", todoHandler.InsertTodo)
	protected.DELETE("/todo", todoHandler.DeleteTodo)
	protected.POST("/comment", commentHandler.InsertComment)

	r.Run()

}
