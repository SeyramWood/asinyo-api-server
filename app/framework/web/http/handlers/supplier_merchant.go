package handlers

import (
	"errors"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/supplier_merchant"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type SupplierMerchantHandler struct {
	service gateways.SupplierMerchantService
}

func NewSupplierMerchantHandler(db *database.Adapter) *SupplierMerchantHandler {
	repo := supplier_merchant.NewSupplierMerchantRepo(db)
	service := supplier_merchant.NewSupplierMerchantService(repo)

	return &SupplierMerchantHandler{
		service: service,
	}
}

func (h *SupplierMerchantHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.SupplierMerchantErrorResponse(err))
		}
		return c.JSON(presenters.SupplierMerchantSuccessResponse(result))
	}
}
func (h *SupplierMerchantHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.SupplierMerchantErrorResponse(err))
		}
		return c.JSON(presenters.SupplierMerchantsSuccessResponse(result))
	}
}

func (h *SupplierMerchantHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.MerchantRequest

		err := c.BodyParser(&request)

		if err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(presenters.SupplierMerchantErrorResponse(err))
		}
		_, err = h.service.Create(&models.SupplierMerchant{
			GhanaCard:      request.Info.GhanaCard,
			LastName:       request.Info.LastName,
			OtherName:      request.Info.OtherName,
			Phone:          request.Info.Phone,
			OtherPhone:     request.Info.OtherPhone,
			Address:        request.Info.Address,
			DigitalAddress: request.Info.DigitalAddress,
		})

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.SupplierMerchantErrorResponse(errors.New("error creating agent")))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}

func (h *SupplierMerchantHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.SupplierMerchantErrorResponse(err))
		}
		return c.JSON(presenters.SupplierMerchantsSuccessResponse(result))
	}

}
func (h *SupplierMerchantHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.SupplierMerchantErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
