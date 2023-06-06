package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/price_model"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type PriceModelHandler struct {
	service gateways.PriceModelService
}

func NewPriceModelHandler(db *database.Adapter) *PriceModelHandler {
	service := price_model.NewPriceModelService(price_model.NewPriceModelRepo(db))
	return &PriceModelHandler{
		service: service,
	}
}

func (h *PriceModelHandler) Create() fiber.Handler {

	return func(c *fiber.Ctx) error {
		var request models.PriceModelRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PriceModelErrorResponse(err))
		}
		response, err := h.service.Create(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.JSON(presenters.PriceModelSuccessResponse(response))
	}
}

func (h *PriceModelHandler) CreatePercentage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		percentage, _ := strconv.Atoi(c.Query("percentage", "0"))
		category, _ := c.ParamsInt("category")
		response, err := h.service.UpdatePercentage(category, percentage)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.JSON(presenters.PriceModelPercentageSuccessResponse(response))
	}
}

func (h *PriceModelHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		model, _ := c.ParamsInt("model")
		response, err := h.service.Fetch(model)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.JSON(presenters.PriceModelSuccessResponse(response))
	}
}
func (h *PriceModelHandler) FetchAll() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		response, err := h.service.FetchAll(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.JSON(presenters.PriceModelsSuccessResponse(response))
	}
}
func (h *PriceModelHandler) FetchAllPercentage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		response, err := h.service.FetchAllPercentage(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.JSON(presenters.PriceModelPercentagesSuccessResponse(response))
	}
}

func (h *PriceModelHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.PriceModelRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PriceModelErrorResponse(err))
		}
		model, _ := c.ParamsInt("model")
		response, err := h.service.Update(model, &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.JSON(presenters.PriceModelSuccessResponse(response))
	}
}

func (h *PriceModelHandler) Remove() fiber.Handler {
	return func(c *fiber.Ctx) error {
		model, _ := c.ParamsInt("model")
		if err := h.service.Remove(model); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}

func (h *PriceModelHandler) RemovePercentage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		category, _ := c.ParamsInt("category")
		if err := h.service.RemovePercentage(category); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PriceModelErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
