	package router

	import (
		"github.com/gorilla/handlers"
		"github.com/gorilla/mux"
		"golang-book-api/controllers"
		"golang-book-api/middleware"
		"golang-book-api/repository"
		"os"
		"net/http"
	)

	func SetupRouter(bookController *controllers.BookController, apiKeyController *controllers.APIKeyController, apiKeyRepo *repository.APIKeyRepository) *mux.Router {
		r := mux.NewRouter()

		// Middleware CORS
		r.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

		// Middleware Logging
		r.Use(func(next http.Handler) http.Handler {
			return handlers.LoggingHandler(os.Stdout, next)
		})

		// Route untuk mendapatkan API key tanpa middleware autentikasi
		r.HandleFunc("/get-api-key", apiKeyController.GetAPIKeyHandler).Methods("GET")

		// Routes dengan middleware API Key Authentication
		apiRoutes := r.PathPrefix("/").Subrouter()
		apiRoutes.Use(middleware.APIKeyAuthMiddleware(apiKeyRepo))

		apiRoutes.HandleFunc("/books", bookController.GetBooks).Methods("GET")
		apiRoutes.HandleFunc("/books/{id}", bookController.GetBook).Methods("GET")
		apiRoutes.HandleFunc("/books", bookController.CreateBook).Methods("POST")
		apiRoutes.HandleFunc("/books/{id}", bookController.UpdateBook).Methods("PUT")
		apiRoutes.HandleFunc("/books/{id}", bookController.DeleteBook).Methods("DELETE")

		return r
	}
