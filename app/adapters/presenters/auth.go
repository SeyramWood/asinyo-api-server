package presenters

import (
	"fmt"
	"strings"

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
	AuthSession struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		SessionName string `json:"sessionName"`
		DisplayName string `json:"displayName"`
		UserType    string `json:"userType"`
		OTP         any    `json:"otp,omitempty"`
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
	if c, err := data.Edges.IndividualOrErr(); err == nil {
		return successResponse(
			&AuthSession{
				ID:          data.ID,
				Username:    data.Username,
				SessionName: strings.Split(c.OtherName, " ")[0],
				DisplayName: fmt.Sprintf("%s %s", c.OtherName, c.LastName),
				UserType:    data.Type,
			},
		)
	}
	if c, err := data.Edges.BusinessOrErr(); err == nil {
		return successResponse(
			&AuthSession{
				ID:          data.ID,
				Username:    data.Username,
				SessionName: c.Name,
				DisplayName: c.Name,
				UserType:    data.Type,
			},
		)
	}

	return successResponse(nil)
}
func AuthAgentResponse(data *ent.Agent) *fiber.Map {
	return successResponse(
		&AuthSession{
			ID:          data.ID,
			Username:    data.Username,
			SessionName: strings.Split(data.OtherName, " ")[0],
			DisplayName: fmt.Sprintf("%s %s", data.OtherName, data.LastName),
			UserType:    "agent",
		},
	)
}

func AuthMerchantResponse(data *ent.Merchant) *fiber.Map {
	if s, err := data.Edges.SupplierOrErr(); err == nil {
		return successResponse(
			&AuthSession{
				ID:          data.ID,
				Username:    data.Username,
				SessionName: strings.Split(s.OtherName, " ")[0],
				DisplayName: fmt.Sprintf("%s %s", s.OtherName, s.LastName),
				UserType:    data.Type,
				OTP:         data.Otp,
			},
		)
	}
	if r, err := data.Edges.RetailerOrErr(); err == nil {
		return successResponse(
			&AuthSession{
				ID:          data.ID,
				Username:    data.Username,
				SessionName: strings.Split(r.OtherName, " ")[0],
				DisplayName: fmt.Sprintf("%s %s", r.OtherName, r.LastName),
				UserType:    data.Type,
				OTP:         data.Otp,
			},
		)
	}
	return successResponse(nil)
}

func AuthErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err,
	}
}
