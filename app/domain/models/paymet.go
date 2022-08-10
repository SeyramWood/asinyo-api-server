package models

type (
	ProductDetails struct {
		ID         int     `json:"id"`
		Store      int     `json:"store"`
		Quantity   int     `json:"quantity"`
		Price      float64 `json:"price"`
		PromoPrice float64 `json:"promoPrice"`
	}
	ProductDetailsResponse struct {
		ID         string `json:"id"`
		Store      string `json:"store"`
		Quantity   string `json:"quantity"`
		Price      string `json:"price"`
		PromoPrice string `json:"promoPrice"`
	}
	CustomFields struct {
		DisplayName  string `json:"display_name"`
		VariableName string `json:"variable_name"`
		Value        string `json:"value"`
	}
	OrderRequestMetadata struct {
		User           int               `json:"user"`
		UserType       string            `json:"userType"`
		OrderNumber    string            `json:"orderNumber"`
		Address        int               `json:"address"`
		DeliveryMethod string            `json:"deliveryMethod"`
		DeliveryFee    float64           `json:"deliveryFee"`
		Pickup         int               `json:"pickup"`
		Products       []*ProductDetails `json:"products"`
		CustomFields   []*CustomFields   `json:"custom_fields"`
	}
	OrderResponseMetadata struct {
		User           string                    `json:"user"`
		UserType       string                    `json:"userType"`
		OrderNumber    string                    `json:"orderNumber"`
		Address        string                    `json:"address"`
		DeliveryMethod string                    `json:"deliveryMethod"`
		PaymentMethod  string                    `json:"paymentMethod"`
		DeliveryFee    string                    `json:"deliveryFee"`
		Pickup         string                    `json:"pickup"`
		Products       []*ProductDetailsResponse `json:"products"`
	}
	OrderRequest struct {
		Amount      float64               `json:"amount" validate:"required"`
		Email       string                `json:"email" validate:"required|string"`
		Currency    string                `json:"currency" validate:"required|string"`
		Reference   string                `json:"reference" validate:"required|string"`
		CallbackUrl string                `json:"callback_url" validate:"required|string"`
		MetaData    *OrderRequestMetadata `json:"metadata"`
	}

	OrderResponse struct {
		Event     string                 `json:"event"`
		Amount    float64                `json:"amount"`
		Currency  string                 `json:"currency"`
		Channel   string                 `json:"channel"`
		Reference string                 `json:"reference"`
		PaidAt    string                 `json:"paidAt"`
		MetaData  *OrderResponseMetadata `json:"metadata"`
	}

	Transaction struct {
		Amount    float64 `json:"amount"`
		Reference string  `json:"reference"`
		Channel   string  `json:"channel"`
		Currency  string  `json:"currency"`
		PaidAt    string  `json:"paid_at"`
	}
)
