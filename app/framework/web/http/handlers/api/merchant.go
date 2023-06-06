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
	"github.com/SeyramWood/app/application/merchant"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
)

type MerchantHandler struct {
	service gateways.MerchantService
	maps    gateways.MapService
}

func NewMerchantHandler(
	db *database.Adapter, noti notification.NotificationService, maps gateways.MapService,
) *MerchantHandler {
	repo := merchant.NewMerchantRepo(db)
	service := merchant.NewMerchantService(repo, noti)
	mapService := maps.SetMerchantRepo(repo)
	return &MerchantHandler{
		service: service,
		maps:    mapService,
	}
}

func (h *MerchantHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantSuccessResponse(result))
	}
}

func (h *MerchantHandler) FetchStorefrontByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.FetchStorefront(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantStoreSuccessResponse(result))
	}
}

func (h *MerchantHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		result, err := h.service.FetchAll(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantsSuccessResponse(result))
	}
}

func (h *MerchantHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		_, err = h.service.Create(&request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}

func (h *MerchantHandler) OnboardMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.MerchantStoreInfo

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		agentId, _ := c.ParamsInt("agent")
		file, err := c.FormFile("banner")
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
		cache := cacheDriver.New("./mnt/cache/")
		stepOne := fmt.Sprintf("step_one_%s", c.Params("agent"))
		stepTwo := fmt.Sprintf("step_two_%s", c.Params("agent"))
		cachedData := cache.FetchMulti([]string{stepOne, stepTwo})
		if len(cachedData) < 2 {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Something went wrong",
				},
			)
		}
		pInfo := models.RetailerStorePersonalInfo{}
		aInfo := models.MerchantStoreAddress{}
		err = json.Unmarshal([]byte(cachedData[stepOne]), &pInfo)
		err = json.Unmarshal([]byte(cachedData[stepTwo]), &aInfo)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Something went wrong",
				},
			)
		}

		requestData := models.OnboardMerchantFullRequest{
			PersonalInfo: &pInfo,
			Address:      &aInfo,
			StoreInfo:    &request,
		}

		logo, images, upErr := storage.NewUploadCare().Client().UploadMerchantStore(file, form)
		if upErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": upErr,
				},
			)
		}

		result, errr := h.service.Onboard(&requestData, agentId, logo, images)
		if errr != nil {
			// TODO Delete all files from remote server
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}

		for key := range cachedData {
			cache.Delete(key)
		}

		h.maps.ExecuteTask(result, "geocoding", "merchant")

		return c.JSON(presenters.EmptySuccessResponse())
	}
}

func (h *MerchantHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		merchantId, _ := c.ParamsInt("id")
		if c.Get("userType") == "retailer" {
			var request models.RetailerProfileUpdate
			err := c.BodyParser(&request)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
			}

			result, err := h.service.Update(merchantId, &request)
			if err != nil {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(presenters.MerchantErrorResponse(err))
			}
			return c.JSON(presenters.MerchantSuccessResponse(result))
		}

		var request models.SupplierProfileUpdate
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		result, err := h.service.Update(merchantId, &request)
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantSuccessResponse(result))
	}
}

func (h *MerchantHandler) Delete() fiber.Handler {
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
