package services

type MockAPIKeyService struct{}

func (m *MockAPIKeyService) GenerateAndSaveAPIKey() (string, error) {
	return "mocked-api-key", nil
}
