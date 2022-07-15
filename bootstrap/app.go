package bootstrap

import (
	"context"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent/migrate"
	"github.com/SeyramWood/pkg/env"
	"github.com/SeyramWood/pkg/router"
	"github.com/SeyramWood/pkg/server"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
)

func init() {
	env.Setup()
}

func App() {

	db := database.NewDB("mysql")
	defer db.DB.Close()

	ctx := context.Background()
	// Run migration.
	if err := db.DB.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app := server.NewHTTPServer()

	app.Server.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	app.Server.Use(recover.New())

	app.Server.Use(logger.New())

	app.Server.Get("/dashboard", monitor.New())

	router.NewRouter(app.Server, db)

	app.Run()
	//return app
}
