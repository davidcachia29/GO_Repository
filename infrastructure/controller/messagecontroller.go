package controller

import (
	"TestProject/infrastructure/usecase"

	"github.com/gofiber/fiber/v2"
)

type MessageController struct {
	UseCase *usecase.MessageUseCase
}

// @Summary Get Call
// @Description Displayes the call that is made
// @Tags messages
// @Accept json
// @Produce  json
// @Success 200 {array} entity.Message
// @Router /path [get]
func (c *MessageController) HandleGetMessage(ctx *fiber.Ctx) error {
	messageText := c.UseCase.GetMessageText()
	return ctx.JSON(messageText)
}
