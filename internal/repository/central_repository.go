package repository

import (
	"api-golang/internal/domain"

	"gorm.io/gorm"
)

type CentralRepository struct {
	DB *gorm.DB
}

func NewCentralRepository(db *gorm.DB) *CentralRepository {
	return &CentralRepository{DB: db}
}

func (r *CentralRepository) Create(user *domain.Central) error {
	return r.DB.Create(user).Error
}

func (r *CentralRepository) GetAll() ([]domain.Central, error) {
	var users []domain.Central
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *CentralRepository) GetByID(id uint) (*domain.Central, error) {
	var user domain.Central
	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *CentralRepository) Update(user *domain.Central) error {
	return r.DB.Save(user).Error
}

func (r *CentralRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Central{}, id).Error
}
