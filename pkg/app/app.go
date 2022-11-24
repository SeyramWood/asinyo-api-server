package app

import (
	"log"
	"os"
	"sync"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
)

type serverconfig struct {
	HTTP   *fiber.App
	Mailer *config.MailServer
	WG     *sync.WaitGroup
	Logger *zap.Logger
}

func New() *serverconfig {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	var wg = sync.WaitGroup{}

	errorChan := make(chan error)
	mailerChan := make(chan *services.Message, 1024)
	mailerDoneChan := make(chan bool)

	return &serverconfig{
		HTTP: fiber.New(
			fiber.Config{
				Prefork:       config.Server().Prefork,
				CaseSensitive: config.Server().CaseSensitive,
				StrictRouting: config.Server().StrictRouting,
				ServerHeader:  config.Server().ServerHeader,
				AppName:       config.App().Name,
				BodyLimit:     10485760,
				JSONEncoder:   json.Marshal,
				JSONDecoder:   json.Unmarshal,
			},
		),
		Mailer: &config.MailServer{
			SMTP:       config.SMTPServer(),
			WG:         &wg,
			MailerChan: mailerChan,
			DoneChan:   mailerDoneChan,
			ErrorChan:  errorChan,
		},
		WG:     &wg,
		Logger: logger,
	}

}

func (http *serverconfig) Run() {
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
