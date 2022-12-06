package api

import (
	"github.com/gofiber/fiber/v2"
)

func getComment(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"comments": "hello world",
	})
}

func SetCommentRoutes(api fiber.Router) {
	commentApi := api.Group("/comment")
	commentApi.Get("/", getComment)
}
