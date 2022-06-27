package handlers

import (
	"errors"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/pickup_station"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type PickupStationHandler struct {
	service gateways.PickupStationService
}

func NewPickupStationHandler(db *database.Adapter) *PickupStationHandler {
	return &PickupStationHandler{
		service: pickup_station.NewPickupStationService(pickup_station.NewPickupStationRepo(db)),
	}
}

func (h *PickupStationHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//storeId, _ := c.ParamsInt("storeId")
		//
		//result, err := h.service.Fetch(storeId)
		//
		//if err != nil {
		//	return c.Status(fiber.StatusInternalServerError).JSON(presenters.MerchantStoreErrorResponse(err))
		//}
		//return c.JSON(presenters.MerchantStoreSuccessResponse(result))
		return nil

	}
}

func (h *PickupStationHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		results, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PickupErrorResponse(err))
		}
		return c.JSON(presenters.PickupSuccessResponses(results))
		
	}
}

func (h *PickupStationHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.PickupStation

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}

		result, err := h.service.Create(&request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PickupErrorResponse(errors.New("error creating merchant")))
		}
		return c.JSON(presenters.PickupSuccessResponse(result))

	}
}

func (h *PickupStationHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return nil
	}

}
func (h *PickupStationHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//if err := h.service.Remove(c.Params("id")); err != nil {
		//	return c.Status(fiber.StatusNotFound).JSON(presenters.MerchantErrorResponse(err))
		//}
		//return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		//	"status": true,
		//	"error":  nil,
		//})
		return nil
	}
}
