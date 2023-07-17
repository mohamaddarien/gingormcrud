package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect("postgresql://postgres:root@localhost:5432/albums")
	database.Migrate()

	r := gin.Default()
	r.GET("/albums", controllers.GetAlbums)
	r.GET("/albums/:id", controllers.GetAlbumByID)
	r.POST("/albums", controllers.PostAlbums)
	r.PATCH("/albums/:id", controllers.UpdateAlbum)
	r.DELETE("/albums/:id", controllers.DeleteAlbum)
	r.Run() 
}
