package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/routes"
)

func NewRouter(app *fiber.App, db *database.Adapter, mail gateways.EmailService) {
	setup(app, routes.NewApiRouter(db, mail), routes.NewPageRouter(db, mail))
}
func setup(app *fiber.App, routers ...Router) {
	for _, r := range routers {
		r.Router(app)
	}
}
