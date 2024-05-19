package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snrnapa/todo-everyone-go-back/db"
)

func main() {

	dsn := "host=localhost user=todo-postgres dbname=todo-postgres password=todo-postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db.Init(dsn)
	db.CreateInitData()

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
