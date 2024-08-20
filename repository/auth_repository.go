package repository

import (
	"database/sql"
	// "github.com/google/uuid"
	"golang-book-api/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo *AuthRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (id, username, password, email) VALUES (?, ?, ?, ?)"
	_, err := repo.db.Exec(query, user.ID, user.Username, user.Password, user.Email)
	return err
}

func (repo *AuthRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password, email FROM users WHERE username = ?"
	err := repo.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
