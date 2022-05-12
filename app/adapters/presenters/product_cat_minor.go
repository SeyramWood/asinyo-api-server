package presenters

import (
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	ProductCatMinor struct {
		ID        int       `json:"id"`
		Major     int       `json:"major"`
		Category  string    `json:"category"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func ProductCatMinorSuccessResponse(data *ent.ProductCategoryMinor) *fiber.Map {
	return successResponse(Merchant{
		ID: data.ID,
		// GhanaCard:      data.GhanaCard,
		// LastName:       data.LastName,
		// OtherName:      data.OtherName,
		// Phone:          data.Phone,
		// OtherPhone:     data.OtherPhone,
		// Address:        data.Address,
		// DigitalAddress: data.DigitalAddress,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func ProductCatMinorsSuccessResponse(data []*ent.ProductCategoryMinor) *fiber.Map {
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
func ProductCatMinorErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
