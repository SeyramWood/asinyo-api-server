package request

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/pkg/validator"
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
func ValidateChangePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.ChangePassword
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
func ValidateUserName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Get("userType")
		if userType == "merchant" {
			verifyUsername := struct {
				Username string `json:"username" validate:"required|email_phone|unique:merchants"`
			}{}
			err := c.BodyParser(&verifyUsername)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
			}
			if er := validator.Validate(&verifyUsername); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
		}
		if userType == "customer" {
			verifyUsername := struct {
				Username string `json:"username" validate:"required|email_phone|unique:customers"`
			}{}
			err := c.BodyParser(&verifyUsername)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
			}
			if er := validator.Validate(&verifyUsername); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
		}
		if userType == "agent" {
			verifyUsername := struct {
				Username string `json:"username" validate:"required|email_phone|unique:agents"`
			}{}
			err := c.BodyParser(&verifyUsername)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
			}
			if er := validator.Validate(&verifyUsername); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
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

func ValidateAgentCompliance() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var compliance models.AgentCompliance
		var guarantor models.AgentGuarantor

		if c.Get("step") == "personal" {
			err := c.BodyParser(&compliance)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}

			if er := validator.Validate(&compliance); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})

		} else {
			err := c.BodyParser(&guarantor)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}
			if er := validator.Validate(&guarantor); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Next()
		}
	}
}

func ValidateMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var retailerInfo models.RetailMerchantRequestInfo
		var supplierInfo models.SupplierMerchantRequestInfo

		var request models.MerchantRequest

		if c.Get("step") == "one" {
			var err error
			if c.Get("merchantType") == "supplier" {
				err = c.BodyParser(&supplierInfo)
			} else {
				err = c.BodyParser(&retailerInfo)
			}
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}

			if c.Get("merchantType") == "supplier" {
				if er := validator.Validate(&supplierInfo); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			} else {
				if er := validator.Validate(&retailerInfo); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
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

func ValidateNewMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var retailerInfo models.RetailerStorePersonalInfo
		var supplierInfo models.SupplierStorePersonalInfo

		if c.Get("step") == "one" {
			var err error
			if c.Get("merchantType") == "supplier" {
				err = c.BodyParser(&supplierInfo)
			} else {
				err = c.BodyParser(&retailerInfo)
			}
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if c.Get("merchantType") == "supplier" {
				if er := validator.Validate(&supplierInfo); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			} else {
				if er := validator.Validate(&retailerInfo); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})

		} else {
			var allRequest models.StoreFinalRequest
			err := c.BodyParser(&allRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(
				&models.MerchantStore{
					BusinessName:     allRequest.BusinessName,
					About:            allRequest.About,
					DescriptionTitle: allRequest.DescriptionTitle,
					Description:      allRequest.Description,
					Image:            allRequest.Image,
					OtherImages:      allRequest.OtherImages,
					Region:           allRequest.Region,
					District:         allRequest.District,
					City:             allRequest.City,
					Account:          allRequest.Account,
					MerchantType:     allRequest.MerchantType,
				},
			); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}

			return c.Next()
		}

	}
}

func ValidateAddress() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.Address
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		return c.Next()

	}
}

func ValidateMerchantStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantStore
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		return c.Next()

	}
}

func ValidateMerchantMomoAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantMomoAccountRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		return c.Next()

	}
}

func ValidateMerchantBankAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantBankAccountRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		return c.Next()

	}
}

func ValidateProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Product

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}

		return c.Next()

	}
}

func ValidateProductCatMajor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.ProductCategoryMajor

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}

		return c.Next()

	}
}

func ValidateProductCatMinor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.ProductCategoryMinor

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}

		return c.Next()

	}
}
