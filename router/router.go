package router

import (
	"github.com/gorilla/mux"
	"golang-book-api/controllers"
	"golang-book-api/middleware"
	// "net/http"
)

func SetupRouter(bookController *controllers.BookController) *mux.Router {
	r := mux.NewRouter()

	r.Use(middleware.AuthMiddleware)

	r.HandleFunc("/books", bookController.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
	r.HandleFunc("/books", bookController.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")

	return r
}
