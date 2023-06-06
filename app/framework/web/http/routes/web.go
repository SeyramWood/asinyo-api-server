package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application/app_cache"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/framework/database"
	handler "github.com/SeyramWood/app/framework/web/http/handlers/page"
	"github.com/SeyramWood/pkg/app"
)

type PageRouter struct {
}

func NewPageRouter(
	app *app.Server,
	db *database.Adapter,
	noti notification.NotificationService,
	dbNoti gateways.DBNotificationService,
	storageSrv gateways.StorageService,
	logis gateways.LogisticService,
	ms gateways.MapService,
	appcache *app_cache.AppCache,
) *PageRouter {
	return &PageRouter{}
}

func (h *PageRouter) Router(app *fiber.App) {

	r := app.Group("")

	r.Get("/", handler.Index())

	pageRouter(r)

	// Custom config
	r.Static(
		"/", "./public/storage", fiber.Static{
			Compress:      true,
			Browse:        false,
			CacheDuration: 10 * time.Second,
			MaxAge:        3600,
		},
	)
	// 404 Handler
	app.Use(
		func(c *fiber.Ctx) error {
			return c.SendStatus(404) // => 404 "Not Found"
		},
	)

}

func pageRouter(r fiber.Router) {
	r.Get("/", handler.Index())

	r.Get("/dashboard", monitor.New())

}
