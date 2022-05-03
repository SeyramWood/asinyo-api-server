package router

import (
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/routes"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, db *database.Adapter) {
	setup(app, routes.NewApiRouter(db), routes.NewHttpRouter(db))
}
func setup(app *fiber.App, routers ...Router) {
	for _, r := range routers {
		r.Router(app)
	}
}
