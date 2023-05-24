package controller

import (
	"fitAPI/app/src/services"
	"github.com/gin-gonic/gin"
)

// StartControllers starts the controllers
// @Summary Decode a .fit file
// @Description Decode a .fit file
// @Param FitFile formData file as "File to decode"
// @Produce application/json
// @Success 200 {object} models.DataFormat
// @Router /decode [post]
func StartControllers(router *gin.Engine) {
	router.POST("/decode", services.DecodeFitFile)
}
