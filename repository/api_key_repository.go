package repository

import (
	"database/sql"
)

type APIKeyRepository struct {
	db *sql.DB
}

func NewAPIKeyRepository(db *sql.DB) *APIKeyRepository {
	return &APIKeyRepository{db}
}

func (r *APIKeyRepository) SaveAPIKeyToDB(apiKey string) error {
	_, err := r.db.Exec("INSERT INTO api_keys(api_key) VALUES(?)", apiKey)
	return err
}

func (r *APIKeyRepository) IsAPIKeyValid(apiKey string) (bool, error) {
	query := "SELECT COUNT(*) FROM api_keys WHERE api_key = ?"
	var count int
	err := r.db.QueryRow(query, apiKey).Scan(&count)
	if err != nil || count == 0 {
		return false, err
	}
	return true, nil
}
