package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Jeffail/gabs"
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/order"
	"github.com/SeyramWood/app/application/payment"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
)

type PaystackHandler struct {
	service   gateways.PaymentService
	orderServ gateways.OrderService
}

func NewPaystackHandler(db *database.Adapter, logis gateways.LogisticService) *PaystackHandler {
	service := payment.NewPaymentService(payment.NewPaymentRepo(db))
	serv := order.NewOrderService(order.NewOrderRepo(db), logis)
	return &PaystackHandler{
		service:   service,
		orderServ: serv,
	}
}

func (h *PaystackHandler) InitiateTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request services.PaystackPayload

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		response, err := h.service.Pay(request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}
		res := response.(*http.Response)
		defer res.Body.Close()
		return c.JSON(presenters.PaystackInitiateTransactionResponse(res))
	}
}

func (h *PaystackHandler) VerifyTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {

		response, err := h.service.Verify(c.Params("reference"))
		res := response.(*http.Response)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}

		body, bErr := ioutil.ReadAll(res.Body)
		if bErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(bErr))
		}
		defer res.Body.Close()

		resBody, err := gabs.ParseJSON(body)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"res": map[string]interface{}{
					"status":  resBody.Path("status").Data(),
					"message": resBody.Path("message").Data(),
				},
			},
		)

	}
}

func (h *PaystackHandler) SaveOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		orderData, err := h.service.FormatPayload(c.Body())
		fmt.Println(orderData)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		_, err = h.orderServ.Create(orderData)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		return nil

	}
}
