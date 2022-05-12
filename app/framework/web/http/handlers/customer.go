package handlers

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/customer"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service gateways.CustomerService
}

func NewCustomerHandler(db *database.Adapter) *CustomerHandler {
	repo := customer.NewCustomerRepo(db)
	service := customer.NewCustomerService(repo)

	return &CustomerHandler{
		service: service,
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

		var request models.Customer

		err := c.BodyParser(&request)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.CustomerErrorResponse(err))
		}

		result, err := h.service.Create(&request)

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
func (h *CustomerHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
