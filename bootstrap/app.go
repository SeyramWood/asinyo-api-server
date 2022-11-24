package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/SeyramWood/app/application/logistic"
	"github.com/SeyramWood/app/application/mailer"
	"github.com/SeyramWood/app/application/maps"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/migrate"
	"github.com/SeyramWood/pkg/app"
	"github.com/SeyramWood/pkg/env"
	"github.com/SeyramWood/pkg/router"
)

func init() {
	env.Setup()
}

func App() {

	db := database.NewDB()
	defer func(DB *ent.Client) {
		_ = DB.Close()
	}(db.DB)

	ctx := context.Background()
	// Run migration.
	if err := db.DB.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	newApp := app.New()

	newApp.HTTP.Use(
		cors.New(
			cors.Config{
				AllowCredentials: true,
				AllowOrigins:     "*",
			},
		),
	)

	newApp.HTTP.Use(recover.New())

	newApp.HTTP.Use(logger.New())

	mail := mailer.NewEmail(newApp.Mailer)

	logis := logistic.NewLogistic(newApp.WG, db)

	ms := maps.NewMaps(newApp.WG)

	router.NewRouter(newApp.HTTP, db, mail, logis, ms)

	go mail.Listen()

	go logis.Listen()

	go ms.Listen()

	go newApp.Run()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, os.Interrupt, syscall.SIGTERM,
	) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")

	newApp.WG.Wait()

	_ = newApp.HTTP.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	_ = db.DB.Close()
	mail.Done()
	mail.CloseChannels()
	logis.Done()
	logis.CloseChannels()
	ms.Done()
	ms.CloseChannels()

	fmt.Println("Fiber was successful shutdown.")
}
