package db

import (
	"fmt"
	"strconv"
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
	})
}

func GetDbInstantce() *gorm.DB {
	return db
}

func CreateInitData() {
	fmt.Println("start createting init user data")
	var users []model.MstUser
	var count int64

	result := db.Find(&users).Count(&count)

	if result.Error != nil {
		fmt.Println("Error:", result.Error)
	} else {
		fmt.Println("initial date:", count)
	}

	if count == 0 {
		count := 100
		for i := 0; i < count; i++ {
			countString := strconv.Itoa(i)
			userName := "TestUser" + countString
			userEmail := "testuser" + countString + "@gmail.com"
			age := i
			id := i + 10000

			db.Create(
				&model.MstUser{
					Id:    id,
					Name:  userName,
					Email: userEmail,
					Age:   age,
				})
		}
		db.AutoMigrate(&model.MstUser{})
	} else {
		fmt.Println("dont create init data Because you already have", count, "records")

	}

	fmt.Println("completed createting init user data")

}
