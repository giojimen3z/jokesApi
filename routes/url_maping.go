package routes

import (
	"github.com/gin-gonic/gin"

	"awesomeProject1/controller"
)

func MapUrl(router *gin.Engine) {
	router.GET("/jokes", controller.GetJokes())
}
