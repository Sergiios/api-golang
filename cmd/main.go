package main

import (
	"api-golang/internal/config"
	"api-golang/internal/domain"
	"api-golang/internal/handler"
	"api-golang/internal/repository"
	"api-golang/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	db.AutoMigrate(&domain.Central{})

	app := fiber.New()

	repo := repository.NewCentralRepository(db)
	uc := usecase.NewCentralUseCase(repo)
	handler := handler.NewCentralHandler(uc)

	app.Post("/central", handler.CreateCentral)
	app.Get("/centrals", handler.GetAllCentrals)
	app.Get("/central/:id", handler.GetCentralByID)
	app.Put("/central/:id", handler.UpdateCentral)
	app.Delete("/central/:id", handler.DeleteCentral)

	log.Fatal(app.Listen(":8080"))
}
