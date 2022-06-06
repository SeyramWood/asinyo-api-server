package presenters

import (
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
	"time"
)

type (
	store struct {
		ID               int                         `json:"id"`
		BusinessName     string                      `json:"businessName"`
		About            string                      `json:"about"`
		DescriptionTitle string                      `json:"descTitle"`
		Description      string                      `json:"description"`
		Logo             string                      `json:"logo"`
		Images           []string                    `json:"images"`
		DefaultAccount   string                      `json:"defaultAccount"`
		BankAccount      *models.MerchantBankAccount `json:"bankAccount"`
		MomoAccount      *models.MerchantMomoAccount `json:"momoAccount"`
		CreatedAt        time.Time                   `json:"created_at"`
		UpdatedAt        time.Time                   `json:"updated_at"`
		MerchantInfo     *MerchantInfo               `json:"merchant"`
	}
	MerchantInfo struct {
		ID int `json:"merchantId"`
	}
)

func MerchantStoreSuccessResponse(data *ent.Merchant) *fiber.Map {
	if data.Edges.Store == nil {
		return successResponse(nil)
	}
	return successResponse(&store{
		ID:               data.Edges.Store.ID,
		BusinessName:     data.Edges.Store.Name,
		About:            data.Edges.Store.About,
		DescriptionTitle: data.Edges.Store.DescTitle,
		Description:      data.Edges.Store.Description,
		Logo:             data.Edges.Store.Logo,
		Images:           data.Edges.Store.Images,
		DefaultAccount:   string(data.Edges.Store.DefaultAccount),
		BankAccount:      data.Edges.Store.BankAccount,
		MomoAccount:      data.Edges.Store.MomoAccount,
		CreatedAt:        data.Edges.Store.CreatedAt,
		UpdatedAt:        data.Edges.Store.UpdatedAt,
		MerchantInfo: &MerchantInfo{
			ID: data.ID,
		},
	})
}

func MerchantStorefrontSuccessResponse(data *ent.MerchantStore) *fiber.Map {
	return successResponse(&store{
		ID:               data.ID,
		BusinessName:     data.Name,
		About:            data.About,
		DescriptionTitle: data.DescTitle,
		Description:      data.Description,
		Logo:             data.Logo,
		Images:           data.Images,
		DefaultAccount:   string(data.DefaultAccount),
		BankAccount:      data.BankAccount,
		MomoAccount:      data.MomoAccount,
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	})
}
func MerchantStoresSuccessResponse(data []*ent.MerchantStore) *fiber.Map {
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

func MerchantStoreErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
