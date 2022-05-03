package request

import (
	"fmt"

	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

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
			fmt.Println(er)
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()
	}
}
func ValidateAgent() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var info models.AgentInfo
		var credentials models.AgentCredentials

		fmt.Println(c.Get("step"))

		if c.Get("step") == "one" {
			err := c.BodyParser(&info)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}

			if er := validator.Validate(&info); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

		} else {
			err := c.BodyParser(&credentials)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}

			if er := validator.Validate(&credentials); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
		}

		return c.Next()
	}
}

func ValidateRetailMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.RetailMerchant

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.RetailMerchantErrorResponse(err))
		}

		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()
	}
}
func ValidateSupplierMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.SupplierMerchant

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.SupplierMerchantErrorResponse(err))
		}

		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()
	}
}
