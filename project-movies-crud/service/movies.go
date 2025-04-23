package service

import (
	"fmt"

	"github.com/Ritika-Agrawal811/learn-golang/project-movies-crud/entities"
	"github.com/Ritika-Agrawal811/learn-golang/project-movies-crud/utils"
)

type MoviesService interface {
	GetAllMovies() (entities.MovieList, error)
	GetMovieById(id string) (entities.Movie, error)
	DeleteMovieById(id string) error
	UpdateMovie(id string, movie entities.Movie) error
	CreateMovie(movie entities.Movie) error
}

type moviesService struct{}

const (
	filePath = "data/movies.json"
)

func NewMoviesService() MoviesService {
	return &moviesService{}
}

func (s *moviesService) GetAllMovies() (entities.MovieList, error) {
	/* load the movies data from json file */
	movies, err := utils.LoadFromJSON[entities.MovieList](filePath)
	if err != nil {
		return entities.MovieList{}, err
	}

	return movies, nil
}

func (s *moviesService) GetMovieById(id string) (entities.Movie, error) {

	moviesList, err := s.GetAllMovies()
	if err != nil {
		return entities.Movie{}, err
	}

	for _, movie := range moviesList.Movies {
		if movie.Id == id {
			return movie, nil
		}
	}

	return entities.Movie{}, fmt.Errorf("no movie found for id: %s", id)
}

func (s *moviesService) DeleteMovieById(id string) error {
	moviesList, err := s.GetAllMovies()
	if err != nil {
		return err
	}

	for index, movie := range moviesList.Movies {
		if movie.Id == id {
			moviesList.Movies = append(moviesList.Movies[:index], moviesList.Movies[index+1:]...)
		}
	}

	/* update the new movies list with deleted value in json */
	if err := utils.SaveToJSON[entities.MovieList](filePath, moviesList); err != nil {
		return err
	}

	return nil
}

func (s *moviesService) UpdateMovie(id string, movie entities.Movie) error {
	moviesList, err := s.GetAllMovies()
	if err != nil {
		return err
	}

	for index, item := range moviesList.Movies {
		if item.Id == id {
			updatedMovie := entities.Movie{
				Id:       id,
				Isbn:     movie.Isbn,
				Title:    movie.Title,
				Director: movie.Director,
			}
			moviesList.Movies[index] = updatedMovie
		}
	}

	/* update the new movies list in json */
	if err := utils.SaveToJSON[entities.MovieList](filePath, moviesList); err != nil {
		return err
	}

	return nil
}

func (s *moviesService) CreateMovie(movie entities.Movie) error {
	moviesList, err := s.GetAllMovies()
	if err != nil {
		return err
	}

	moviesList.Movies = append(moviesList.Movies, movie)

	/* update the new movies list  in json */
	if err := utils.SaveToJSON[entities.MovieList](filePath, moviesList); err != nil {
		return err
	}

	return nil
}
