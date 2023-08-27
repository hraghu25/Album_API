package handlers

import (
	"albumapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAlbumsHandler(c *gin.Context) {
	var newAlbum models.Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	models.Albums = append(models.Albums, newAlbum)
	c.JSON(http.StatusOK, newAlbum)
}

func GetAlbumsHandler(c *gin.Context) {
	album := models.Albums
	c.JSON(http.StatusOK, album)
}

func LowPriceAlbum(c *gin.Context) {
	var LowPriceAlbum models.AlbumSlice
	for _, album := range models.Albums {
		if album.Price < 20.0 {
			LowPriceAlbum = append(LowPriceAlbum, album)
		}
	}
	c.JSON(http.StatusOK, LowPriceAlbum)
}
