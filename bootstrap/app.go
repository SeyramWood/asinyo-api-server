package bootstrap

import (
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/env"
	"github.com/SeyramWood/pkg/router"
	"github.com/SeyramWood/pkg/server"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	env.Setup()
}

func App() *server.HTTP {

	db := database.NewDB("mysql")

	app := server.NewHTTPServer()

	app.Server.Server().MaxConnsPerIP = 1

	app.Server.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Server.Use(recover.New())
	app.Server.Use(logger.New())

	app.Server.Get("/dashboard", monitor.New())

	router.NewRouter(app.Server, db)

	return app

}
