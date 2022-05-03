package presenters

import (
	"sync"
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

// Book is the presenter object which will be passed in the response by Handler

type (
	User struct {
		ID        int       `json:"id"`
		Username  string    `json:"username"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

// UserSuccessResponse is the singular SuccessResponse that will be passed in the response by
//Handler
func UserSuccessResponse(data *ent.User) *fiber.Map {
	return successResponse(User{
		ID:        data.ID,
		Username:  data.Username,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}

// UsersSuccessResponse is the list SuccessResponse that will be passed in the response by Handler
func UsersSuccessResponse(data []*ent.User) *fiber.Map {
	var response []User
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.User) {
			defer wg.Done()
			response = append(response, User{
				ID:        v.ID,
				Username:  v.Username,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			})
		}(v)
	}
	wg.Wait()
	return successResponse(response)
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func UserErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
