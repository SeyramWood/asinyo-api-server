package handlers

import (
	"errors"
	"fmt"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/merchant_store"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
	"github.com/gofiber/fiber/v2"
)

type MerchantStoreHandler struct {
	service gateways.MerchantStoreService
}

func NewMerchantStoreHandler(db *database.Adapter) *MerchantStoreHandler {
	return &MerchantStoreHandler{
		service: merchant_store.NewMerchantStoreService(merchant_store.NewMerchantStoreRepo(db)),
	}
}

func (h *MerchantStoreHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("merchantId")

		result, err := h.service.Fetch(id)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantStoreErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoreSuccessResponse(result))
	}
}
func (h *MerchantStoreHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoresSuccessResponse(result))
	}
}

func (h *MerchantStoreHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantStore
		merchantId, _ := c.ParamsInt("merchantId")
		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		dir := fmt.Sprintf("merchant/stores/%s", c.Params("merchantId"))

		logo, logoErr := storage.NewStorage().SaveFile(c, "image", dir)

		if logoErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Upload error",
			})
		}
		images, imagesErr := storage.NewStorage().SaveFiles(c, "otherImages", dir)

		if imagesErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Upload error",
			})
		}

		result, err := h.service.Create(&request, merchantId, logo.(string), images.([]string))

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.MerchantStoreSuccessResponse(result))
	}
}
func (h *MerchantStoreHandler) SaveMomoAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantMomoAccountRequest

		storeId, _ := c.ParamsInt("storeId")

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		result, err := h.service.SaveAccount(&request, storeId, "momo")

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.MerchantStorefrontSuccessResponse(result))
	}
}
func (h *MerchantStoreHandler) SaveBankAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantBankAccountRequest
		storeId, _ := c.ParamsInt("storeId")
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		result, err := h.service.SaveAccount(&request, storeId, "bank")

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.MerchantStorefrontSuccessResponse(result))
	}
}

func (h *MerchantStoreHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoresSuccessResponse(result))
	}

}
func (h *MerchantStoreHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
