package request

import (
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func ValidateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {

		userType := c.Get("Asinyo-Authorization-Type")
		var request models.User
		var merchantRequest models.UserMerchant

		if userType == "supplier" || userType == "retailer" {
			err := c.BodyParser(&merchantRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
			}
			if er := validator.Validate(&merchantRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Next()
		}

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
		}

		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()
	}
}

func ValidateAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Admin

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AdminErrorResponse(err))
		}

		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()
	}
}

func ValidateCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.Customer
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.CustomerErrorResponse(err))
		}

		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()
	}
}
func ValidateAgent() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var info models.AgentInfo
		var request models.AgentRequest

		if c.Get("step") == "one" {
			err := c.BodyParser(&info)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}

			if er := validator.Validate(&info); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})

		} else {
			err := c.BodyParser(&request)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}
			if er := validator.Validate(&request.Credentials); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Next()
		}
	}
}

func ValidateMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var info models.MerchantRequestInfo
		var request models.MerchantRequest

		if c.Get("step") == "one" {
			err := c.BodyParser(&info)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}

			if er := validator.Validate(&info); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})

		} else {
			err := c.BodyParser(&request)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&request.Credentials); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Next()
		}

	}
}

func ValidateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var retquest models.Product

		err := c.BodyParser(&retquest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}

		if er := validator.Validate(&retquest); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}

		return c.Next()

	}
}

func ValidateProductCatMajor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var retquest models.ProductCategoryMajor

		err := c.BodyParser(&retquest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		if er := validator.Validate(&retquest); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}

		return c.Next()

	}
}

func ValidateProductCatMinor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var retquest models.ProductCategoryMinor

		err := c.BodyParser(&retquest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		if er := validator.Validate(&retquest); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}

		return c.Next()

	}
}
