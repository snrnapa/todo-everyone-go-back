package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MstUser struct {
	Id    int
	Name  string
	Age   int
	Email string
	gorm.Model
}

func main() {

	dsn := "host=localhost user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to db open(postgresql)")
	}
	db.AutoMigrate(&MstUser{})
	fmt.Println("completed postgres start and migrate")

	fmt.Println("start createting init user data")

	count := 100
	for i := 0; i < count; i++ {
		countString := strconv.Itoa(i)
		userName := "TestUser" + countString
		userEmail := "testuser" + countString + "@gmail.com"
		age := i
		id := i + 10000

		db.Create(
			&MstUser{
				Id:    id,
				Name:  userName,
				Email: userEmail,
				Age:   age,
			})
	}

	fmt.Println("completed createting init user data")

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hogehogesigeru",
		})
	})

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
