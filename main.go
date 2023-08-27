package main

import (
	"albumapi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", handlers.DefaultHandler)
	router.GET("/GetAlbums", handlers.GetAlbumsHandler)
	router.GET("/GetPrice", handlers.LowPriceAlbum)
	router.POST("/AddAlbum", handlers.AddAlbumsHandler)

	router.Run(":8080")
}
