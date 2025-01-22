package main

import (
	"github.com/gin-gonic/gin"
	"movie-server/controllers"
)

func main() {
	router := gin.Default()

	// Movie routes
	movieController := controllers.NewMovieController()
	router.GET("/movies", movieController.GetMovies)
	router.GET("/movies/:id", movieController.GetMovie)
	router.POST("/movies", movieController.CreateMovie)
	router.PUT("/movies/:id", movieController.UpdateMovie)
	router.DELETE("/movies/:id", movieController.DeleteMovie)

	router.Run(":8080")
}
