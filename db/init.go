package db

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/google/uuid"
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

		err = db.AutoMigrate(&model.Todo{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}

		err = db.AutoMigrate(&model.Comment{})
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}
	})
}

func GetDbInstantce() *gorm.DB {
	return db
}

func CreateInitData() {
	var users []model.User
	var count int64

	result := db.Find(&users).Count(&count)

	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Println("initial date:", count)
	}

	fmt.Println("start createting init user data")
	if count == 0 {
		count := 3
		for i := 0; i < count; i++ {
			countString := strconv.Itoa(i)
			userName := "TestUser" + countString
			userEmail := "testuser" + countString + "@gmail.com"
			age := i
			id := uuid.New().String()

			db.Create(
				&model.User{
					Id:       id,
					Password: "dummypass",
					Name:     userName,
					Email:    userEmail,
					Age:      age,
				})
		}
	} else {
		fmt.Println("dont create init data Because you already have", count, "records")

	}

	fmt.Println("completed createting init user data")

}
