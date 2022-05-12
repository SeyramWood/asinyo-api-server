package presenters

import (
	"sync"
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	Customer struct {
		ID        int       `json:"id"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Username  string    `json:"username"`
		Phone     string    `json:"phone"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func CustomerSuccessResponse(data *ent.Customer) *fiber.Map {
	return successResponse(Customer{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Phone:     data.Phone,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func CustomersSuccessResponse(data []*ent.Customer) *fiber.Map {
	var response []Customer
	wg := sync.WaitGroup{}

	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Customer) {
			defer wg.Done()
			response = append(response, Customer{
				ID:        v.ID,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Username:  v.Username,
				Phone:     v.Phone,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}(v)
	}

	wg.Wait()
	return successResponse(response)
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func CustomerErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
