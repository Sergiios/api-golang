package handler

import (
	"api-golang/internal/domain"
	"api-golang/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CentralUseCase interface {
	CreateCentral(central *domain.Central) error
	GetAllCentrals() ([]domain.Central, error)
	GetCentralByID(id uint) (*domain.Central, error)
	UpdateCentral(central *domain.Central) error
	DeleteCentral(id uint) error
}

type CentralHandler struct {
	UseCase   CentralUseCase
	Validator *validator.Validate
}

func NewCentralHandler(uc CentralUseCase) *CentralHandler {
	return &CentralHandler{
		UseCase:   uc,
		Validator: validator.New(),
	}
}

// Create Central
func (h *CentralHandler) CreateCentral(c *fiber.Ctx) error {
	var central domain.Central

	// Parse JSON do corpo da requisição
	if err := c.BodyParser(&central); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	// Validação usando Validator
	if err := h.Validator.Struct(central); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.FormatValidationErrors(err),
		})
	}

	// Chama o caso de uso para criar a central
	if err := h.UseCase.CreateCentral(&central); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(central)
}

// Get All Centrals
func (h *CentralHandler) GetAllCentrals(c *fiber.Ctx) error {
	centrals, err := h.UseCase.GetAllCentrals()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(centrals)
}

// Get Central by ID
func (h *CentralHandler) GetCentralByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	central, err := h.UseCase.GetCentralByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Central not found"})
	}
	return c.JSON(central)
}

// Update Central
func (h *CentralHandler) UpdateCentral(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var central domain.Central

	// Parse JSON do corpo da requisição
	if err := c.BodyParser(&central); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}

	// Validação usando Validator
	if err := h.Validator.Struct(central); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": utils.FormatValidationErrors(err),
		})
	}

	// Define o ID da central antes de atualizar
	central.ID = uint(id)
	if err := h.UseCase.UpdateCentral(&central); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(central)
}

// Delete Central
func (h *CentralHandler) DeleteCentral(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.UseCase.DeleteCentral(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
