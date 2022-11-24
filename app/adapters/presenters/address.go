package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent"
)

type AddressResponse struct {
	ID           int                  `json:"id"`
	LastName     string               `json:"lastName"`
	OtherName    string               `json:"otherName"`
	Phone        string               `json:"phone"`
	OtherPhone   string               `json:"otherPhone"`
	Address      string               `json:"address"`
	Country      string               `json:"country,omitempty"`
	Region       string               `json:"region"`
	City         string               `json:"city"`
	District     string               `json:"district"`
	StreetName   string               `json:"streetName"`
	StreetNumber string               `json:"streetNumber"`
	Coordinate   *services.Coordinate `json:"coordinate"`
	Default      bool                 `json:"default"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
}

func AddressSuccessResponse(data *ent.Address) *fiber.Map {
	if data != nil {
		return successResponse(
			&AddressResponse{
				ID:           data.ID,
				LastName:     data.LastName,
				OtherName:    data.OtherName,
				Phone:        data.Phone,
				OtherPhone:   data.OtherPhone,
				Address:      data.Address,
				Region:       data.Region,
				District:     data.District,
				City:         data.City,
				StreetName:   data.StreetName,
				StreetNumber: data.StreetNumber,
				Coordinate:   data.Coordinate,
				Default:      data.Default,
				CreatedAt:    data.CreatedAt,
				UpdatedAt:    data.UpdatedAt,
			},
		)
	}
	return nil
}

func AddressSuccessResponses(data []*ent.Address) *fiber.Map {
	var response []*AddressResponse
	for _, v := range data {
		response = append(
			response, &AddressResponse{
				ID:           v.ID,
				LastName:     v.LastName,
				OtherName:    v.OtherName,
				Phone:        v.Phone,
				OtherPhone:   v.OtherPhone,
				Address:      v.Address,
				Region:       v.Region,
				District:     v.District,
				City:         v.City,
				StreetName:   v.StreetName,
				StreetNumber: v.StreetNumber,
				Coordinate:   v.Coordinate,
				Default:      v.Default,
				CreatedAt:    v.CreatedAt,
				UpdatedAt:    v.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}

func AddressErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
