package controllers

import (
	"movie-server/models"
	"movie-server/repositories"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	repository repositories.MovieRepository
}

func NewMovieController(repo repositories.MovieRepository) *MovieController {
	return &MovieController{repository: repo}
}

func (ctrl *MovieController) GetMovies(c *gin.Context) {
	c.JSON(http.StatusOK, ctrl.repository.GetAll())
}

func (ctrl *MovieController) GetMovie(c *gin.Context) {
	movieId := c.Param("id")

	// Convert movieId to int
	id, _ := strconv.Atoi(movieId)

	movie := ctrl.repository.GetById(id)

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

	movie = ctrl.repository.Create(movie)
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
	movie = ctrl.repository.Update(movie)
	c.JSON(http.StatusOK, movie)
}

func (ctrl *MovieController) DeleteMovie(c *gin.Context) {
	movieId := c.Param("id")

	// Convert movieId to int
	id, _ := strconv.Atoi(movieId)

	if ctrl.repository.Delete(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Movie deleted"})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
}

func (ctrl *MovieController) SearchMovie(c *gin.Context) {
	title := c.Query("title")
	m := ctrl.repository.SearchMoviesByTitle(title)
	if len(m) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
	} else {
		c.JSON(http.StatusOK, m)
	}
}

// SearchMoviesByTitle function seraches movies by title words

// }
