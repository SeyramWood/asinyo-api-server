package presenters

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"

	"github.com/SeyramWood/ent"
)

type (
	AuthSession struct {
		ID          int      `json:"id"`
		Username    string   `json:"username"`
		SessionName string   `json:"sessionName"`
		DisplayName string   `json:"displayName,omitempty"`
		UserType    string   `json:"userType,omitempty"`
		OTP         any      `json:"otp,omitempty"`
		Storefront  int      `json:"storefront,omitempty"`
		Permissions []string `json:"permissions,omitempty"`
	}
)

func AuthAdminResponse(data *ent.Admin) *fiber.Map {
	return successResponse(
		&AuthSession{
			ID:          data.ID,
			Username:    data.Username,
			SessionName: strings.Split(data.OtherName, " ")[0],
			Permissions: func() []string {
				roles, err := data.Edges.RolesOrErr()
				if err != nil {
					return nil
				}
				var permissions []string
				for _, role := range roles {
					perms, err := role.Edges.PermissionsOrErr()
					if err != nil {
						return nil
					}
					for _, perm := range perms {
						if lo.Contains(permissions, perm.Slug) {
							continue
						}
						permissions = append(permissions, perm.Slug)
					}
				}
				return permissions
			}(),
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
				Storefront:  data.Edges.Store.ID,
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
				Storefront:  data.Edges.Store.ID,
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
