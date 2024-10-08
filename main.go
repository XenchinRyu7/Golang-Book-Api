package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang-book-api/config"
	"golang-book-api/controllers"
	// "golang-book-api/middleware"
	"golang-book-api/repository"
	"golang-book-api/router"
	"golang-book-api/services"
	"log"
	"net/http"
)

func seedData(db *sql.DB) {
	_, err := db.Exec(`INSERT INTO books (isbn, title, author, publisher, year)
        VALUES
        ('978-3-16-148410-0', 'Book Title 1', 'Author 1', 'Publisher 1', 2020),
        ('978-1-56619-909-4', 'Book Title 2', 'Author 2', 'Publisher 2', 2021);`)
	if err != nil {
		log.Fatal("Could not seed data", err)
	}

	log.Println("Seed data applied successfully")
}

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}
	defer db.Close()

	bookRepo := repository.NewBookRepository(db)
	apiKeyRepo := repository.NewAPIKeyRepository(db)

	bookService := services.NewBookService(bookRepo)
	apiKeyService := services.NewAPIKeyService(apiKeyRepo)

	bookController := controllers.NewBookController(bookService)
	apiKeyController := controllers.NewAPIKeyController(apiKeyService)

	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	// Migrate Database
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Could not create migration driver", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql", driver)
	if err != nil {
		log.Fatal("Could not start migration", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Could not apply migration", err)
	}

	// Seed Data
	seedData(db)

	// Setup Router
	r := router.SetupRouter(bookController, apiKeyController, apiKeyRepo, authController)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
