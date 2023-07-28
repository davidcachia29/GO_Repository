package main

import (
	"TestProject/infrastructure/controller"
	_ "TestProject/infrastructure/docs"
	"TestProject/infrastructure/entity"
	"TestProject/infrastructure/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using swagger with fiber.
// @host localhost:8080
// @BasePath /
func main() {
	app := fiber.New()

	message := entity.Message{Text: "/"}
	useCase := &usecase.MessageUseCase{Message: message}
	ctrl := &controller.MessageController{UseCase: useCase}

	//Routes
	app.Get("/test", ctrl.HandleGetMessage)

	app.Get("/swagger/*", swagger.HandlerDefault)

	println("Starting server on port 8080")
	app.Listen(":8080")
}
