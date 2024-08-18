package services

import (
	"crypto/rand"
	"encoding/hex"
	"golang-book-api/repository"
)

type APIKeyService struct {
	APIKeyRepo *repository.APIKeyRepository
}

func NewAPIKeyService(apiKeyRepo *repository.APIKeyRepository) *APIKeyService {
	return &APIKeyService{APIKeyRepo: apiKeyRepo}
}

func (s *APIKeyService) GenerateAndSaveAPIKey() (string, error) {
	apiKey, err := generateAPIKey()
	if err != nil {
		return "", err
	}

	err = s.APIKeyRepo.SaveAPIKeyToDB(apiKey)
	if err != nil {
		return "", err
	}

	return apiKey, nil
}

func generateAPIKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(key), nil
}
