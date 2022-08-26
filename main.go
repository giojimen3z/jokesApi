package main

import (
	"github.com/gin-gonic/gin"

	"awesomeProject1/routes"
)

func main() {
	router := gin.Default()

	routes.MapUrl(router)

	router.Run()

}
