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
    var newAlbum entities.Album

    if err := c.ShouldBindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
      }

    database.Instance.Create(&newAlbum)
    c.IndentedJSON(http.StatusOK, gin.H{"data": newAlbum})
}

func UpdateAlbum(c *gin.Context){
    var album entities.Album

    if err := database.Instance.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
      }

    var input requests.UpdateAlbumRequest
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
  