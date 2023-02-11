package api

import (
	"errors"
	"fmt"
	"strconv"

	cacheDriver "github.com/faabiosr/cachego/file"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/merchant_store"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
)

type MerchantStoreHandler struct {
	service    gateways.MerchantStoreService
	maps       gateways.MapService
	storageSrv gateways.StorageService
}

func NewMerchantStoreHandler(
	db *database.Adapter, maps gateways.MapService, storageSrv gateways.StorageService,
) *MerchantStoreHandler {
	repo := merchant_store.NewMerchantStoreRepo(db)
	return &MerchantStoreHandler{
		service:    merchant_store.NewMerchantStoreService(repo),
		maps:       maps.SetMerchantStoreRepo(repo),
		storageSrv: storageSrv,
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
		var infoRequest models.MerchantStore
		err := c.BodyParser(&infoRequest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		cache := cacheDriver.New("./mnt/cache/")
		stepOne := fmt.Sprintf("step_one_%s", c.Params("merchantId"))
		if !cache.Contains(stepOne) {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Something went wrong",
				},
			)
		}
		cachedData, errr := cache.Fetch(stepOne)
		if errr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Something went wrong",
				},
			)
		}
		addressRequest := models.MerchantStoreAddress{}
		err = json.Unmarshal([]byte(cachedData), &addressRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Something went wrong",
				},
			)
		}
		requestData := models.MerchantStoreRequest{
			Info:    &infoRequest,
			Address: &addressRequest,
		}
		file, err := c.FormFile("banner")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		logo, images, upErr := storage.NewUploadCare().Client().UploadMerchantStore(file, form)
		if upErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": upErr,
				},
			)
		}

		merchantId, _ := c.ParamsInt("merchantId")
		result, err := h.service.Create(&requestData, merchantId, logo, images)
		if err != nil {
			// TODO Delete all files from remote server
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}
		cache.Delete(stepOne)
		h.maps.ExecuteTask(result, "geocoding", "store")
		return c.JSON(presenters.MerchantStoreSuccessResponse(result))
	}
}

func (h *MerchantStoreHandler) SaveAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var momoRequest models.MerchantMomoAccountRequest
		var bankRequest models.MerchantBankAccountRequest
		storeId, _ := c.ParamsInt("id")
		accountType := c.Params("accountType")
		if accountType == "bank" {
			err := c.BodyParser(&bankRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			result, err := h.service.SaveAccount(&bankRequest, storeId, accountType)
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
			}
			return c.JSON(presenters.MerchantStorefrontSuccessResponse(result))
		}

		err := c.BodyParser(&momoRequest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		result, err := h.service.SaveAccount(&momoRequest, storeId, accountType)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
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
		var infoRequest models.MerchantStoreUpdate
		var addressRequest models.MerchantStoreAddress

		formType := c.Get("formType")
		storeId, _ := c.ParamsInt("storeId")

		if formType == "information" {
			err := c.BodyParser(&infoRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			result, err := h.service.Update(&infoRequest, storeId)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
			}

			return c.JSON(presenters.MerchantStorefrontSuccessResponse(result))
		}
		if formType == "address" {
			err := c.BodyParser(&addressRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			result, errr := h.service.UpdateAddress(&addressRequest, storeId)
			if errr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
			}
			h.maps.ExecuteTask(result, "geocoding", "store")
			return c.JSON(presenters.MerchantStorefrontSuccessResponse(result))
		}

		if formType == "images" {
			form, err := c.MultipartForm()
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}
			urls, errr := h.storageSrv.Disk("uploadcare").UploadFiles("merchant_store", form.File["images"])
			if errr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
			}
			result, err := h.service.AppendNewImages(storeId, urls)
			if errr != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
			}
			return c.Status(fiber.StatusOK).JSON(
				fiber.Map{
					"status": true,
					"data":   result,
				},
			)

		}

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"status": false,
				"msg":    "Bad request",
			},
		)
	}

}
func (h *MerchantStoreHandler) UpdateBusinessBanner() fiber.Handler {
	return func(c *fiber.Ctx) error {
		banner, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		bannerPath, err := h.storageSrv.Disk("uploadcare").UploadFile("merchant_logo", banner)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}

		storeId, _ := strconv.Atoi(c.Query("id"))
		prevUrl := c.Query("file", "")
		result, err := h.service.UpdateBanner(storeId, bannerPath)
		if err != nil {
			if prevUrl != "" {
				h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		if prevUrl != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"data":   result,
			},
		)
	}

}
func (h *MerchantStoreHandler) UpdateBusinessImages() fiber.Handler {
	return func(c *fiber.Ctx) error {
		imageFile, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		newPath, err := h.storageSrv.Disk("uploadcare").UploadFile("merchant_store", imageFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}

		storeId, _ := strconv.Atoi(c.Query("id"))
		prevUrl := c.Query("file", "")
		result, err := h.service.UpdateImages(storeId, newPath, prevUrl)
		if err != nil {
			if prevUrl != "" {
				h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		if prevUrl != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"data":   result,
			},
		)
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
