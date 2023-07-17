package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example/web-service-gin/database"
	"example/web-service-gin/entities"
)

func GetAlbums(c *gin.Context) {
    var albums []entities.Album
    database.Instance.Find(&albums)

    c.IndentedJSON(http.StatusOK, gin.H{"data": albums})
}

func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")
    var album entities.Album

    database.Instance.Find(&album, id)

    c.IndentedJSON(http.StatusOK, album)
}

func PostAlbums(c *gin.Context) {
    var newAlbum entities.Album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    database.Instance.Create(&newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}