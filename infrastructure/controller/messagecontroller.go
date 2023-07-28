package controller

import (
	"TestProject/infrastructure/usecase"

	"github.com/gofiber/fiber/v2"
)

type MessageController struct {
	UseCase *usecase.MessageUseCase
}

// @Summary Get messages
// @Description Fetches all messages. You can add more details here.
// @Tags messages
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Message
// @Router /messages [get]
func (c *MessageController) HandleGetMessage(ctx *fiber.Ctx) error {
	messageText := c.UseCase.GetMessageText()
	return ctx.JSON(messageText)
}
