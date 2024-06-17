package db

import (
	"fmt"
	"sync"

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
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("failed to connecting database : %v", err))
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
