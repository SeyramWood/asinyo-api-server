package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	PriceModel struct {
		ID            int       `json:"id"`
		Name          string    `json:"name,omitempty"`
		Initials      string    `json:"initials,omitempty"`
		Formula       string    `json:"formula,omitempty"`
		AsinyoFormula string    `json:"asinyoFormula,omitempty"`
		CreatedAt     time.Time `json:"created_at,omitempty"`
		UpdatedAt     time.Time `json:"updated_at,omitempty"`
	}
	PriceModelPercentage struct {
		ID         int    `json:"id"`
		Category   string `json:"category,omitempty"`
		Percentage int    `json:"percentage,omitempty"`
	}
)

func PriceModelPercentageSuccessResponse(data *ent.ProductCategoryMinor) *fiber.Map {
	return successResponse(
		&PriceModelPercentage{
			ID:         data.ID,
			Category:   data.Category,
			Percentage: data.Percentage,
		},
	)
}
func PriceModelSuccessResponse(data *ent.PriceModel) *fiber.Map {
	return successResponse(
		&PriceModel{
			ID:            data.ID,
			Name:          data.Name,
			Initials:      data.Initials,
			Formula:       data.Formula,
			AsinyoFormula: data.AsinyoFormula,
			CreatedAt:     data.CreatedAt,
			UpdatedAt:     data.UpdatedAt,
		},
	)
}

func PriceModelsSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*PriceModel
	for _, data := range res.Data.([]*ent.PriceModel) {
		response = append(
			response, &PriceModel{
				ID:            data.ID,
				Name:          data.Name,
				Initials:      data.Initials,
				Formula:       data.Formula,
				AsinyoFormula: data.AsinyoFormula,
				CreatedAt:     data.CreatedAt,
				UpdatedAt:     data.UpdatedAt,
			},
		)

	}
	return successResponse(
		&PaginationResponse{
			Count: res.Count,
			Data:  response,
		},
	)
}

func PriceModelPercentagesSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*PriceModelPercentage
	for _, data := range res.Data.([]*ent.ProductCategoryMinor) {
		response = append(
			response, &PriceModelPercentage{
				ID:         data.ID,
				Category:   data.Category,
				Percentage: data.Percentage,
			},
		)

	}
	return successResponse(
		&PaginationResponse{
			Count: res.Count,
			Data:  response,
		},
	)
}

func PriceModelErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
