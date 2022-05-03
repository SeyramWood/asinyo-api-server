package presenters

import (
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	Agent struct {
		ID             int       `json:"id"`
		GhanaCard      string    `json:"ghanaCard"`
		LastName       string    `json:"lastName"`
		OtherName      string    `json:"firstName"`
		Phone          string    `json:"phone"`
		OtherPhone     *string   `json:"otherPhone"`
		Address        string    `json:"address"`
		DigitalAddress string    `json:"digitalAddress"`
		Username       string    `json:"terms"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

func AgentSuccessResponse(data *ent.Agent) *fiber.Map {
	return successResponse(Agent{
		ID:             data.ID,
		GhanaCard:      data.GhanaCard,
		LastName:       data.LastName,
		OtherName:      data.OtherName,
		Phone:          data.Phone,
		OtherPhone:     data.OtherPhone,
		Address:        data.Address,
		DigitalAddress: data.DigitalAddress,
		Username:       data.Username,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
	})
}
func AgentsSuccessResponse(data []*ent.Agent) *fiber.Map {
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

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func AgentErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
