package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/product"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
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
		return c.JSON(presenters.ProductWithMerchantResponse(result))
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
			return c.JSON(presenters.ProductWithMerchantResponse(result))
		}

		result, err := h.service.FetchBySupplierMerchant(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductWithMerchantResponse(result))
	}

}

func (h *ProductHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		result, err := h.service.FetchAll(limit, offset)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductsWithStoreResponse(result))
	}
}

func (h *ProductHandler) FetchBySlugMerchantCategoryProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		if c.Params("merchant") == "retailer" {
			if c.Params("cat") == "major" {
				products, err := h.service.FetchBySlugRetailMerchantCategoryMajor(c.Params("slug"), limit, offset)

				if err != nil {

					return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
				}
				return c.JSON(presenters.ProductsCategoryMajorWithMerchantResponse(products))
			}

			if c.Params("cat") == "minor" {
				products, err := h.service.FetchBySlugRetailMerchantCategoryMinor(c.Params("slug"), limit, offset)

				if err != nil {

					return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
				}
				return c.JSON(presenters.ProductsCategoryMinorWithMerchantResponse(products))
			}

		}
		if c.Params("merchant") == "supplier" {
			if c.Params("cat") == "major" {
				products, err := h.service.FetchBySlugSupplierMerchantCategoryMajor(c.Params("slug"), limit, offset)

				if err != nil {

					return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
				}
				return c.JSON(presenters.ProductsCategoryMajorWithMerchantResponse(products))
			}

			if c.Params("cat") == "minor" {
				products, err := h.service.FetchBySlugSupplierMerchantCategoryMinor(c.Params("slug"), limit, offset)

				if err != nil {

					return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
				}
				return c.JSON(presenters.ProductsCategoryMinorWithMerchantResponse(products))
			}

		}
		return c.Status(fiber.StatusNotFound).JSON("Not Found")

	}
}

func (h *ProductHandler) FetchAllBySlugMerchantCategoryProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		if c.Params("cat") == "major" {
			products, err := h.service.FetchAllBySlugCategoryMajor(
				c.Params("merchant"), c.Params("slug"), limit, offset,
			)

			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}

			return c.JSON(presenters.ProductsWithMerchantResponse(products))
		}

		if c.Params("cat") == "minor" {
			products, err := h.service.FetchAllBySlugCategoryMinor(
				c.Params("merchant"), c.Params("slug"), limit, offset,
			)

			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}

			return c.JSON(presenters.ProductsWithMerchantResponse(products))
		}

		return c.Status(fiber.StatusNotFound).JSON("Not Found")

	}
}

func (h *ProductHandler) FetchAllMerchantCategoryMajorProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		if c.Params("merchant") == "retailer" {

			products, err := h.service.FetchAllRetailMerchantCategoryMajor(limit, offset)

			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsCategoryMajorWithMerchantResponse(products))
		}

		if c.Params("merchant") == "supplier" {

			products, err := h.service.FetchAllSupplierMerchantCategoryMajor(limit, offset)

			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}

			return c.JSON(presenters.ProductsCategoryMajorWithMerchantResponse(products))
		}

		return c.Status(fiber.StatusNotFound).JSON("Not Found")

	}
}

func (h *ProductHandler) FetchMerchantProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		id, _ := c.ParamsInt("id")

		if c.Params("merchant") == "supplier" {
			products, err := h.service.FetchAllBySupplier(id, limit, offset)
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsWithMerchantResponse(products))
		}
		if c.Params("merchant") == "retailer" {
			products, err := h.service.FetchAllByRetailer(id, limit, offset)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsWithMerchantResponse(products))
		}
		return c.Status(fiber.StatusNotFound).JSON("Not Found")
	}
}

func (h *ProductHandler) FetchMerchantBestSellerProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		id, _ := c.ParamsInt("id")

		products, err := h.service.FetchBestSellerByMerchant(id, limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
		}
		return c.JSON(presenters.ProductsWithMerchantResponse(products))
		
	}
}

func (h *ProductHandler) FetchBestSellerProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		if c.Params("merchantType") == "supplier" {
			products, err := h.service.FetchBestSellerBySupplier(limit, offset)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsWithMerchantResponse(products))
		}

		if c.Params("merchantType") == "retailer" {
			products, err := h.service.FetchBestSellerByRetailer(limit, offset)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(err))
			}
			return c.JSON(presenters.ProductsWithMerchantResponse(products))
		}

		return c.Status(fiber.StatusNotFound).JSON("Not Found")
	}
}

func (h *ProductHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Product

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}

		fPath, err := storage.NewUploadCare().Client().Upload(file, "product")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}

		result, err := h.service.Create(&request, fPath)

		if err != nil {
			// Delete file from remote server
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error creating merchant")))
		}

		if c.Params("merchant") == "supplier" {
			prod, err := h.service.FetchBySupplierMerchant(result.ID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error fetching supplier merchant product")))
			}
			return c.JSON(presenters.ProductWithMerchantResponse(prod))
		}

		if c.Params("merchant") == "retailer" {
			prod, err := h.service.FetchByRetailMerchant(result.ID)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.ProductErrorResponse(errors.New("error fetching supplier merchant product")))
			}

			return c.JSON(presenters.ProductWithMerchantResponse(prod))
		}

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"msg": "Bad Request",
			},
		)
	}
}

func (h *ProductHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		result, err := h.service.FetchAll(limit, offset)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.ProductsWithStoreResponse(result))
	}

}
func (h *ProductHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.ProductErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
