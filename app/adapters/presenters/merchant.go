package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	Merchant struct {
		ID         int       `json:"id"`
		GhanaCard  string    `json:"ghanaCard"`
		LastName   string    `json:"lastName"`
		OtherName  string    `json:"otherName"`
		Phone      string    `json:"phone"`
		OtherPhone string    `json:"otherPhone"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
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
func MerchantsSuccessResponse(data []*ent.Merchant) *fiber.Map {
	// var response []Customer
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
	return successResponse(nil)
}

func MerchantErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
