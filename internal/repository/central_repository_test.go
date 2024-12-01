package repository_test

import (
	"api-golang/internal/domain"
	"api-golang/internal/repository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupInMemoryDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Cria a tabela Central
	err = db.AutoMigrate(&domain.Central{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}

func TestCreateCentral(t *testing.T) {
	db := setupInMemoryDB()
	repo := repository.NewCentralRepository(db)

	central := &domain.Central{
		Name: "Central Test",
		MAC:  "00:11:22:33:44:55",
		IP:   "192.168.0.1",
	}

	err := repo.Create(central)
	assert.NoError(t, err)

	// Verifica se foi salvo corretamente
	var result domain.Central
	err = db.First(&result, central.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, central.Name, result.Name)
	assert.Equal(t, central.MAC, result.MAC)
	assert.Equal(t, central.IP, result.IP)
}

func TestGetAllCentrals(t *testing.T) {
	db := setupInMemoryDB()
	repo := repository.NewCentralRepository(db)

	// Adiciona dados de teste
	db.Create(&domain.Central{Name: "Central 1", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"})
	db.Create(&domain.Central{Name: "Central 2", MAC: "00:11:22:33:44:56", IP: "192.168.0.2"})

	centrals, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, centrals, 2)
	assert.Equal(t, "Central 1", centrals[0].Name)
	assert.Equal(t, "Central 2", centrals[1].Name)
}

func TestGetCentralByID(t *testing.T) {
	db := setupInMemoryDB()
	repo := repository.NewCentralRepository(db)

	// Adiciona dado de teste
	db.Create(&domain.Central{Name: "Central Test", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"})

	central, err := repo.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, "Central Test", central.Name)
	assert.Equal(t, "00:11:22:33:44:55", central.MAC)

	// Testa ID inexistente
	central, err = repo.GetByID(99)
	assert.Nil(t, central)
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}

func TestUpdateCentral(t *testing.T) {
	db := setupInMemoryDB()
	repo := repository.NewCentralRepository(db)

	// Adiciona dado de teste
	db.Create(&domain.Central{Name: "Central Old", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"})

	central := &domain.Central{ID: 1, Name: "Central Updated", MAC: "00:11:22:33:44:55", IP: "192.168.0.2"}
	err := repo.Update(central)
	assert.NoError(t, err)

	// Verifica atualização
	var result domain.Central
	err = db.First(&result, 1).Error
	assert.NoError(t, err)
	assert.Equal(t, "Central Updated", result.Name)
	assert.Equal(t, "192.168.0.2", result.IP)
}

func TestDeleteCentral(t *testing.T) {
	db := setupInMemoryDB()
	repo := repository.NewCentralRepository(db)

	// Adiciona dado de teste
	db.Create(&domain.Central{Name: "Central To Delete", MAC: "00:11:22:33:44:55", IP: "192.168.0.1"})

	err := repo.Delete(1)
	assert.NoError(t, err)

	// Verifica se foi deletado
	var result domain.Central
	err = db.First(&result, 1).Error
	assert.Error(t, err) // Registro deve estar ausente
	assert.Equal(t, gorm.ErrRecordNotFound, err)

	// Testa exclusão de ID inexistente
	err = repo.Delete(99)
	fmt.Println(err)
	assert.NoError(t, err) // GORM não retorna erro para exclusões de IDs inexistentes
}
