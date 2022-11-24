package models

import "github.com/SeyramWood/app/domain/services"

type (
	OrderPayloadMetadata struct {
		User           int                        `json:"user"`
		Pickup         int                        `json:"pickup"`
		Address        int                        `json:"address"`
		OrderNumber    string                     `json:"orderNumber"`
		DeliveryFee    float64                    `json:"deliveryFee"`
		UserType       string                     `json:"userType"`
		DeliveryMethod string                     `json:"deliveryMethod"`
		PaymentMethod  string                     `json:"paymentMethod"`
		Products       []*services.ProductDetails `json:"products"`
	}
	OrderPayload struct {
		Amount    float64               `json:"amount"`
		Reference string                `json:"reference"`
		Currency  string                `json:"currency"`
		Channel   string                `json:"channel"`
		PaidAt    string                `json:"paidAt"`
		Metadata  *OrderPayloadMetadata `json:"metadata"`
	}
	OrderFareEstimateRequest struct {
		Delivery *services.Coordinate   `json:"delivery,omitempty"`
		Pickups  []*services.Coordinate `json:"pickups,omitempty"`
	}
)
