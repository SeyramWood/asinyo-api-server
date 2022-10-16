package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	AuthAdmin struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		LastName  string `json:"lastName"`
		OtherName string `json:"otherName"`
	}
	AuthCustomer struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		LastName   string `json:"lastName"`
		OtherName  string `json:"otherName"`
		Phone      string `json:"phone"`
		OtherPhone string `json:"otherPhone"`
		UserType   string `json:"userType"`
	}
	AuthMerchant struct {
		ID         int    `json:"id"`
		Username   string `json:"username"`
		LastName   string `json:"lastName"`
		OtherName  string `json:"otherName"`
		Phone      string `json:"phone"`
		OtherPhone string `json:"otherPhone"`
		UserType   string `json:"userType"`
		OTP        any    `json:"otp"`
	}
)

func AuthAdminResponse(data *ent.Admin) *fiber.Map {
	return successResponse(
		&AuthAdmin{
			ID:       data.ID,
			Username: data.Username,
		},
	)
}
func AuthCustomerResponse(data *ent.Customer) *fiber.Map {
	return successResponse(
		&AuthCustomer{
			ID:        data.ID,
			Username:  data.Username,
			OtherName: data.FirstName,
			LastName:  data.LastName,
			UserType:  "customer",
		},
	)
}
func AuthAgentResponse(data *ent.Agent) *fiber.Map {
	return successResponse(
		&AuthMerchant{
			ID:         data.ID,
			Username:   data.Username,
			OtherName:  data.OtherName,
			LastName:   data.LastName,
			Phone:      data.Phone,
			OtherPhone: *data.OtherPhone,
			UserType:   "agent",
		},
	)
}
func AuthSupplierMerchantResponse(data *AuthMerchant) *fiber.Map {
	return successResponse(
		&AuthMerchant{
			ID:         data.ID,
			Username:   data.Username,
			OtherName:  data.OtherName,
			LastName:   data.LastName,
			Phone:      data.Phone,
			OtherPhone: data.OtherPhone,
			UserType:   "supplier",
			OTP:        data.OTP,
		},
	)
}
func AuthRetailMerchantResponse(data *AuthMerchant) *fiber.Map {
	return successResponse(
		&AuthMerchant{
			ID:         data.ID,
			Username:   data.Username,
			OtherName:  data.OtherName,
			LastName:   data.LastName,
			Phone:      data.Phone,
			OtherPhone: data.OtherPhone,
			UserType:   "retailer",
			OTP:        data.OTP,
		},
	)
}

func AuthErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err,
	}
}
