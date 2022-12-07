package utils

import (
	"github.com/gofiber/fiber/v2"

	lApi "go-blogs-api/apis"
	lHttp "go-blogs-api/http"
)

type Router interface {
	InstallRouter(app *fiber.App)
}

func InstallRouter(app *fiber.App) {
	apiRouter := &ApiRouter{}
	httpRouter := &HttpRouter{}
	setup(app, apiRouter, httpRouter)
}

func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}

type ApiRouter struct {
}

func (h ApiRouter) InstallRouter(app *fiber.App) {
	// limitConfig := limiter.Config{Max: 500}
	// api := app.Group("/api", limiter.New(limitConfig))

	api := app.Group("/api")
	lApi.SetRootRoutes(api)
	lApi.SetArticleRoutes(api)
	lApi.SetCommentRoutes(api)
}

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	// limitConfig := limiter.Config{Max: 500}
	// api := app.Group("/api", limiter.New(limitConfig))

	httpApi := app.Group("/")
	lHttp.SetRootRoutes(httpApi)
}
