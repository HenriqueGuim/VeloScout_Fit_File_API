package main

import (
	"fitAPI/app/src/controller"
	_ "fitAPI/app/src/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title VeloScoutFitAPI Swagger API Documentation
// @version 1.0
// @description This is an API for decoding .fit files
// @host veloscout-henriqueguim0nrdln66m6.koyeb.app
// @BasePath /
func main() {
	router := gin.Default()
	controller.StartControllers(router)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = router.Run() // listen and serve on 0.0.0.0:8080
}
