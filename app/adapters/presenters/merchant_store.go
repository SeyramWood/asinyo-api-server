package presenters

import (
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

type (
	storeDetails struct {
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
	store struct {
		ID           int       `json:"id"`
		BusinessName string    `json:"businessName"`
		About        string    `json:"about"`
		Logo         string    `json:"logo"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	MerchantInfo struct {
		ID int `json:"id"`
	}
	AllMerchantStore struct {
		*store
	}
)

func MerchantStoreSuccessResponse(data *ent.MerchantStore) *fiber.Map {
	if data == nil {
		return successResponse(nil)
	}

	return successResponse(&storeDetails{
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
		MerchantInfo: &MerchantInfo{
			ID: data.Edges.Merchant.ID,
		},
	})
}

func MerchantStorefrontSuccessResponse(data *ent.MerchantStore) *fiber.Map {
	return successResponse(&storeDetails{
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
func MerchantStorefrontsSuccessResponse(data []*ent.MerchantStore) *fiber.Map {
	var response []AllMerchantStore
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.MerchantStore) {
			defer wg.Done()
			response = append(response, AllMerchantStore{
				&store{
					ID:           v.ID,
					BusinessName: v.Name,
					About:        v.About,
					Logo:         v.Logo,
					CreatedAt:    v.CreatedAt,
					UpdatedAt:    v.UpdatedAt,
				},
			})
		}(v)
	}
	wg.Wait()
	return successResponse(response)
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
