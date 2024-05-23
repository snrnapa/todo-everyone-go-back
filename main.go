package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	dsn := "host=localhost user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db.Init(dsn)
	// db.CreateInitData()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	userHandler := handler.NewUserHandler(
		usecase.NewUserUsecase(
			repository.NewUserRepository(),
		),
	)

	authHandler := handler.NewAuthHandler(
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

	// auth Logic
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/current-user", authHandler.FindCurrentUser)

	protected := r.Group("/v1")
	protected.Use(middlewares.JwtAuthMiddleware())

	// user Information
	protected.GET("/user", userHandler.GetUser)

	// todo information
	protected.GET("/todos", todoHandler.GetTodos)
	protected.POST("/todo", todoHandler.InsertTodo)
	protected.DELETE("/todo", todoHandler.DeleteTodo)

	// comment Information
	protected.POST("/comment", commentHandler.InsertComment)

	r.Run()

}
