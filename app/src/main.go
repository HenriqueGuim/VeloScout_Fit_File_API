package main

import (
	"fitAPI/app/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	controller.StartControllers(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = router.Run() // listen and serve on 0.0.0.0:8080
}
