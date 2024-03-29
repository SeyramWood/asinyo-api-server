package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	SupplierMerchant struct {
		ID             int       `json:"id"`
		GhanaCard      string    `json:"ghanaCard"`
		LastName       string    `json:"lastName"`
		OtherName      string    `json:"firstName"`
		Phone          string    `json:"phone"`
		OtherPhone     *string   `json:"otherPhone"`
		Address        string    `json:"address"`
		DigitalAddress string    `json:"digitalAddress"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
)

func SupplierMerchantSuccessResponse(data *ent.SupplierMerchant) *fiber.Map {
	return successResponse(
		SupplierMerchant{
			ID:         data.ID,
			GhanaCard:  data.GhanaCard,
			LastName:   data.LastName,
			OtherName:  data.OtherName,
			Phone:      data.Phone,
			OtherPhone: data.OtherPhone,
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
		},
	)
}
func SupplierMerchantsSuccessResponse(data []*ent.SupplierMerchant) *fiber.Map {
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

func SupplierMerchantErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
