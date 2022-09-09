package handler

import (
	"github.com/gofiber/fiber/v2"
	"mangtas-test/core/domain/entity"
	"mangtas-test/core/ports"
)

type Handler struct {
	Div ports.IDivison
}

func (h *Handler) Division(ctx *fiber.Ctx) error {
	//09016016563
	payload := &entity.Payload{}

	err := ctx.BodyParser(payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to decode")
	}
	ans, err := h.Div.Divide(*payload)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to decode")
	}
	return ctx.JSON(map[string]interface{}{"answer": ans})
}
