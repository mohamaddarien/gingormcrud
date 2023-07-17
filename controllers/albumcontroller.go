package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example/web-service-gin/database"
	"example/web-service-gin/entities"
	"example/web-service-gin/requests"
)

func GetAlbums(c *gin.Context) {
    var albums []entities.Album
    database.Instance.Find(&albums)

    c.IndentedJSON(http.StatusOK, gin.H{"data": albums})
}

func GetAlbumByID(c *gin.Context) {    
    var album entities.Album
    if err := database.Instance.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }

    c.IndentedJSON(http.StatusOK, gin.H{"data": album})
}

func PostAlbums(c *gin.Context) {
    tx := database.Instance.Begin()

    var request requests.AlbumRequest

    if err := c.ShouldBindJSON(&request); err != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
      }

    album := entities.Album{Title: request.Artist, Artist: request.Artist, Price: request.Price}
    
    tx.Create(&album)

    if err := tx.Create(&album).Error; err != nil {
        c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to save to Database!"})
        tx.Rollback()
        return
      }

    tx.Commit()
    c.IndentedJSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context){
    var album entities.Album

    if err := database.Instance.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }

    var input requests.AlbumRequest
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.Instance.Model(&album).Updates(input)
    
    c.IndentedJSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbum(c *gin.Context) {
    var album entities.Album
    if err := database.Instance.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
      return
    }
  
    database.Instance.Delete(&album)
  
    c.JSON(http.StatusOK, gin.H{"data": true})
  }
  