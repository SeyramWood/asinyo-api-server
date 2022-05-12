package handlers

import (
	"errors"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/merchant"
	"github.com/SeyramWood/app/application/retail_merchant"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type RetailMerchantHandler struct {
	service  gateways.RetailMerchantService
	merchant gateways.MerchantService
}

func NewRetailMerchantHandler(db *database.Adapter) *RetailMerchantHandler {
	repo := retail_merchant.NewRetailMerchantRepo(db)
	mrepo := merchant.NewMerchantRepo(db)
	service := retail_merchant.NewRetailMerchantService(repo)
	merchant := merchant.NewMerchantService(mrepo)

	return &RetailMerchantHandler{
		service:  service,
		merchant: merchant,
	}
}

func (h *RetailMerchantHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RetailMerchantErrorResponse(err))
		}
		return c.JSON(presenters.RetailMerchantSuccessResponse(result))
	}
}
func (h *RetailMerchantHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RetailMerchantErrorResponse(err))
		}
		return c.JSON(presenters.RetailMerchantsSuccessResponse(result))
	}
}

func (h *RetailMerchantHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.MerchantRequest

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.RetailMerchantErrorResponse(err))
		}
		_, err = h.merchant.Create(&request)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RetailMerchantErrorResponse(errors.New("error creating agent")))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}

func (h *RetailMerchantHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RetailMerchantErrorResponse(err))
		}
		return c.JSON(presenters.RetailMerchantsSuccessResponse(result))
	}

}
func (h *RetailMerchantHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.RetailMerchantErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
