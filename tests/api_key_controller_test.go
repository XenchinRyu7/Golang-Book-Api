package tests

import (
	"golang-book-api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockAPIKeyService struct{}

func (m *MockAPIKeyService) GenerateAndSaveAPIKey() (string, error) {
	return "mocked-api-key", nil
}

func TestGetAPIKeyHandler(t *testing.T) {
	mockService := &MockAPIKeyService{}
	controller := controllers.NewAPIKeyController(mockService)

	req, err := http.NewRequest("GET", "/api-key", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAPIKeyHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "mocked-api-key"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
