package handlers

import (
	"github.com/SeyramWood/app/adapters"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application/auth"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service gateways.AuthService
}

func NewAuthHandler(db *database.Adapter) adapters.AuthHandler {
	repo := auth.NewAuthRepo(db)
	service := auth.NewAuthService(repo)

	return &authHandler{
		service: service,
	}

}

func (auth *authHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return auth.service.Login(c)
	}
}
func (auth *authHandler) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return auth.service.Logout(c)
	}
}
