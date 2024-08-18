package controllers

import (
	"golang-book-api/services" // Import services
	"net/http"
)

type APIKeyController struct {
	APIKeyService *services.APIKeyService
}

func NewAPIKeyController(apiKeyService *services.APIKeyService) *APIKeyController {
	return &APIKeyController{APIKeyService: apiKeyService}
}

func (c *APIKeyController) GetAPIKeyHandler(w http.ResponseWriter, r *http.Request) {
	apiKey, err := c.APIKeyService.GenerateAndSaveAPIKey()
	if err != nil {
		http.Error(w, "Could not generate and save API key", http.StatusInternalServerError)
		return
	}

	// Mengembalikan API key sebagai respons
	w.Write([]byte(apiKey))
}
