package main

import (
	"github.com/gin-gonic/gin"
	"app/handlers"
)

func main() {
	r := gin.Default()
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})*/
	r.GET("/ping", handlers.Ping)
	r.Run(":8081")
}

