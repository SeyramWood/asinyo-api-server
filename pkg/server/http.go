package server

import (
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/pkg/env"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
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
	if env.Get("APP_ENV", "local") == "production" {
		// Get the PORT from heroku env
		//port := os.Getenv("PORT")
		// Verify if heroku provided the port or not
		//if os.Getenv("PORT") == "" {
		//	port = "9000"
		//}
		log.Fatal(http.Server.Listen(":" + config.App().PORT))
	} else {
		log.Fatal(http.Server.Listen(fmt.Sprint(config.App().ServerURL)))
	}
}
