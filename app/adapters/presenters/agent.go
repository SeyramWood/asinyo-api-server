package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type (
	AgentComplianceData struct {
		GhanaCard    []string                    `json:"ghanaCard"`
		PoliceReport string                      `json:"policeReport"`
		Verified     bool                        `json:"verified"`
		Guarantor    *models.AgentGuarantorModel `json:"guarantor"`
		CreatedAt    time.Time                   `json:"created_at"`
		UpdatedAt    time.Time                   `json:"updated_at"`
	}
	Agent struct {
		ID             int                          `json:"id"`
		GhanaCard      string                       `json:"ghanaCard,omitempty"`
		LastName       string                       `json:"lastName,omitempty"`
		OtherName      string                       `json:"otherName,omitempty"`
		Phone          string                       `json:"phone,omitempty"`
		OtherPhone     string                       `json:"otherPhone,omitempty"`
		Address        string                       `json:"address,omitempty"`
		DigitalAddress string                       `json:"digitalAddress,omitempty"`
		Username       string                       `json:"username,omitempty"`
		City           string                       `json:"city,omitempty"`
		District       string                       `json:"district,omitempty"`
		Region         string                       `json:"region,omitempty"`
		Verified       	bool                         `json:"verified"`
		Compliance     *models.AgentComplianceModel `json:"compliance,omitempty"`
		DefaultAccount string                       `json:"defaultAccount,omitempty"`
		BankAccount    *models.AgentBankAccount     `json:"bankAccount,omitempty"`
		MomoAccount    *models.AgentMomoAccount     `json:"momoAccount,omitempty"`
		CreatedAt      time.Time                    `json:"created_at"`
		UpdatedAt      time.Time                    `json:"updated_at"`
	}
	AgentComplianceResponse struct {
		Compliance *models.AgentComplianceModel `json:"compliance"`
	}
)

func AgentSuccessResponse(data *ent.Agent) *fiber.Map {
	return successResponse(
		Agent{
			ID:             data.ID,
			GhanaCard:      data.GhanaCard,
			LastName:       data.LastName,
			OtherName:      data.OtherName,
			Phone:          data.Phone,
			OtherPhone:     data.OtherPhone,
			Address:        data.Address,
			DigitalAddress: data.DigitalAddress,
			Username:       data.Username,
			City:           data.City,
			District:       data.District,
			Region:         data.Region,
			Verified:        data.Verified,
			Compliance:     data.Compliance,
			DefaultAccount: string(data.DefaultAccount),
			BankAccount:    data.BankAccount,
			MomoAccount:    data.MomoAccount,
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
		},
	)
}
func AgentsComplianceSuccessResponse(data *ent.Agent) *fiber.Map {
	return successResponse(
		AgentComplianceResponse{
			Compliance: data.Compliance,
		},
	)
}
func AgentsSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*Agent
	for _, v := range res.Data.([]*ent.Agent) {
		response = append(
			response, &Agent{
				ID:             v.ID,
				GhanaCard:      v.GhanaCard,
				LastName:       v.LastName,
				OtherName:      v.OtherName,
				Phone:          v.Phone,
				OtherPhone:     v.OtherPhone,
				Address:        v.Address,
				DigitalAddress: v.DigitalAddress,
				Username:       v.Username,
				City:           v.City,
				District:       v.District,
				Region:         v.Region,
				// Compliance: &AgentComplianceData{
				// 	Region:       data.Region,
				// 	District:     data.District,
				// 	City:         data.City,
				// 	GhanaCard:    data.Compliance.GhanaCard,
				// 	PoliceReport: data.Compliance.PoliceReport,
				// 	Verified:     data.Verified,
				// 	Guarantor: &models.AgentGuarantorModel{
				// 		GhanaCard:      data.Compliance.Guarantor.GhanaCard,
				// 		LastName:       data.Compliance.Guarantor.LastName,
				// 		OtherName:      data.Compliance.Guarantor.OtherName,
				// 		Phone:          data.Compliance.Guarantor.Phone,
				// 		OtherPhone:     data.Compliance.Guarantor.OtherPhone,
				// 		Address:        data.Compliance.Guarantor.Address,
				// 		DigitalAddress: data.Compliance.Guarantor.DigitalAddress,
				// 		Relation:       data.Compliance.Guarantor.Relation,
				// 		Occupation:     data.Compliance.Guarantor.Occupation,
				// 	},
				// 	CreatedAt: data.CreatedAt,
				// 	UpdatedAt: data.UpdatedAt,
				// },
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}
func AgentsPaginationResponse(res *PaginationResponse) *fiber.Map {
	var response []*Agent
	for _, v := range res.Data.([]*ent.Agent) {
		response = append(
			response, &Agent{
				ID:             v.ID,
				GhanaCard:      v.GhanaCard,
				LastName:       v.LastName,
				OtherName:      v.OtherName,
				Phone:          v.Phone,
				OtherPhone:     v.OtherPhone,
				Address:        v.Address,
				DigitalAddress: v.DigitalAddress,
				City:           v.City,
				District:       v.District,
				Region:         v.Region,
				Verified:       v.Verified,
				CreatedAt:      v.CreatedAt,
				UpdatedAt:      v.UpdatedAt,
			},
		)
	}
	return successResponse(
		&PaginationResponse{
			Count: res.Count,
			Data:  response,
		},
	)
}

func AgentErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
