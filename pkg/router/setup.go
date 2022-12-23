package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/routes"
)

func NewRouter(
	app *fiber.App, db *database.Adapter, mail gateways.EmailService, logis gateways.LogisticService,
	maps gateways.MapService,
	storageSrv gateways.StorageService,
) {
	setup(app, routes.NewApiRouter(db, mail, logis, maps, storageSrv), routes.NewPageRouter(db, mail, logis))
}

func setup(app *fiber.App, routers ...Router) {
	for _, r := range routers {
		r.Router(app)
	}
}
