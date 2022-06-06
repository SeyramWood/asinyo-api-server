package server

import (
	"fmt"
	"log"

	"github.com/SeyramWood/config"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
		Server: fiber.New(fiber.Config{
			Prefork:       config.Server().Prefork,
			CaseSensitive: config.Server().CaseSensitive,
			StrictRouting: config.Server().StrictRouting,
			ServerHeader:  config.Server().ServerHeader,
			AppName:       config.App().Name,
			JSONEncoder:   json.Marshal,
			JSONDecoder:   json.Unmarshal,
		}),
		Logger: logger,
	}

}

func (http *HTTP) Run() {
	log.Fatal(http.Server.Listen(fmt.Sprint(config.App().ServerURL)))
}
