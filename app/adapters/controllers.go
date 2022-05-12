package adapters

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		FetchByID() fiber.Handler
		Fetch() fiber.Handler
		Create() fiber.Handler
		Update() fiber.Handler
		Delete() fiber.Handler
	}
)
