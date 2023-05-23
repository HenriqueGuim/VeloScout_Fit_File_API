package controller

import (
	"fitAPI/app/src/services"
	"github.com/gin-gonic/gin"
)

func StartControllers(router *gin.Engine) {
	router.POST("/decode", services.DecodeFitFile)
}
