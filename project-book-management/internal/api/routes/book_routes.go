package routes

import (
	"github.com/Ritika-Agrawal811/learn-golang/project-book-management/internal/api/handlers"
	"github.com/Ritika-Agrawal811/learn-golang/project-book-management/internal/services"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)
	r := mux.NewRouter()

	r.HandleFunc("/books", bookHandler.CreateBookHandler).Methods("POST")
	r.HandleFunc("/books", bookHandler.GetAllBooksHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.GetBookByIdHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.UpdateBookByIdHandler).Methods("PUT")
	r.HandleFunc("/books/{id}", bookHandler.DeleteBookByIdHandler).Methods("DELETE")
	return r
}
