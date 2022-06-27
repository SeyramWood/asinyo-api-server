package presenters

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type paystackInitiateTransaction struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationUrl string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

func PaystackInitiateTransactionResponse(res *http.Response) *fiber.Map {
	var response paystackInitiateTransaction
	err := json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return errorResponse(err)
	}
	return successResponse(&response)

}
func PaymentErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
