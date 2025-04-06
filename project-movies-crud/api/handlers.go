package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/learn-golang/project-movies-crud/entities"
	"github.com/learn-golang/project-movies-crud/service"
)

type Handler struct {
	service service.MoviesService
}

func NewHandler(service service.MoviesService) *Handler {
	return &Handler{
		service: service,
	}
}

// @Summary      Get all movies
// @Description  Retrieves a list of all movies from the database
// @Tags         movies
// @Produce      json
// @Success      200  {array}   entities.Movie
// @Failure      500  {object}  map[string]string
// @Router       /movies [get]
func (h *Handler) getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	movies, err := h.service.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(movies)
}

// @Summary      Get movie by ID
// @Description  Retrieves a single movie by its ID
// @Tags         movies
// @Produce      json
// @Param        id   path      string  true  "Movie ID"
// @Success      200  {object}  entities.Movie
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /movies/{id} [get]
func (h *Handler) getMovieByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	movie, err := h.service.GetMovieById(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

// @Summary      Create a new movie
// @Description  Adds a new movie to the database
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movie  body      entities.Movie         true  "Movie object to create"
// @Success      201    {object}  entities.SuccessResponse
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /movies [post]
func (h *Handler) creatMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie entities.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := entities.SuccessResponse{
		Status: true,
		Data:   &movie,
	}

	json.NewEncoder(w).Encode(response)

}

// @Summary      Update a movie
// @Description  Updates an existing movie's details
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movie  body      entities.Movie         true  "Movie object to update"
// @Success      200    {object}  entities.SuccessResponse
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /movies/{id} [put]
func (h *Handler) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie entities.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateMovie(movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := entities.SuccessResponse{
		Status: true,
		Data:   &movie,
	}

	json.NewEncoder(w).Encode(response)
}

// @Summary Delete a movie
// @Description Delete a movie by ID
// @Tags movies
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} entities.SuccessResponse
// @Failure 500 {object} map[string]string
// @Router /movies/{id} [delete]
func (h *Handler) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if err := h.service.DeleteMovieById(params["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := entities.SuccessResponse{
		Status: true,
		Data:   nil,
	}

	json.NewEncoder(w).Encode(response)
}
