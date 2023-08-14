package bootstrap

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/SeyramWood/app/application/db_notification"
	"github.com/SeyramWood/app/application/logistic"
	"github.com/SeyramWood/app/application/mailer"
	"github.com/SeyramWood/app/application/maps"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/application/sms"
	"github.com/SeyramWood/app/application/storage"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/config"
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

	ctx := context.Background()
	// Run migration.
	if err := db.DB.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	newApp := app.New()

	newApp.HTTP.Use(
		limiter.New(
			limiter.Config{
				Max:               20,
				Expiration:        30 * time.Second,
				LimiterMiddleware: limiter.SlidingWindow{},
			},
		),
	)
	newApp.HTTP.Use(
		helmet.New(
			helmet.Config{
				XSSProtection:             "1; mode=block",
				ContentTypeNosniff:        "nosniff",
				XFrameOptions:             "SAMEORIGIN",
				ReferrerPolicy:            "no-referrer",
				CrossOriginEmbedderPolicy: "require-corp",
				CrossOriginOpenerPolicy:   "same-origin",
				CrossOriginResourcePolicy: "same-origin",
				OriginAgentCluster:        "?1",
				XDNSPrefetchControl:       "off",
				XDownloadOptions:          "noopen",
				XPermittedCrossDomain:     "none",
			},
		),
	)
	newApp.HTTP.Use(
		cors.New(
			cors.Config{
				AllowOrigins: config.App().AllowOrigins,
			},
		),
	)
	newApp.HTTP.Use(idempotency.New())
	newApp.HTTP.Use(compress.New())
	newApp.HTTP.Use(recover.New())

	newApp.HTTP.Use(logger.New())

	mail := mailer.NewEmail(newApp.Mailer)
	smsServ := sms.NewSMSService(newApp.WG)
	dbNoti := db_notification.NewDBNotificationService(newApp.WG, db)

	noti := notification.NewNotification(dbNoti, smsServ, mail)
	noti.Subscribe(noti)

	logis := logistic.NewLogistic(newApp.WG, db)

	ms := maps.NewMaps(newApp.WG)

	storageSrv := storage.NewStorageService(newApp.WG)

	router.NewRouter(
		newApp, db, noti, dbNoti, storageSrv, logis, ms,
	)

	// go appcache.CleanUp()

	go mail.Listen()

	go smsServ.Listen()

	go dbNoti.Listen()

	go logis.Listen()

	go ms.Listen()

	go storageSrv.Listen()

	go newApp.Run()

	c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
	signal.Notify(
		c, syscall.SIGINT, syscall.SIGTERM,
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
	smsServ.Done()
	smsServ.CloseChannels()
	dbNoti.Done()
	dbNoti.CloseChannels()

	logis.Done()
	logis.CloseChannels()

	ms.Done()
	ms.CloseChannels()

	storageSrv.Done()
	storageSrv.CloseChannels()

	fmt.Println("API Server was successful shutdown.")
}
