package http

import "github.com/gofiber/fiber/v2"

func index(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"Title": "Custom title",
	})
}

func SetRootRoutes(api fiber.Router) {
	api.Get("/", index)
}
