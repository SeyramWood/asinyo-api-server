package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/merchant_store"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
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
		storeId, _ := c.ParamsInt("storeId")

		result, err := h.service.Fetch(storeId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantStoreErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoreSuccessResponse(result))

	}
}
func (h *MerchantStoreHandler) FetchByMerchantID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		merchantId, _ := c.ParamsInt("merchantId")

		result, err := h.service.FetchByMerchant(merchantId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantStoreErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoreSuccessResponse(result))

	}
}
func (h *MerchantStoreHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		merchantType := c.Params("merchantType")
		result, err := h.service.FetchAllByMerchant(merchantType, limit, offset)
		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStorefrontsSuccessResponse(result))
	}
}

func (h *MerchantStoreHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantStore
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		merchantId, _ := c.ParamsInt("merchantId")

		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}

		logo, images, upErr := storage.NewUploadCare().Client().UploadMerchantStore(file, form)
		if upErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": upErr,
				},
			)
		}
		result, err := h.service.Create(&request, merchantId, logo, images)
		if err != nil {
			// Delete all files from remote server
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

func (h *MerchantStoreHandler) SaveDefaultAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		storeId, _ := c.ParamsInt("storeId")
		accountType := c.Params("type")

		result, err := h.service.SaveDefaultAccount(storeId, accountType)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.MerchantStorefrontSuccessResponse(result))

	}
}
func (h *MerchantStoreHandler) SaveAgentPermission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		storeId, _ := c.ParamsInt("storeId")
		permission, _ := strconv.ParseBool(c.Params("permission"))

		_, err := h.service.SaveAgentPermission(permission, storeId)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}

		return c.JSON(fiber.Map{"status": true, "msg": "Permission"})

	}
}

func (h *MerchantStoreHandler) FetchMerchantAgent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		storeId, _ := c.ParamsInt("storeId")

		result, err := h.service.FetchAgent(storeId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoreAgentSuccessResponse(result))

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
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
