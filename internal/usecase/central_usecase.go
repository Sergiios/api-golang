package usecase

import (
	"api-golang/internal/domain"
	"api-golang/internal/repository"
)

type CentralUseCase struct {
	Repo *repository.CentralRepository
}

func NewCentralUseCase(repo *repository.CentralRepository) *CentralUseCase {
	return &CentralUseCase{Repo: repo}
}

func (uc *CentralUseCase) CreateCentral(user *domain.Central) error {
	return uc.Repo.Create(user)
}

func (uc *CentralUseCase) GetAllCentrals() ([]domain.Central, error) {
	return uc.Repo.GetAll()
}

func (uc *CentralUseCase) GetCentralByID(id uint) (*domain.Central, error) {
	return uc.Repo.GetByID(id)
}

func (uc *CentralUseCase) UpdateCentral(user *domain.Central) error {
	return uc.Repo.Update(user)
}

func (uc *CentralUseCase) DeleteCentral(id uint) error {
	return uc.Repo.Delete(id)
}
