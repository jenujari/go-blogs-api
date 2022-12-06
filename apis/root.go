package api

import "github.com/gofiber/fiber/v2"

func index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Jenish Jariwala",
	})
}

func SetRootRoutes(api fiber.Router) {
	api.Get("/", index)
}
