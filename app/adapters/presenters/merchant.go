package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	Merchant struct {
		ID           int       `json:"id"`
		Username     string    `json:"username,omitempty"`
		GhanaCard    string    `json:"ghanaCard"`
		LastName     string    `json:"lastName"`
		OtherName    string    `json:"otherName"`
		Phone        string    `json:"phone"`
		OtherPhone   string    `json:"otherPhone"`
		MerchantType string    `json:"merchantType,omitempty"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)

func MerchantSuccessResponse(data *ent.Merchant) *fiber.Map {
	if m, err := data.Edges.SupplierOrErr(); err == nil {
		return successResponse(
			&Merchant{
				ID:         data.ID,
				GhanaCard:  m.GhanaCard,
				LastName:   m.LastName,
				OtherName:  m.OtherName,
				Phone:      m.Phone,
				OtherPhone: *m.OtherPhone,
				CreatedAt:  data.CreatedAt,
				UpdatedAt:  data.UpdatedAt,
			},
		)
	}
	if m, err := data.Edges.RetailerOrErr(); err == nil {
		return successResponse(
			&Merchant{
				ID:         data.ID,
				GhanaCard:  m.GhanaCard,
				LastName:   m.LastName,
				OtherName:  m.OtherName,
				Phone:      m.Phone,
				OtherPhone: *m.OtherPhone,
				CreatedAt:  data.CreatedAt,
				UpdatedAt:  data.UpdatedAt,
			},
		)
	}
	return successResponse(nil)
}

func MerchantsSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*Merchant
	for _, data := range res.Data.([]*ent.Merchant) {
		if profile, err := data.Edges.SupplierOrErr(); err == nil {
			response = append(
				response, &Merchant{
					ID:           data.ID,
					Username:     data.Username,
					GhanaCard:    profile.GhanaCard,
					LastName:     profile.LastName,
					OtherName:    profile.OtherName,
					Phone:        profile.Phone,
					OtherPhone:   *profile.OtherPhone,
					MerchantType: "supplier",
					CreatedAt:    data.CreatedAt,
					UpdatedAt:    data.UpdatedAt,
				},
			)
		}
		if profile, err := data.Edges.RetailerOrErr(); err == nil {
			response = append(response, &Merchant{
				ID:           data.ID,
				Username:     data.Username,
				GhanaCard:    profile.GhanaCard,
				LastName:     profile.LastName,
				OtherName:    profile.OtherName,
				Phone:        profile.Phone,
				OtherPhone:   *profile.OtherPhone,
				MerchantType: "retailer",
				CreatedAt:    data.CreatedAt,
				UpdatedAt:    data.UpdatedAt,
			})
		}
	}
	return successResponse(
		&PaginationResponse{
			Count: res.Count,
			Data:  response,
		},
	)
}

func MerchantErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
