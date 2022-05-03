package routes

import (
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

type HttpRouter struct {
}

func NewHttpRouter(db *database.Adapter) *HttpRouter {
	return &HttpRouter{}
}

func (h *HttpRouter) Router(app *fiber.App) {
	r := app.Group("", csrf.New())
	r.Get("/", handlers.Index())
	pageRouter(r)
	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

}

func pageRouter(r fiber.Router) {
	r.Get("/", handlers.Index())
}
