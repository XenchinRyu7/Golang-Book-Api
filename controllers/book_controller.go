package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-book-api/helpers"
	"golang-book-api/models"
	"golang-book-api/services"
	"net/http"
	"strconv"
)

type BookController struct {
	Service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{Service: service}
}

func (bc *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bc.Service.GetBooks()
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helpers.RespondJSON(w, http.StatusOK, books)
}

func (bc *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	book, err := bc.Service.GetBookByID(id)
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}
	helpers.RespondJSON(w, http.StatusOK, book)
}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := bc.Service.CreateBook(book); err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	// helpers.RespondJSON(w, http.StatusCreated, book) respon with object
	helpers.RespondSuccess(w, http.StatusCreated, "Book created successfully")
}

func (bc *BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid Book Id")
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.ID = id

	err = bc.Service.UpdateBook(book, id)
	if err != nil {
		if err.Error() == "book not found" {
			helpers.RespondError(w, http.StatusNotFound, "Book not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "Could not update book")
		return
	}

	// helpers.RespondJSON(w, http.StatusOK, book)
	helpers.RespondSuccess(w, http.StatusCreated, "Book updated successfully")

}

func (bc *BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid Book Id")
		return
	}

	err = bc.Service.DeleteBook(id)
	if err != nil {
		if err.Error() == "book not found" {
			helpers.RespondError(w, http.StatusNotFound, "Book not found")
			return
		}
		helpers.RespondError(w, http.StatusInternalServerError, "Could not delete book")
		return
	}

	// helpers.RespondJSON(w, http.StatusNoContent, nil)
	helpers.RespondSuccess(w, http.StatusOK, "Book deleted successfully")
}
