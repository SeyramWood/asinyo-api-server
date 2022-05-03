package router

import "github.com/gofiber/fiber/v2"

type Router interface {
	Router(app *fiber.App)
}
