package handlers

import (
	"errors"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/product"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service gateways.ProductService
}

func NewProductHandler(db *database.Adapter) *ProductHandler {
	repo := product.NewProductRepo(db)
	service := product.NewProductService(repo)

	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductSuccessResponse(result))
	}
}
func (h *ProductHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductsSuccessResponse(result))
	}
}

func (h *ProductHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Product

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}

		_, err = h.service.Create(&request)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}

func (h *ProductHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.ProductsSuccessResponse(result))
	}

}
func (h *ProductHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.ProductErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
