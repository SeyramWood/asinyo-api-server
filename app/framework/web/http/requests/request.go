package request

import (
	"fmt"
	"time"

	"github.com/faabiosr/cachego/file"
	"github.com/goccy/go-json"
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
func ValidateChangePassword(changeType string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var updateRequest models.ChangePassword
		var resetRequest models.ResetPassword
		if changeType == "update" {
			err := c.BodyParser(&updateRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
			}
			if er := validator.Validate(&updateRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
		}
		if changeType == "reset" {
			err := c.BodyParser(&resetRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
			}
			if er := validator.Validate(&resetRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
		}

		return c.Next()
	}
}
func ValidateUserName(checkUsernameExists bool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Get("userType")

		verifyUsername := struct {
			Username string `json:"username" validate:"required|email_phone"`
		}{}

		if userType == "merchant" {
			if checkUsernameExists {
				verifyUsernameExist := struct {
					Username string `json:"username" validate:"required|email_phone|unique:merchants"`
				}{}
				err := c.BodyParser(&verifyUsernameExist)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsernameExist); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			} else {
				err := c.BodyParser(&verifyUsername)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsername); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			}
		}
		if userType == "customer" {
			if checkUsernameExists {
				verifyUsernameExist := struct {
					Username string `json:"username" validate:"required|email_phone|unique:customers"`
				}{}
				err := c.BodyParser(&verifyUsernameExist)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsernameExist); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			} else {
				err := c.BodyParser(&verifyUsername)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsername); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			}
		}
		if userType == "agent" {
			if checkUsernameExists {
				verifyUsernameExist := struct {
					Username string `json:"username" validate:"required|email_phone|unique:agents"`
				}{}
				err := c.BodyParser(&verifyUsernameExist)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsernameExist); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			} else {
				err := c.BodyParser(&verifyUsername)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsername); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			}
		}

		if userType == "admin" {
			if checkUsernameExists {
				verifyUsernameExist := struct {
					Username string `json:"username" validate:"required|email_phone|unique:admins"`
				}{}
				err := c.BodyParser(&verifyUsernameExist)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsernameExist); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
			} else {
				err := c.BodyParser(&verifyUsername)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(presenters.AuthErrorResponse(err))
				}
				if er := validator.Validate(&verifyUsername); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
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
		var merchAddress models.MerchantStoreAddress

		cache := file.New("./mnt/cache/")
		agentId := c.Params("agent")

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
				dataKey := fmt.Sprintf("step_one_%s", agentId)
				if cache.Contains(dataKey) {
					cache.Delete(dataKey)
				}
				data, _ := json.Marshal(supplierInfo)
				if err := cache.Save(dataKey, string(data), 25*time.Minute); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(
						fiber.Map{
							"status": false,
							"msg":    "Could not cached data",
						},
					)
				}

			} else {
				if er := validator.Validate(&retailerInfo); er != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
				}
				dataKey := fmt.Sprintf("step_one_%s", agentId)
				if cache.Contains(dataKey) {
					cache.Delete(dataKey)
				}
				data, _ := json.Marshal(supplierInfo)
				if err := cache.Save(dataKey, string(data), 25*time.Minute); err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(
						fiber.Map{
							"status": false,
							"msg":    "Could not cached data",
						},
					)
				}
			}

			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})

		} else if c.Get("step") == "two" {
			var err error
			err = c.BodyParser(&merchAddress)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&merchAddress); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
			dataKey := fmt.Sprintf("step_two_%s", agentId)
			if cache.Contains(dataKey) {
				cache.Delete(dataKey)
			}
			data, _ := json.Marshal(merchAddress)
			if err := cache.Save(dataKey, string(data), 25*time.Minute); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
						"msg":    "Could not cached data",
					},
				)
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})

		} else {
			var merchantStoreInfo models.MerchantStoreInfo
			err := c.BodyParser(&merchantStoreInfo)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&merchantStoreInfo); er != nil {
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
		if er := validator.Validate(&request); er != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
		}
		return c.Next()

	}
}

func ValidateMerchantStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("step") == "one" {
			var addressRequest models.MerchantStoreAddress
			err := c.BodyParser(&addressRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&addressRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
			cache := file.New("./mnt/cache/")
			merchantId := c.Params("merchantId")
			dataKey := fmt.Sprintf("step_one_%s", merchantId)
			if cache.Contains(dataKey) {
				cache.Delete(dataKey)
			}
			data, _ := json.Marshal(addressRequest)
			if err := cache.Save(dataKey, string(data), 25*time.Minute); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(
					fiber.Map{
						"status": false,
						"msg":    "Could not cached data",
					},
				)
			}
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})
		} else {
			var storeRequest models.MerchantStore
			err := c.BodyParser(&storeRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&storeRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
		}
		return c.Next()

	}
}
func ValidateMerchantStoreUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var infoRequest models.MerchantStore
		var addressRequest models.MerchantStoreAddress

		formType := c.Get("formType")
		if formType == "information" {
			err := c.BodyParser(&infoRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&infoRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
			
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"ok": true})
		}
		if formType == "address" {
			err := c.BodyParser(&addressRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			if er := validator.Validate(&addressRequest); er != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(er)
			}
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
