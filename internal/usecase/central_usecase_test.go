package usecase_test

import (
	"api-golang/internal/domain"
	"api-golang/internal/usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do Repositório
type MockCentralRepository struct {
	mock.Mock
}

func (m *MockCentralRepository) Create(user *domain.Central) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockCentralRepository) GetAll() ([]domain.Central, error) {
	args := m.Called()
	return args.Get(0).([]domain.Central), args.Error(1)
}

func (m *MockCentralRepository) GetByID(id uint) (*domain.Central, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Central), args.Error(1)
}

func (m *MockCentralRepository) Update(user *domain.Central) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockCentralRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupUseCase() (*usecase.CentralUseCase, *MockCentralRepository) {
	mockRepo := new(MockCentralRepository)
	uc := usecase.NewCentralUseCase(mockRepo)
	return uc, mockRepo
}

func TestCreateCentral(t *testing.T) {
	uc, mockRepo := setupUseCase()

	// Dados de entrada
	central := &domain.Central{Name: "Central Test", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"}

	// Configura o mock
	mockRepo.On("Create", central).Return(nil)

	// Chama o método
	err := uc.CreateCentral(central)

	// Valida os resultados
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", central)
}

func TestGetAllCentrals(t *testing.T) {
	uc, mockRepo := setupUseCase()

	// Dados simulados
	centrals := []domain.Central{
		{Name: "Central 1", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"},
		{Name: "Central 2", MAC: "00:11:22:33:44:56", IP: "192.168.0.2"},
	}

	// Configura o mock
	mockRepo.On("GetAll").Return(centrals, nil)

	// Chama o método
	result, err := uc.GetAllCentrals()

	// Valida os resultados
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, centrals, result)
	mockRepo.AssertCalled(t, "GetAll")
}

func TestGetCentralByID_ValidID(t *testing.T) {
	uc, mockRepo := setupUseCase()

	// Dado simulado
	mockCentral := &domain.Central{Name: "Central 1", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"}

	// Configura o mock
	mockRepo.On("GetByID", uint(1)).Return(mockCentral, nil)

	// Chama o método
	result, err := uc.GetCentralByID(1)

	// Valida os resultados
	assert.NoError(t, err)
	assert.Equal(t, mockCentral, result)
	mockRepo.AssertCalled(t, "GetByID", uint(1))
}

func TestGetCentralByID_InvalidID(t *testing.T) {
	uc, mockRepo := setupUseCase()

	// Configura o mock para ID inválido
	mockRepo.On("GetByID", uint(99)).Return((*domain.Central)(nil), errors.New("not found"))

	// Chama o método
	result, err := uc.GetCentralByID(99)

	// Valida os resultados
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertCalled(t, "GetByID", uint(99))
}

func TestUpdateCentral(t *testing.T) {
	uc, mockRepo := setupUseCase()

	// Dados de entrada
	central := &domain.Central{ID: 1, Name: "Updated Central", MAC: "00:11:22:33:44:55", IP: "192.168.0.2"}

	// Configura o mock
	mockRepo.On("Update", central).Return(nil)

	// Chama o método
	err := uc.UpdateCentral(central)

	// Valida os resultados
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Update", central)
}

func TestDeleteCentral(t *testing.T) {
	uc, mockRepo := setupUseCase()

	// Configura o mock
	mockRepo.On("Delete", uint(1)).Return(nil)

	// Chama o método
	err := uc.DeleteCentral(1)

	// Valida os resultados
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", uint(1))
}
