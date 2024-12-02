package usecase

import (
	"api-golang/internal/domain"
)

type CentralRepository interface {
	Create(user *domain.Central) error
	GetAll() ([]domain.Central, error)
	GetByID(id uint) (*domain.Central, error)
	Update(user *domain.Central) error
	Delete(id uint) error
}

type CentralUseCase struct {
	Repo CentralRepository
}

func NewCentralUseCase(repo CentralRepository) *CentralUseCase {
	return &CentralUseCase{Repo: repo}
}

func (uc *CentralUseCase) CreateCentral(central *domain.Central) error {
	return uc.Repo.Create(central)
}

func (uc *CentralUseCase) GetAllCentrals() ([]domain.Central, error) {
	return uc.Repo.GetAll()
}

func (uc *CentralUseCase) GetCentralByID(id uint) (*domain.Central, error) {
	return uc.Repo.GetByID(id)
}

func (uc *CentralUseCase) UpdateCentral(central *domain.Central) error {
	return uc.Repo.Update(central)
}

func (uc *CentralUseCase) DeleteCentral(id uint) error {
	return uc.Repo.Delete(id)
}
