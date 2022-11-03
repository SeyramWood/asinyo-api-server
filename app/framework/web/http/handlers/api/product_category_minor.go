package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/product_cat_minor"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
)

type ProductCatMinorHandler struct {
	service gateways.ProductCatMinorService
}

func NewProductCatMinorHandler(db *database.Adapter) *ProductCatMinorHandler {
	repo := product_cat_minor.NewProductCatMinorRepo(db)
	service := product_cat_minor.NewProductCatMinorService(repo)

	return &ProductCatMinorHandler{
		service: service,
	}
}

func (h *ProductCatMinorHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMinorSuccessResponse(result))
	}
}
func (h *ProductCatMinorHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		results, err := h.service.FetchAll(limit, offset)
		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMinorErrorResponse(err))
		}

		return c.JSON(presenters.ProductCatMinorsSuccessResponse(results))
	}
}

func (h *ProductCatMinorHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.ProductCategoryMinor

		err := c.BodyParser(&request)

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMinorErrorResponse(err))
		}

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error.",
				},
			)
		}
		fPath, err := storage.NewUploadCare().Client().Upload(file, "category_minor")

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error.",
				},
			)
		}
		cat, err := h.service.Create(&request, fPath)
		if err != nil {
			// Delete file from remote storage
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMinorSuccessResponse(cat))
	}
}

func (h *ProductCatMinorHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		result, err := h.service.FetchAll(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		return c.JSON(presenters.ProductCatMinorsSuccessResponse(result))
	}

}
func (h *ProductCatMinorHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
