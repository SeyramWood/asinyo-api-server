package server

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/SeyramWood/config"
)

type HTTP struct {
	Server *fiber.App
	Logger *zap.Logger
}

func NewHTTPServer() *HTTP {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &HTTP{
		Server: fiber.New(
			fiber.Config{
				Prefork:       config.Server().Prefork,
				CaseSensitive: config.Server().CaseSensitive,
				StrictRouting: config.Server().StrictRouting,
				ServerHeader:  config.Server().ServerHeader,
				AppName:       config.App().Name,
				JSONEncoder:   json.Marshal,
				JSONDecoder:   json.Unmarshal,
			},
		),
		Logger: logger,
	}

}

func (http *HTTP) Run() {
	if os.Getenv("APP_ENV") == "production" {
		port := os.Getenv("PORT")
		if port == "" {
			port = config.App().PORT
		}
		log.Fatal(http.Server.Listen(":" + port))
	} else {
		log.Fatal(http.Server.Listen(":" + config.App().PORT))
	}
}
