package handlers

import (
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/order"
	"github.com/SeyramWood/app/application/payment"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"strings"
	"sync"
)

type PaymentHandler struct {
	service   gateways.PaymentService
	orderServ gateways.OrderService
}

func NewPaystackHandler(db *database.Adapter) *PaymentHandler {
	service := payment.NewPaymentService(payment.NewPaymentRepo(db), "paystack")
	serv := order.NewOrderService(order.NewOrderRepo(db))
	return &PaymentHandler{
		service:   service,
		orderServ: serv,
	}
}

func (h *PaymentHandler) InitiateTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.OrderRequest

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}

		response, err := h.service.Pay(request)

		defer response.Body.Close()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}

		return c.JSON(presenters.PaystackInitiateTransactionResponse(response))
	}
}

func (h *PaymentHandler) VerifyTransaction() fiber.Handler {
	return func(c *fiber.Ctx) error {

		response, err := h.service.Verify(c.Params("reference"))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}

		body, bErr := ioutil.ReadAll(response.Body)
		if bErr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(bErr))
		}
		defer response.Body.Close()

		resBody, err := gabs.ParseJSON(body)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.PaymentErrorResponse(err))
		}

		return c.JSON(fiber.Map{"res": map[string]interface{}{
			"status":  resBody.Path("status").Data(),
			"message": resBody.Path("message").Data(),
		}})

	}
}

func (h *PaymentHandler) WebhookResponse() fiber.Handler {
	return func(c *fiber.Ctx) error {

		resBody, err := gabs.ParseJSON(c.Body())

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.PaymentErrorResponse(err))
		}

		response := models.OrderResponse{
			Event:     resBody.Path("event").Data().(string),
			Amount:    resBody.Path("data.amount").Data().(float64),
			Currency:  resBody.Path("data.currency").Data().(string),
			Channel:   resBody.Path("data.channel").Data().(string),
			Reference: resBody.Path("data.reference").Data().(string),
			PaidAt:    resBody.Path("data.paid_at").Data().(string),
			MetaData: &models.OrderResponseMetadata{
				User:           resBody.Path("data.metadata.user").Data().(string),
				UserType:       resBody.Path("data.metadata.userType").Data().(string),
				OrderNumber:    resBody.Path("data.metadata.orderNumber").Data().(string),
				Address:        resBody.Path("data.metadata.address").Data().(string),
				DeliveryMethod: resBody.Path("data.metadata.deliveryMethod").Data().(string),
				DeliveryFee:    resBody.Path("data.metadata.deliveryFee").Data().(string),
				Pickup:         resBody.Path("data.metadata.pickup").Data().(string),
				Products: func() []*models.ProductDetailsResponse {
					var products []*models.ProductDetailsResponse
					children, _ := resBody.Path("data.metadata.products").Children()
					wg := sync.WaitGroup{}
					for _, child := range children {
						wg.Add(1)
						go func(child *gabs.Container) {
							defer wg.Done()
							pro := child.Data().(map[string]interface{})
							products = append(products, &models.ProductDetailsResponse{
								ID:         pro["id"].(string),
								Store:      pro["store"].(string),
								Quantity:   pro["quantity"].(string),
								Price:      pro["price"].(string),
								PromoPrice: pro["promoPrice"].(string),
							})
						}(child)
					}
					wg.Wait()
					return products
				}(),
			},
		}

		if strings.Compare(response.Event, "charge.success") == 0 {
			if err := h.orderServ.Create(&response); err != nil {
				fmt.Println("Ooops! Error while creating order\nERROR:\n", err)
			}
		}

		return nil
	}
}
