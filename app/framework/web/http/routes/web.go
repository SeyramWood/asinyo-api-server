package routes

import (
	"time"

	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/handlers"
	"github.com/gofiber/fiber/v2"
)

type HttpRouter struct {
}

func NewHttpRouter(db *database.Adapter) *HttpRouter {
	return &HttpRouter{}
}

func (h *HttpRouter) Router(app *fiber.App) {
	// Custom config
	app.Static("/", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	r := app.Group("")
	// r.Get("/", handlers.Index())

	pageRouter(r)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

}

func pageRouter(r fiber.Router) {
	r.Get("/", handlers.Index())
}
