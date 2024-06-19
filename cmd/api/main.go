package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/db"
	"github.com/snrnapa/todo-everyone-go-back/handler"
	"github.com/snrnapa/todo-everyone-go-back/middlewares"
	"github.com/snrnapa/todo-everyone-go-back/repository"
	"github.com/snrnapa/todo-everyone-go-back/usecase"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting current directory: %v", err)
	}
	credentalFilePath := filepath.Join(currentDir, "serviceAccountKey.json")
	middlewares.InitFirebase(credentalFilePath)

	dsn := "host=db user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db.Init(dsn)
	// db.CreateInitData()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://todo-everyone.web.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173"
		},
		MaxAge: 12 * time.Hour,
	}))

	authHandler := handler.NewAuthHandler(
		usecase.NewUserUsecase(
			repository.NewUserRepository(),
		),
	)

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

	additionHandler := handler.NewAdditionHandler(
		usecase.NewAdditionUsecase(
			repository.NewAdditionRepository(),
		),
	)

	commentHandler := handler.NewCommentHandler(
		usecase.NewCommentUsecase(
			repository.NewCommentRepository(),
		),
	)

	protected := r.Group("/v1")
	protected.Use(middlewares.AuthMiddleware())

	// auth Logic
	protected.POST("/register", authHandler.Register)

	// user Information
	protected.GET("/user", userHandler.GetUserById)
	protected.POST("/contact", userHandler.InsertContact)

	// todo information
	protected.GET("/todos/:user_id", todoHandler.GetTodos)
	protected.GET("/summary/:user_id", todoHandler.GetSummary)
	protected.GET("/todo/:id", todoHandler.GetTodoById)
	protected.POST("/todo", todoHandler.InsertTodo)
	protected.DELETE("/todo", todoHandler.DeleteTodo)
	protected.PATCH("/todo", todoHandler.UpdateTodo)

	// addition information for todo
	protected.POST("/addition", additionHandler.UpsertAddition)
	// protected.POST("/book", additionHandler.UpdateBook)

	// comment Information
	protected.POST("/comment", commentHandler.InsertComment)

	r.Run()

}
