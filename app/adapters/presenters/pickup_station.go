package presenters

import (
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
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

	return successResponse(&PickupStation{
		ID:        data.ID,
		Address:   data.Address,
		Region:    data.Region,
		City:      data.City,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}

func PickupSuccessResponses(data []*ent.PickupStation) *fiber.Map {
	var response []*PickupStation
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.PickupStation) {
			defer wg.Done()
			response = append(response, &PickupStation{
				ID:        v.ID,
				Address:   v.Address,
				Region:    v.Region,
				City:      v.City,
				Name:      v.Name,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}(v)
	}
	wg.Wait()

	return successResponse(response)
}

func PickupErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
