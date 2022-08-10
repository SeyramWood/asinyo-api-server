package api

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/order"
	"github.com/SeyramWood/app/framework/database"
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
		results, err := h.service.FetchAllByUser(userType, usrId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.OrdersSuccessResponse(results))
	}
}

func (h *OrderHandler) FetchAllByStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		merchantId, _ := c.ParamsInt("merchant")

		results, err := h.service.FetchAllByStore(merchantId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(
			presenters.StoreOrdersSuccessResponse(results),
		)
	}
}
func (h *OrderHandler) FetchByStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		merchantId, _ := c.ParamsInt("merchant")
		orderId, _ := c.ParamsInt("order")

		result, err := h.service.FetchByStore(orderId, merchantId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.StoreOrderSuccessResponse(result))
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

func (h *OrderHandler) SaveOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {

		orderData, err := h.service.FormatOrderRequest(c.Body())

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		if strings.Compare(orderData.Event, "charge.success") == 0 {
			o, err := h.service.Create(orderData)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
			}
			return c.JSON(presenters.OrderSuccessResponse(o))
		}

		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"msg": "Bad Request",
			},
		)
	}
}

func (h *OrderHandler) UpdateOrderDetailStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.UpdateOrderDetailStatus(c.Body())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.StoreOrderSuccessResponse(result))
	}
}
