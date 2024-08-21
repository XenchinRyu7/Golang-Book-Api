package repository

import (
	"database/sql"
	"golang-book-api/models"
	"errors"
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

func (repo *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password, email FROM users WHERE email = ?"
	err := repo.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("email tidak ditemukan")
		}
		return nil, err
	}
	return &user, nil
}
