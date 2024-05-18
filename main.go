package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
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
