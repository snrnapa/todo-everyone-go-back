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

	// 環境変数で環境を判別
	env := os.Getenv("ENV")
	var logDir string

	if env == "docker" {
		logDir = "/app/log"
	} else {
		logDir = filepath.Join(currentDir, "log")
	}

	// ログディレクトリが存在しない場合は作成する
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			log.Fatalf("error creating log directory: %v", err)
		}
		log.Printf("created log directory: %s", logDir)
	} else {
		log.Printf("log directory already exists: %s", logDir)
	}

	var logFile *os.File
	setLogFile := func() {
		logFileName := "app-" + time.Now().Format("2006-01-02") + ".log"
		logFIlePath := filepath.Join(logDir, logFileName)

		if logFile != nil {
			logFile.Close()
		}

		var err error
		logFile, err = os.OpenFile(logFIlePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failled to open log file: %v", err)
		}

		log.SetOutput(logFile)
	}

	setLogFile()
	defer func() {
		if logFile != nil {
			logFile.Close()
			log.Println("Log file closed")
		}
	}()

	credentalFilePath := filepath.Join(currentDir, "serviceAccountKey.json")
	middlewares.InitFirebase(credentalFilePath)

	var dsn string
	if env == "docker" {
		dsn = "host=db user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	} else {
		dsn = "host=localhost user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	}

	db.Init(dsn)
	dbInstance := db.GetDbInstantce()
	// db.CreateInitData()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://todo-everyone.web.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:5173" || origin == "https://todo-everyone.web.app"
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
			repository.NewTodoRepository(dbInstance),
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
