package main

import (
	"TestProject/infrastructure/controller"
	_ "TestProject/infrastructure/docs"
	"TestProject/infrastructure/entity"
	"TestProject/infrastructure/usecase"

	custom_logger "TestProject/infrastructure/custom-logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using swagger with fiber.
// @host 127.0.0.1:8080
// @BasePath /
func main() {

	customLogger := custom_logger.GetInstance()

	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			customLogger.Write([]byte(err.Error() + "\n"))
			return c.SendStatus(fiber.StatusInternalServerError)
		},
	})

	app.Use(logger.New(logger.Config{
		Format:     "${time} ${method} ${path} - ${ip} - ${status}\n",
		TimeFormat: "2006-01-02 15:04:05",
		Output:     customLogger,
	}))

	message := entity.Message{Text: "/"}
	useCase := &usecase.MessageUseCase{Message: message}
	ctrl := &controller.MessageController{UseCase: useCase}

	// Initialize default config
	app.Use(logger.New())

	//Routes
	app.Get("/path", ctrl.HandleGetMessage)

	app.Get("/swagger/*", swagger.HandlerDefault)

	//implementation of metrics endpoint below
	app.Get("/metrics", monitor.New())

	println("Starting server on port 8080")
	app.Listen(":8080")
}
