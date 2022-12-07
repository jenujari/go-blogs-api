package apis

import "github.com/gofiber/fiber/v2"

func SetRootRoutes(api fiber.Router) {
	api.Get("/", index)
}

func index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello world",
	})
}
