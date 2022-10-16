package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	Admin struct {
		ID        int       `json:"id"`
		Username  string    `json:"firstName"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func AdminSuccessResponse(data *ent.Admin) *fiber.Map {
	return successResponse(
		Admin{
			ID:        data.ID,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	)
}
func AdminsSuccessResponse(data []*ent.Admin) *fiber.Map {
	var response []Customer
	// wg := sync.WaitGroup{}
	// for _, v := range data {
	// 	wg.Add(1)
	// 	go func(v *ent.Customer) {
	// 		defer wg.Done()
	// 		response = append(response, User{
	// 			ID:        v.ID,
	// 			Username:  v.Username,
	// 			CreatedAt: v.CreatedAt,
	// 			UpdatedAt: v.UpdatedAt,
	// 		})
	// 	}(v)
	// }
	// wg.Wait()
	return successResponse(response)
}

func AdminErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
