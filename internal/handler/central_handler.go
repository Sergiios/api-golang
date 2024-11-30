package handler

import (
	"api-golang/internal/domain"
	"api-golang/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type CentralHandler struct {
	UseCase *usecase.CentralUseCase
}

func NewCentralHandler(uc *usecase.CentralUseCase) *CentralHandler {
	return &CentralHandler{UseCase: uc}
}

// Create User
func (h *CentralHandler) CreateCentral(c *fiber.Ctx) error {
	var cental domain.Central
	if err := c.BodyParser(&cental); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.UseCase.CreateCentral(&cental); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(cental)
}

// Get All Users
func (h *CentralHandler) GetAllCentrals(c *fiber.Ctx) error {
	centrals, err := h.UseCase.GetAllCentrals()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(centrals)
}

// Get User by ID
func (h *CentralHandler) GetCentralByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	central, err := h.UseCase.GetCentralByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Central not found"})
	}
	return c.JSON(central)
}

// Update User
func (h *CentralHandler) UpdateCentral(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var central domain.Central
	if err := c.BodyParser(&central); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	central.ID = uint(id)
	if err := h.UseCase.UpdateCentral(&central); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(central)
}

// Delete User
func (h *CentralHandler) DeleteCentral(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.UseCase.DeleteCentral(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
