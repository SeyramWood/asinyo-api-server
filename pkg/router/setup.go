package router

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application/app_cache"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/app/framework/web/http/routes"
	"github.com/SeyramWood/pkg/app"
)

func NewRouter(
	app *app.Server,
	db *database.Adapter,
	noti notification.NotificationService,
	dbNoti gateways.DBNotificationService,
	storageSrv gateways.StorageService,
	logis gateways.LogisticService,
	ms gateways.MapService,
	appcache *app_cache.AppCache,
) {
	setup(
		app, routes.NewApiRouter(app, db, noti, dbNoti, storageSrv, logis, ms, appcache),
		routes.NewPageRouter(app, db, noti, dbNoti, storageSrv, logis, ms, appcache),
	)
}

func setup(app *app.Server, routers ...Router) {
	for _, r := range routers {
		r.Router(app.HTTP)
	}
}
