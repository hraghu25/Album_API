package main

import (
	"albumapi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.DefaultHandler)

	router.Run(":8080")
}
