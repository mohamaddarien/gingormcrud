package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect("postgresql://postgres:root@localhost:5432/albums")
	database.Migrate()

	getAlbums := controllers.GetAlbums
	postAlbums := controllers.PostAlbums
	getAlbumById := controllers.GetAlbumByID

	r := gin.Default()
	r.GET("/albums", getAlbums)
	r.GET("/albums/:id", getAlbumById)
	r.POST("/albums", postAlbums)
	r.Run() 
}
