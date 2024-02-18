// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"music-api/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)

	router.Run("localhost:8080")
}
