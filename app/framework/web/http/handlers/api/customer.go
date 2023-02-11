package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/customer"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type CustomerHandler struct {
	service    gateways.CustomerService
	storageSrv gateways.StorageService
}

func NewCustomerHandler(db *database.Adapter, storageSrv gateways.StorageService) *CustomerHandler {
	repo := customer.NewCustomerRepo(db)
	service := customer.NewCustomerService(repo)

	return &CustomerHandler{
		service:    service,
		storageSrv: storageSrv,
	}
}

func (h *CustomerHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.CustomerSuccessResponse(result))
	}
}
func (h *CustomerHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.CustomersSuccessResponse(result))
	}
}

func (h *CustomerHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var individualRequest models.IndividualCustomer
		var businessRequest models.BusinessCustomer
		if c.Get("customerType") == "individual" {
			err := c.BodyParser(&individualRequest)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(presenters.CustomerErrorResponse(err))
			}
			result, err := h.service.Create(&individualRequest, c.Get("customerType"))
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(presenters.CustomerErrorResponse(err))
			}
			return c.JSON(presenters.CustomerSuccessResponse(result))
		}

		err := c.BodyParser(&businessRequest)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.CustomerErrorResponse(err))
		}
		result, err := h.service.Create(&businessRequest, c.Get("customerType"))
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.CustomerSuccessResponse(result))

	}
}

func (h *CustomerHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		customerId, _ := c.ParamsInt("id")
		if c.Get("customerType") == "individual" {
			var individualRequest models.IndividualCustomerUpdate
			err := c.BodyParser(&individualRequest)
			if err != nil {
				c.Status(fiber.StatusBadRequest)
				return c.JSON(presenters.CustomerErrorResponse(err))
			}
			result, err := h.service.Update(customerId, &individualRequest)
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(presenters.CustomerErrorResponse(err))
			}
			return c.JSON(presenters.CustomerSuccessResponse(result))
		}

		var businessRequest models.BusinessCustomerUpdate
		err := c.BodyParser(&businessRequest)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.CustomerErrorResponse(err))
		}
		result, err := h.service.Update(customerId, &businessRequest)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.CustomerSuccessResponse(result))
	}

}
func (h *CustomerHandler) UpdateBusinessLogo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logo, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		logoPath, err := h.storageSrv.Disk("uploadcare").UploadFile("customer_logo", logo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		customerId, _ := strconv.Atoi(c.Query("id"))
		prevUrl := c.Query("file", "")
		result, err := h.service.UpdateLogo(customerId, logoPath)
		if err != nil {
			if prevUrl != "" {
				h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		if prevUrl != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"data":   result,
			},
		)
	}

}
func (h *CustomerHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
