package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/snrnapa/todo-everyone-go-back/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Init(dsn string) {
	once.Do(func() {
		maxRetries := 20
		retryInterval := time.Second * 20

		var err error
		for i := 0; i < maxRetries; i++ {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				log.Printf("Database connected successfully.")
				break
			}
			log.Printf("Failed to connect to database: %v", err)
			if i < maxRetries-1 {
				log.Printf("Retrying in %s...", retryInterval)
				time.Sleep(retryInterval)
			} else {
				panic(fmt.Sprintf("Failed to connect to database after %d attempts: %v", maxRetries, err))
			}
		}

		err = db.AutoMigrate(&model.User{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}

		err = db.AutoMigrate(&model.Contact{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}

		err = db.AutoMigrate(&model.Todo{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}

		err = db.AutoMigrate(&model.Comment{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}

		err = db.AutoMigrate(&model.Addition{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}
	})
}

func GetDbInstantce() *gorm.DB {
	return db
}
