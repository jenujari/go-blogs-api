package api

import (
	"github.com/gofiber/fiber/v2"
)

func getArticle(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"article": "hello world",
	})
}

func SetArticleRoutes(api fiber.Router) {
	articleApi := api.Group("/article")
	articleApi.Get("/", getArticle)
}
