package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	PickupStation struct {
		ID        int       `json:"id"`
		Region    string    `json:"region"`
		City      string    `json:"city"`
		Name      string    `json:"name"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func PickupSuccessResponse(data *ent.PickupStation) *fiber.Map {

	return successResponse(
		&PickupStation{
			ID:        data.ID,
			Address:   data.Address,
			Region:    data.Region,
			City:      data.City,
			Name:      data.Name,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	)
}

func PickupSuccessResponses(data []*ent.PickupStation) *fiber.Map {
	var response []*PickupStation
	for _, v := range data {
		response = append(
			response, &PickupStation{
				ID:        v.ID,
				Address:   v.Address,
				Region:    v.Region,
				City:      v.City,
				Name:      v.Name,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
		)
	}

	return successResponse(response)
}

func PickupErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
