package configs

import (
	"os"
	"path"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django"
)

func InitFiber() *fiber.App {
	viewPath := path.Join(os.Getenv("ABS_PATH"), "views")
	engine := django.New(viewPath, ".html")

	var config fiber.Config = fiber.Config{
		AppName:           os.Getenv("APP_NAME"),
		EnablePrintRoutes: false,
		Views:             engine,
	}

	app := fiber.New(config)

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/health", monitor.New())

	return app
}
