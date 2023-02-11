package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/product_cat_major"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type ProductCatMajorHandler struct {
	service    gateways.ProductCatMajorService
	storageSrv gateways.StorageService
}

func NewProductCatMajorHandler(db *database.Adapter, storageSrv gateways.StorageService) *ProductCatMajorHandler {
	repo := product_cat_major.NewProductCatMajorRepo(db)
	service := product_cat_major.NewProductCatMajorService(repo)

	return &ProductCatMajorHandler{
		service:    service,
		storageSrv: storageSrv,
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
		var request models.ProductCategoryMajor

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		_, err = h.service.Update(id, &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductCatMajorErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}

}

func (h *ProductCatMajorHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.ProductCatMinorErrorResponse(err))
		}
		if c.Query("files") != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(strings.Split(c.Query("files"), "&"), "delete_files")
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
