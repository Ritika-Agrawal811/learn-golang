package services

type BookService interface{}

type bookService struct{}

func NewBookService() BookService {
	return &bookService{}
}
