package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/order"
	"github.com/SeyramWood/app/application/payment"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type OrderHandler struct {
	service        gateways.OrderService
	paymentService gateways.PaymentService
	logis          gateways.LogisticService
}

func NewOrderHandler(db *database.Adapter, logis gateways.LogisticService, mail gateways.EmailService) *OrderHandler {
	service := order.NewOrderService(order.NewOrderRepo(db), logis, mail)
	paymentService := payment.NewPaymentService(payment.NewPaymentRepo(db), "pay_on_delivery")
	return &OrderHandler{
		service:        service,
		paymentService: paymentService,
		logis:          logis,
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
func (h *OrderHandler) FetchAllByAgentStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		agentId, _ := c.ParamsInt("agent")

		results, err := h.service.FetchAllByAgentStore(agentId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(
			presenters.AgentStoreOrdersSuccessResponse(results),
		)
	}
}
func (h *OrderHandler) FetchByStore() fiber.Handler {
	return func(c *fiber.Ctx) error {
		merchantId, _ := c.ParamsInt("merchant")
		orderId, _ := c.ParamsInt("order")
		userType := c.Get("userType")
		result, err := h.service.FetchByStore(orderId, merchantId, userType)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		if userType == "agent" {
			return c.JSON(presenters.AgentStoreOrderSuccessResponse(result))
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

		orderData, err := h.paymentService.FormatPayload(c.Body())

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		_, err = h.service.Create(orderData)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"data": map[string]any{
					"status": true,
					"msg":    "Order successfully placed",
				},
			},
		)
	}
}
func (h *OrderHandler) TestOrderCreation() fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := struct {
			ID int `json:"id"`
		}{}
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
		}
		_, err = h.service.TestCreate(request.ID)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"data": map[string]any{
					"status": true,
					"msg":    "Order successfully placed",
				},
			},
		)
	}
}

func (h *OrderHandler) ListenTookanWebhook() fiber.Handler {
	return func(c *fiber.Ctx) error {
		h.logis.ExecuteWebhook(c.Body())
		return nil
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

func (h *OrderHandler) FetchOrderFareEstimate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var coordinates models.OrderFareEstimateRequest

		if err := c.BodyParser(&coordinates); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.OrderErrorResponse(err))
		}

		result, err := h.logis.FareEstimate(&coordinates)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.OrderErrorResponse(err))
		}
		return c.JSON(presenters.OrderFareEstimateSuccessResponse(result))
	}
}
