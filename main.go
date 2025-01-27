package main

import (
	"movie-server/controllers"
	"movie-server/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize repository and controller
	movieRepository := repositories.NewMovieMemoryRepository()
	movieController := controllers.NewMovieController(movieRepository)

	// Movie routes
	router.GET("/movies", movieController.GetMovies)
	router.GET("/movies/:id", movieController.GetMovie)
	router.POST("/movies", movieController.CreateMovie)
	router.PUT("/movies/:id", movieController.UpdateMovie)
	router.DELETE("/movies/:id", movieController.DeleteMovie)
	router.GET("/movies/search", movieController.SearchMovie)

	router.Run(":8080")
}
