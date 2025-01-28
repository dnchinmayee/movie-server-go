package repositories

import (
	"log"
	"movie-server/models"
	"strings"
)

type MovieRepository interface {
	GetAll() []models.Movie
	GetById(id int) models.Movie
	Create(movie models.Movie) models.Movie
	Update(movie models.Movie) models.Movie
	Delete(id int) bool
	SearchMoviesByTitle(title string) []models.Movie
	SearchMoviesByTitleAndDirector(title, director string) []models.Movie
}

type MovieMemoryRepository struct {
	movies []models.Movie
}

func NewMovieMemoryRepository() *MovieMemoryRepository {
	repo := &MovieMemoryRepository{}
	// Mock data
	movies := []models.Movie{
		{ID: 1, Title: "Inception", Director: "Christopher Nolan"},
		{ID: 2, Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski"},
	}
	repo.movies = movies
	return repo
}

func (r *MovieMemoryRepository) GetAll() []models.Movie {
	return r.movies
}

func (r *MovieMemoryRepository) GetById(id int) models.Movie {
	for _, m := range r.movies {
		if m.ID == id {
			return m
		}
	}
	return models.Movie{}
}

func (r *MovieMemoryRepository) Create(movie models.Movie) models.Movie {
	movie.ID = r.movies[len(r.movies)-1].ID + 1
	r.movies = append(r.movies, movie)
	return movie
}

func (r *MovieMemoryRepository) Update(movie models.Movie) models.Movie {
	for i, m := range r.movies {
		if m.ID == movie.ID {
			r.movies[i] = movie
			return movie
		}
	}
	return models.Movie{}
}

func (r *MovieMemoryRepository) Delete(id int) bool {
	for i, m := range r.movies {
		if m.ID == id {
			r.movies = append(r.movies[:i], r.movies[i+1:]...)
			return true
		}
	}
	return false
}

func (r *MovieMemoryRepository) SearchMoviesByTitle(title string) []models.Movie {
	// movies := ctrl.repository.GetAll()
	var result []models.Movie

	//find range of movies and seach titile by separated by space or the whole title name
	for _, m := range r.movies {
		if title == "" {
			result = append(result, m)
		} else {
			// if m.Title == title {
			if strings.Contains(m.Title, title) {
				result = append(result, m)
			}
		} // SearchMoviesByTitle searches movies by title
	}
	return result
}

func (r *MovieMemoryRepository) SearchMoviesByTitleAndDirector(title, director string) []models.Movie {
	log.Printf("Entering SearchMoviesByTitleAndDirector")
	/*
		if only title or director is available then search by that specific term
		if both present, then search movies which have both title & director
		if none are present then send all movies

	*/

	if title == "" && director == "" {
		return r.movies
	}

	if title != "" && director == "" {
		return r.SearchMoviesByTitle(title)
		// return movies with matching title
	}

	if title == "" && director != "" {
		// retrun movies with matching director
		return r.SearchMovieDirector(director)
	}

	// return movies matching both title and director

	return r.SearchMovieByTitleAndDirector(title, director)
}

func (r *MovieMemoryRepository) SearchMovieDirector(dir string) []models.Movie {
	var result []models.Movie

	for _, m := range r.movies {
		if strings.Contains(m.Director, dir) {
			result = append(result, m)
		}
	}

	return result
}

func (r *MovieMemoryRepository) SearchMovieByTitleAndDirector(title, dir string) []models.Movie {
	var result []models.Movie

	for _, m := range r.movies {
		if strings.Contains(m.Director, dir) && strings.Contains(m.Title, title) {
			result = append(result, m)
		}
	}

	return result
}
