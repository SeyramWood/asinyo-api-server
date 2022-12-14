package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/customer"
	"github.com/SeyramWood/app/application/storage"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type CustomerHandler struct {
	service    gateways.CustomerService
	storageSrv gateways.StorageService
}

func NewCustomerHandler(db *database.Adapter) *CustomerHandler {
	repo := customer.NewCustomerRepo(db)
	service := customer.NewCustomerService(repo)

	return &CustomerHandler{
		service:    service,
		storageSrv: storage.NewStorageService(),
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

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.CustomersSuccessResponse(result))
	}

}
func (h *CustomerHandler) UpdateBusinessLogo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request struct {
			Logo []byte `json:"logo" form:"logo"`
		}
		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		logo, err := c.FormFile("logo")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		logoPath, _ := h.storageSrv.Disk("uploadcare").UploadFile("customer", logo)
		fmt.Println(logoPath)
		return nil
		// return c.JSON(presenters.CustomersSuccessResponse(result))
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
