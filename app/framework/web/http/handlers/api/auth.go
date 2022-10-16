package api

import (
	"strconv"
	"time"

	"github.com/faabiosr/cachego/file"
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/auth"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type authHandler struct {
	service gateways.AuthService
}

func NewAuthHandler(db *database.Adapter, mail gateways.EmailService) *authHandler {
	repo := auth.NewAuthRepo(db)
	service := auth.NewAuthService(repo, mail)

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
func (auth *authHandler) ChangePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.ChangePassword
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		id := c.Params("user")
		userType := c.Params("userType")
		isOTP, _ := strconv.ParseBool(c.Get("OTP-Password-Change"))
		_, err = auth.service.UpdatePassword(id, &request, userType, isOTP)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"msg":    "Password successfully updated.",
			},
		)
	}
}
func (auth *authHandler) ResetPassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.ResetPassword
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		username := c.Params("user")
		userType := c.Params("userType")
		_, err = auth.service.ResetPassword(&request, username, userType)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"msg":    "Password successfully reset.",
			},
		)
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

		cache := file.New("./mnt/cache/otp/")
		if cache.Contains(request.Username) {
			userCode, err := cache.Fetch(request.Username)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			} else {
				return c.Status(fiber.StatusOK).JSON(
					fiber.Map{
						"status": true,
						"code":   userCode,
					},
				)
			}
		}
		code, err := auth.service.SendUserVerificationCode(request.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
					"msg":    "Could not send verification",
				},
			)
		}
		if err := cache.Save(request.Username, code, 24*time.Hour); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
					"msg":    "Could not saved OTP in cache",
				},
			)
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"code":   code,
			},
		)
	}
}
func (auth *authHandler) SendPasswordResetCode() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := struct {
			Username string
		}{}
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		cache := file.New("./mnt/cache/otp/")
		if cache.Contains(request.Username) {
			userCode, err := cache.Fetch(request.Username)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			} else {
				return c.Status(fiber.StatusOK).JSON(
					fiber.Map{
						"status": true,
						"code":   userCode,
					},
				)
			}
		}
		code, err := auth.service.SendPasswordResetCode(request.Username, c.Get("userType"))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"status": false,
					"msg":    "Could not send reset code",
				},
			)
		}
		if code != "" {
			if err := cache.Save(request.Username, code, 1*time.Hour); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
						"msg":    "Could not saved OTP in cache",
					},
				)
			}
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"code":   code,
			},
		)
	}
}
