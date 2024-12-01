package handler_test

import (
	"api-golang/internal/domain"
	"api-golang/internal/handler"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do UseCase
type MockCentralUseCase struct {
	mock.Mock
}

func (m *MockCentralUseCase) CreateCentral(central *domain.Central) error {
	args := m.Called(central)
	return args.Error(0)
}

func (m *MockCentralUseCase) GetAllCentrals() ([]domain.Central, error) {
	args := m.Called()
	return args.Get(0).([]domain.Central), args.Error(1)
}

func (m *MockCentralUseCase) GetCentralByID(id uint) (*domain.Central, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Central), args.Error(1)
}

func (m *MockCentralUseCase) UpdateCentral(central *domain.Central) error {
	args := m.Called(central)
	return args.Error(0)
}

func (m *MockCentralUseCase) DeleteCentral(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Função auxiliar para configurar o handler
func setupHandler() (*handler.CentralHandler, *MockCentralUseCase) {
	mockUseCase := new(MockCentralUseCase)
	centralHandler := handler.NewCentralHandler(mockUseCase)
	return centralHandler, mockUseCase
}

func TestCreateCentral_ValidData(t *testing.T) {
	app := fiber.New()
	centralHandler, mockUseCase := setupHandler()

	app.Post("/central", centralHandler.CreateCentral)

	// Dados válidos
	data := domain.Central{Name: "Central 1", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"}
	payload, _ := json.Marshal(data)

	mockUseCase.On("CreateCentral", mock.AnythingOfType("*domain.Central")).Return(nil)

	req := httptest.NewRequest(http.MethodPost, "/central", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	mockUseCase.AssertCalled(t, "CreateCentral", mock.AnythingOfType("*domain.Central"))
}

func TestCreateCentral_InvalidData(t *testing.T) {
	app := fiber.New()
	centralHandler, _ := setupHandler()

	app.Post("/central", centralHandler.CreateCentral)

	// Dados inválidos
	data := domain.Central{Name: "", MAC: "", IP: ""}
	payload, _ := json.Marshal(data)

	req := httptest.NewRequest(http.MethodPost, "/central", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestGetAllCentrals(t *testing.T) {
	app := fiber.New()
	centralHandler, mockUseCase := setupHandler()

	app.Get("/centrals", centralHandler.GetAllCentrals)

	// Dados simulados
	centrals := []domain.Central{
		{Name: "Central 1", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"},
		{Name: "Central 2", MAC: "00:11:22:33:44:56", IP: "192.168.0.2"},
	}
	mockUseCase.On("GetAllCentrals").Return(centrals, nil)

	req := httptest.NewRequest(http.MethodGet, "/centrals", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUseCase.AssertCalled(t, "GetAllCentrals")
}

func TestGetCentralByID_ValidID(t *testing.T) {
	app := fiber.New()
	centralHandler, mockUseCase := setupHandler()

	app.Get("/central/:id", centralHandler.GetCentralByID)

	// Simula central encontrada
	mockCentral := &domain.Central{
		ID:   1,
		Name: "Central Test",
		MAC:  "00:11:22:33:44:55",
		IP:   "192.168.0.1",
	}
	mockUseCase.On("GetCentralByID", uint(1)).Return(mockCentral, nil)

	req := httptest.NewRequest(http.MethodGet, "/central/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUseCase.AssertCalled(t, "GetCentralByID", uint(1))
}

func TestGetCentralByID_InvalidID(t *testing.T) {
	app := fiber.New()
	centralHandler, mockUseCase := setupHandler()

	app.Get("/central/:id", centralHandler.GetCentralByID)

	// Simula central não encontrada
	mockUseCase.On("GetCentralByID", uint(99)).Return((*domain.Central)(nil), errors.New("not found"))

	req := httptest.NewRequest(http.MethodGet, "/central/99", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	mockUseCase.AssertCalled(t, "GetCentralByID", uint(99))
}

func TestUpdateCentral(t *testing.T) {
	app := fiber.New()
	centralHandler, mockUseCase := setupHandler()

	app.Put("/central/:id", centralHandler.UpdateCentral)

	// Dados válidos
	data := domain.Central{Name: "Updated Central", MAC: "00:11:22:33:44:55", IP: "192.168.0.2"}
	payload, _ := json.Marshal(data)

	mockUseCase.On("UpdateCentral", mock.AnythingOfType("*domain.Central")).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/central/1", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	mockUseCase.AssertCalled(t, "UpdateCentral", mock.AnythingOfType("*domain.Central"))
}

func TestDeleteCentral(t *testing.T) {
	app := fiber.New()
	centralHandler, mockUseCase := setupHandler()

	app.Delete("/central/:id", centralHandler.DeleteCentral)

	mockUseCase.On("DeleteCentral", uint(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/central/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	mockUseCase.AssertCalled(t, "DeleteCentral", uint(1))
}
