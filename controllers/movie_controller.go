package controllers

import (
	"movie-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	// You can add services here
	Movies []models.Movie
}

func NewMovieController() *MovieController {
	ctrl := &MovieController{}
	// Mock data
	movies := []models.Movie{
		{ID: 1, Title: "Inception", Director: "Christopher Nolan"},
		{ID: 2, Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski"},
	}
	ctrl.Movies = movies

	return ctrl
}

func (ctrl *MovieController) GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.getMovies())
}

func (ctrl *MovieController) GetMovie(c *gin.Context) {
	movieId := c.Param("id")

	// Convert movieId to int
	id, _ := strconv.Atoi(movieId)

	movie := ctrl.getMovieById(id)

	if movie.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func (ctrl *MovieController) CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie = ctrl.createMovie(movie)
	c.JSON(http.StatusOK, movie)
}

func (ctrl *MovieController) UpdateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movieId := c.Param("id")

	// Convert movieId to int
	movie.ID, _ = strconv.Atoi(movieId)
	movie = ctrl.updateMovie(movie)
	c.JSON(http.StatusOK, movie)
}

func (ctrl *MovieController) DeleteMovie(c *gin.Context) {
	movieId := c.Param("id")

	// Convert movieId to int
	id, _ := strconv.Atoi(movieId)

	if ctrl.deleteMovie(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

// handle CRUD operations for movies
func (ctrl *MovieController) getMovies() []models.Movie {
	return ctrl.Movies
}

func (ctrl *MovieController) createMovie(movie models.Movie) models.Movie {
	// Mock response
	movie.ID = ctrl.Movies[len(ctrl.Movies)-1].ID + 1
	ctrl.Movies = append(ctrl.Movies, movie)

	return movie
}

func (ctrl *MovieController) updateMovie(movie models.Movie) models.Movie {
	for i, m := range ctrl.Movies {
		if m.ID == movie.ID {
			ctrl.Movies[i] = movie
			return movie
		}
	}

	return models.Movie{}
}

func (ctrl *MovieController) deleteMovie(id int) bool {
	for i, m := range ctrl.Movies {
		if m.ID == id {
			ctrl.Movies = append(ctrl.Movies[:i], ctrl.Movies[i+1:]...)
			return true
		}
	}

	return false
}

func (ctrl *MovieController) getMovieById(id int) models.Movie {
	for _, m := range ctrl.Movies {
		if m.ID == id {
			return m
		}
	}

	return models.Movie{}
}
