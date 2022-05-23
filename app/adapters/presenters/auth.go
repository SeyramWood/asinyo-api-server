package presenters

import (
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthAdmin struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		LastName  string `json:"lastName"`
		OtherName string `json:"otherName"`
	}
	AuthCustomer struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		LastName  string `json:"lastName"`
		OtherName string `json:"otherName"`
		UserType  string `json:"userType"`
	}
	AuthMerchant struct {
		ID        int    `json:"id"`
		Username  string `json:"username"`
		LastName  string `json:"lastName"`
		OtherName string `json:"otherName"`
		UserType  string `json:"userType"`
	}
)

//Handler
func AuthAdminResponse(data *ent.Admin) *fiber.Map {
	return successResponse(&AuthAdmin{
		ID:       data.ID,
		Username: data.Username,
	})
}
func AuthCustomerResponse(data *ent.Customer) *fiber.Map {
	return successResponse(&AuthCustomer{
		ID:        data.ID,
		Username:  data.Username,
		OtherName: data.FirstName,
		LastName:  data.LastName,
		UserType:  "customer",
	})
}
func AuthAgentResponse(data *ent.Agent) *fiber.Map {
	return successResponse(&AuthMerchant{
		ID:        data.ID,
		Username:  data.Username,
		OtherName: data.OtherName,
		LastName:  data.LastName,
		UserType:  "agent",
	})
}
func AuthSupplierMerchantResponse(data *AuthMerchant) *fiber.Map {
	return successResponse(&AuthMerchant{
		ID:        data.ID,
		Username:  data.Username,
		OtherName: data.OtherName,
		LastName:  data.LastName,
		UserType:  "supplier",
	})
}
func AuthRetailMerchantResponse(data *AuthMerchant) *fiber.Map {
	return successResponse(&AuthMerchant{
		ID:        data.ID,
		Username:  data.Username,
		OtherName: data.OtherName,
		LastName:  data.LastName,
		UserType:  "retailer",
	})
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func AuthErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err,
	}
}
