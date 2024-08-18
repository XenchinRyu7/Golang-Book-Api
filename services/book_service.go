package services

import (
	"database/sql"
	"errors"
	"golang-book-api/models"
	"golang-book-api/repository"
)

type BookService interface {
	GetBooks() ([]models.Book, error)
	GetBookByID(id int) (*models.Book, error)
	CreateBook(book models.Book) error
	UpdateBook(book models.Book, id int) error
	DeleteBook(id int) error
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{repo}
}

func (s *bookService) GetBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

func (s *bookService) GetBookByID(id int) (*models.Book, error) {
	return s.repo.FindByID(id)
}

func (s *bookService) CreateBook(book models.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book models.Book, id int) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("book not found")
		}
		return err
	}
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id int) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("book not found")
		}
		return err
	}
	return s.repo.Delete(id)
}
