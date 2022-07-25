package api

import (
	"fmt"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/order"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	service gateways.OrderService
}

func NewOrderHandler(db *database.Adapter) *OrderHandler {
	service := order.NewOrderService(order.NewOrderRepo(db))
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) FetchByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		usrId, _ := c.ParamsInt("id")
		results, err := h.service.FetchByAllUser(userType, usrId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.OrdersSuccessResponse(results))
	}
}
func (h *OrderHandler) FetchById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		orderId, _ := c.ParamsInt("id")

		result, err := h.service.Fetch(orderId)
		if err != nil {
			fmt.Println(err)

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.OrderSuccessResponse(result))
	}
}

func (h *OrderHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.OrdersSuccessResponse(result))
	}
}
