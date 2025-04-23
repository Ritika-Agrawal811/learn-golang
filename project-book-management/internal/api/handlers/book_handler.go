package handlers

import (
	"net/http"

	"github.com/Ritika-Agrawal811/learn-golang/project-book-management/internal/services"
)

type BookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{
		bookService,
	}
}

func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) GetBookByIdHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) UpdateBookByIdHandler(w http.ResponseWriter, r *http.Request) {

}

func (h *BookHandler) DeleteBookByIdHandler(w http.ResponseWriter, r *http.Request) {

}
