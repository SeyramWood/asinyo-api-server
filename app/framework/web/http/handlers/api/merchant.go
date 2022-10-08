package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/merchant"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/pkg/storage"
)

type MerchantHandler struct {
	service gateways.MerchantService
}

func NewMerchantHandler(db *database.Adapter, mail gateways.EmailService) *MerchantHandler {
	repo := merchant.NewMerchantRepo(db)
	service := merchant.NewMerchantService(repo, mail)

	return &MerchantHandler{
		service: service,
	}
}

func (h *MerchantHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RetailMerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantSuccessResponse(result))
	}
}
func (h *MerchantHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

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

		var request models.StoreFinalRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		agentId, _ := c.ParamsInt("agent")
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
		_, err = h.service.Onboard(&request, agentId, logo, images)
		if err != nil {
			// Delete all files from remote server
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}
func (h *MerchantHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantErrorResponse(err))
		}
		return c.JSON(presenters.MerchantsSuccessResponse(result))
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
