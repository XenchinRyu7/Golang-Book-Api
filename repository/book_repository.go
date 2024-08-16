package repository

import (
	"database/sql"
	"golang-book-api/models"
)

type BookRepository interface {
	FindAll() ([]models.Book, error)
	FindByID(id int) (*models.Book, error)
	Create(book models.Book) error
	Update(book models.Book) error
	Delete(id int) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll() ([]models.Book, error) {
	rows, err := r.db.Query("SELECT id, isbn, title, author, publisher, year FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.ISBN, &book.Title, &book.Author, &book.Publisher, &book.Year); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *bookRepository) FindByID(id int) (*models.Book, error) {
	var book models.Book
	err := r.db.QueryRow("SELECT id, isbn, title, author, publisher, year FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.ISBN, &book.Title, &book.Author, &book.Publisher, &book.Year)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepository) Create(book models.Book) error {
	_, err := r.db.Exec("INSERT INTO books(isbn, title, author, publisher, year) VALUES(?, ?, ?, ?, ?)",
		book.ISBN, book.Title, book.Author, book.Publisher, book.Year)
	return err
}

func (r *bookRepository) Update(book models.Book) error {
	_, err := r.db.Exec("UPDATE books SET isbn = ?, title = ?, author = ?, publisher = ?, year = ? WHERE id = ?",
		book.ISBN, book.Title, book.Author, book.Publisher, book.Year, book.ID)
	return err
}

func (r *bookRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
