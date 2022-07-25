package api

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/product_cat_major"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type ProductCatMajorHandler struct {
	service gateways.ProductCatMajorService
}

func NewProductCatMajorHandler(db *database.Adapter) *ProductCatMajorHandler {
	repo := product_cat_major.NewProductCatMajorRepo(db)
	service := product_cat_major.NewProductCatMajorService(repo)

	return &ProductCatMajorHandler{
		service: service,
	}
}

func (h *ProductCatMajorHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMajorSuccessResponse(result))
	}
}
func (h *ProductCatMajorHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		results, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMajorsSuccessResponse(results))
	}
}

func (h *ProductCatMajorHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.ProductCategoryMajor

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		cat, err := h.service.Create(&request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMajorSuccessResponse(cat))
	}

}

func (h *ProductCatMajorHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMajorsSuccessResponse(result))
	}

}
func (h *ProductCatMajorHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
