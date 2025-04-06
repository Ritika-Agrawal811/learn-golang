package api

import (
	"github.com/gorilla/mux"
	_ "github.com/learn-golang/project-movies-crud/docs"
	"github.com/learn-golang/project-movies-crud/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes() *mux.Router {
	service := service.NewMoviesService()
	handler := NewHandler(service)

	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/movies", handler.getMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", handler.getMovieByIdHandler).Methods("GET")
	r.HandleFunc("/movies", handler.creatMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", handler.updateMovieHandler).Methods("PUT")
	r.HandleFunc("/movies/{id}", handler.deleteMovieHandler).Methods("DELETE")

	return r
}
