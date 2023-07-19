package app

import (
	"log"
	"os"
	"sync"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application/cache"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/pkg/jwt"
)

type Server struct {
	HTTP   *fiber.App
	Mailer *config.MailServer
	WG     *sync.WaitGroup
	JWT    *jwt.JWT
	Logger *zap.Logger
	Cache  gateways.CacheService
}

func New() *Server {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	jwtServ := jwt.NewJWT().Command()
	appcache := cache.New()

	var wg = sync.WaitGroup{}
	errorChan := make(chan error)
	mailerChan := make(chan *services.MailerMessage, 1024)
	failedMailerChan := make(chan *services.MailerMessage, 1024)
	mailerDoneChan := make(chan bool)

	return &Server{
		HTTP: fiber.New(
			fiber.Config{
				Prefork:                 config.Server().Prefork,
				CaseSensitive:           config.Server().CaseSensitive,
				StrictRouting:           config.Server().StrictRouting,
				ServerHeader:            config.Server().ServerHeader,
				EnablePrintRoutes:       config.Server().EnablePrintRoutes,
				EnableTrustedProxyCheck: config.Server().EnableTrustedProxyCheck,
				AppName:                 config.App().Name,
				BodyLimit:               10485760,
				JSONEncoder:             json.Marshal,
				JSONDecoder:             json.Unmarshal,
			},
		),
		Mailer: &config.MailServer{
			SMTP:           config.SMTPServer(),
			FailedDataChan: failedMailerChan,
			WG:             &wg,
			MailerChan:     mailerChan,
			DoneChan:       mailerDoneChan,
			ErrorChan:      errorChan,
		},
		WG:     &wg,
		JWT:    jwtServ,
		Logger: logger,
		Cache:  appcache,
	}

}

func (http *Server) Run() {
	if os.Getenv("APP_ENV") == "production" {
		port := os.Getenv("PORT")
		if port == "" {
			port = config.App().PORT
		}
		log.Fatal(http.HTTP.Listen(":" + port))
	} else {
		log.Fatal(http.HTTP.Listen(":" + config.App().PORT))
	}
}
