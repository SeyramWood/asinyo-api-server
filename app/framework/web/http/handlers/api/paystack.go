package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/order"
	"github.com/SeyramWood/app/application/payment"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type PaystackHandler struct {
	service   gateways.PaymentService
	orderServ gateways.OrderService
}

func NewPaystackHandler(db *database.Adapter) *PaystackHandler {
	service := payment.NewPaymentService(payment.NewPaymentRepo(db), "paystack")
	serv := order.NewOrderService(order.NewOrderRepo(db))
	return &PaystackHandler{
		service:   service,
		orderServ: serv,
	}
}

func (h *PaystackHandler) InitiateTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.OrderRequest

		err := c.BodyParser(&request)

		if err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}
		response, err := h.service.Pay(request)
		res := response.(*http.Response)
		defer res.Body.Close()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}

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

		orderData, err := h.orderServ.FormatOrderRequest(c.Body())

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}

		if strings.Compare(orderData.Event, "charge.success") == 0 {
			if _, err := h.orderServ.Create(orderData); err != nil {
				fmt.Println("Ooops! Error while creating order\nERROR:\n", err)
			}
			return c.Status(fiber.StatusOK).JSON(
				fiber.Map{
					"msg": "Order saved successfully",
				},
			)
		}
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"msg": "Bad Request",
			},
		)
	}
}
