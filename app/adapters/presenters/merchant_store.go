package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type (
	storeDetails struct {
		ID             int                          `json:"id"`
		BusinessName   string                       `json:"businessName"`
		BusinessSlogan string                       `json:"businessSlogan"`
		About          string                       `json:"about"`
		Description    string                       `json:"description"`
		Logo           string                       `json:"logo"`
		Images         []string                     `json:"images"`
		DefaultAccount string                       `json:"defaultAccount"`
		BankAccount    *models.MerchantBankAccount  `json:"bankAccount"`
		MomoAccount    *models.MerchantMomoAccount  `json:"momoAccount"`
		Address        *models.MerchantStoreAddress `json:"address"`
		PermitAgent    bool                         `json:"permitAgent"`
		Agent          *agentMerchantTypeDetails    `json:"agent"`
		CreatedAt      time.Time                    `json:"created_at"`
		UpdatedAt      time.Time                    `json:"updated_at"`
		MerchantInfo   *MerchantInfo                `json:"merchant"`
	}
	agentMerchantTypeDetails struct {
		ID             int    `json:"id"`
		LastName       string `json:"lastName"`
		OtherName      string `json:"otherName"`
		Phone          string `json:"phone"`
		OtherPhone     string `json:"otherPhone"`
		Address        string `json:"address"`
		DigitalAddress string `json:"digitalAddress"`
	}
	store struct {
		ID           int       `json:"id"`
		BusinessName string    `json:"businessName"`
		About        string    `json:"about"`
		Logo         string    `json:"logo"`
		PermitAgent  bool      `json:"permitAgent"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	MerchantInfo struct {
		ID int `json:"id"`
	}
	AllMerchantStore struct {
		*store
	}
	AgentMerchantDetail struct {
		ID      int                       `json:"id"`
		Type    string                    `json:"type"`
		Profile *agentMerchantTypeDetails `json:"profile"`
	}
	AgentAllMerchantStore struct {
		ID           int                  `json:"id"`
		BusinessName string               `json:"businessName"`
		About        string               `json:"about"`
		Logo         string               `json:"logo"`
		PermitAgent  bool                 `json:"permitAgent"`
		CreatedAt    time.Time            `json:"created_at"`
		UpdatedAt    time.Time            `json:"updated_at"`
		Merchant     *AgentMerchantDetail `json:"merchant"`
	}
)

func MerchantStoreSuccessResponse(data *ent.MerchantStore) *fiber.Map {
	if data == nil {
		return successResponse(nil)
	}

	return successResponse(
		&storeDetails{
			ID:             data.ID,
			BusinessName:   data.Name,
			BusinessSlogan: data.Slogan,
			About:          data.About,
			Description:    data.Description,
			Logo:           data.Logo,
			Images:         data.Images,
			DefaultAccount: string(data.DefaultAccount),
			BankAccount:    data.BankAccount,
			MomoAccount:    data.MomoAccount,
			Address:        data.Address,
			PermitAgent:    data.PermitAgent,
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
			MerchantInfo: &MerchantInfo{
				ID: data.Edges.Merchant.ID,
			},
			Agent: func() *agentMerchantTypeDetails {
				if a, err := data.Edges.AgentOrErr(); err == nil {
					return &agentMerchantTypeDetails{
						ID:             a.ID,
						LastName:       a.LastName,
						OtherName:      a.OtherName,
						Phone:          a.Phone,
						OtherPhone:     *a.OtherPhone,
						Address:        a.Address,
						DigitalAddress: a.DigitalAddress,
					}
				}
				return nil
			}(),
		},
	)
}

func MerchantStorefrontSuccessResponse(data *ent.MerchantStore) *fiber.Map {
	return successResponse(
		&storeDetails{
			ID:             data.ID,
			BusinessName:   data.Name,
			BusinessSlogan: data.Slogan,
			About:          data.About,
			Description:    data.Description,
			Logo:           data.Logo,
			Images:         data.Images,
			DefaultAccount: string(data.DefaultAccount),
			BankAccount:    data.BankAccount,
			MomoAccount:    data.MomoAccount,
			Address:        data.Address,
			PermitAgent:    data.PermitAgent,
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
		},
	)
}

func MerchantStoreAgentSuccessResponse(data *ent.Agent) *fiber.Map {
	return successResponse(
		&agentMerchantTypeDetails{
			ID:             data.ID,
			LastName:       data.LastName,
			OtherName:      data.OtherName,
			Phone:          data.Phone,
			OtherPhone:     *data.OtherPhone,
			Address:        data.Address,
			DigitalAddress: data.DigitalAddress,
		},
	)
}

func MerchantStorefrontsSuccessResponse(data []*ent.MerchantStore) *fiber.Map {
	var response []*AllMerchantStore
	for _, v := range data {
		response = append(
			response, &AllMerchantStore{
				&store{
					ID:           v.ID,
					BusinessName: v.Name,
					About:        v.About,
					Logo:         v.Logo,
					PermitAgent:  v.PermitAgent,
					CreatedAt:    v.CreatedAt,
					UpdatedAt:    v.UpdatedAt,
				},
			},
		)
	}
	return successResponse(response)
}

func AgentMerchantStorefrontsSuccessResponse(data []*ent.MerchantStore) *fiber.Map {
	var response []*AgentAllMerchantStore
	for _, v := range data {
		response = append(
			response, &AgentAllMerchantStore{
				ID:           v.ID,
				BusinessName: v.Name,
				About:        v.About,
				Logo:         v.Logo,
				PermitAgent:  v.PermitAgent,
				CreatedAt:    v.CreatedAt,
				UpdatedAt:    v.UpdatedAt,
				Merchant: &AgentMerchantDetail{
					ID:   v.Edges.Merchant.ID,
					Type: v.Edges.Merchant.Type,
					Profile: func() *agentMerchantTypeDetails {
						if s, err := v.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
							return &agentMerchantTypeDetails{
								ID:         s.ID,
								LastName:   s.LastName,
								OtherName:  s.OtherName,
								Phone:      s.Phone,
								OtherPhone: *s.OtherPhone,
							}
						}
						if s, err := v.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
							return &agentMerchantTypeDetails{
								ID:         s.ID,
								LastName:   s.LastName,
								OtherName:  s.OtherName,
								Phone:      s.Phone,
								OtherPhone: *s.OtherPhone,
							}
						}
						return nil
					}(),
				},
			},
		)
	}

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
