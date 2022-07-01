package handlers

import (
	"errors"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/product"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
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

func (h *ProductHandler) FetchByIDMerchantProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")

		if c.Params("merchant") == "retailer" {

			result, err := h.service.FetchByRetailMerchant(id)

			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductRetailerMerchantSuccessResponse(result))
		}

		result, err := h.service.FetchBySupplierMerchant(id)

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

func (h *ProductHandler) FetchBySlugMerchantCategoryProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Params("merchant") == "retailer" {
			if c.Params("cat") == "major" {
				products, err := h.service.FetchBySlugRetailMerchantCategoryMajor(c.Params("slug"))

				if err != nil {

					return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
				}
				return c.JSON(presenters.ProductsRetailerCategoryMajorResponse(products))
			}

			if c.Params("cat") == "minor" {
				products, err := h.service.FetchBySlugRetailMerchantCategoryMinor(c.Params("slug"))

				if err != nil {

					return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
				}
				return c.JSON(presenters.ProductsRetailerCategoryMinorResponse(products))
			}

		}

		products, err := h.service.FetchBySlugRetailMerchantCategoryMajor(c.Params("slug"))

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}

		return c.JSON(presenters.ProductsSupplierCategoryMajorResponse(products))

	}
}

func (h *ProductHandler) FetchAllMerchantCategoryMajorProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Params("merchant") == "retailer" {

			products, err := h.service.FetchAllRetailMerchantCategoryMajor()

			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsRetailerCategoryMajorResponse(products))
		}

		products, err := h.service.FetchAllSupplierMerchantCategoryMajor()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}

		return c.JSON(presenters.ProductsSupplierCategoryMajorResponse(products))

	}
}
func (h *ProductHandler) FetchMerchantProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, _ := c.ParamsInt("id")

		if c.Params("merchant") == "supplier" {
			products, err := h.service.FetchAllBySupplier(id)
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsSupplierResponse(products))
		}

		products, err := h.service.FetchAllByRetailer(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductsRetailerResponse(products))
	}
}

func (h *ProductHandler) FetchBestSellerProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		if c.Params("merchantType") == "supplier" {
			products, err := h.service.FetchBestSellerBySupplier()
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsSupplierResponse(products))
		}

		products, err := h.service.FetchBestSellerByRetailer()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductsRetailerResponse(products))
	}
}

func (h *ProductHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Product

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}

		path, err := storage.NewStorage().SaveFile(c, "image", "products")

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Bad request",
			})
		}

		result, err := h.service.Create(&request, path.(string))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error creating merchant")))
		}

		if c.Params("merchant") == "supplier" {
			prod, err := h.service.FetchBySupplierMerchant(result.ID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error fetching supplier merchant product")))
			}
			return c.JSON(presenters.ProductSupplierResponse(prod))
		}

		prod, err := h.service.FetchByRetailMerchant(result.ID)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error fetching supplier merchant product")))
		}

		return c.JSON(presenters.ProductRetailerResponse(prod))
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
