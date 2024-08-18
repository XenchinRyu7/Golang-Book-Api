package middleware

import (
	"golang-book-api/repository"
	"net/http"
)

func APIKeyAuthMiddleware(apiKeyRepo *repository.APIKeyRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Pengecualian untuk route /get-api-key
			if r.URL.Path == "/get-api-key" {
				next.ServeHTTP(w, r)
				return
			}

			apiKey := r.Header.Get("x-api-key")

			if apiKey == "" {
				http.Error(w, "API key is required", http.StatusUnauthorized)
				return
			}

			isValid, err := apiKeyRepo.IsAPIKeyValid(apiKey)
			if err != nil || !isValid {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
