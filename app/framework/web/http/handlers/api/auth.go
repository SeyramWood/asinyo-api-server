package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/auth"
	"github.com/SeyramWood/app/framework/database"
)

type authHandler struct {
	service gateways.AuthService
}

func NewAuthHandler(db *database.Adapter) *authHandler {
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
func (auth *authHandler) FetchAuthUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return auth.service.FetchAuthUser(c)
	}
}
func (auth *authHandler) SendVerificationCode() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := struct {
			Username string
		}{}
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		code, err := auth.service.SendUserVerificationCode(request.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Could not send verification",
				},
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"code": code,
			},
		)
	}
}
