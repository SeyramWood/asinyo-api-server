package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
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
		Username       string    `json:"username"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}
	AgentComplianceData struct {
		ID           int                         `json:"id"`
		Region       string                      `json:"region"`
		District     string                      `json:"district"`
		City         string                      `json:"city"`
		GhanaCard    []string                    `json:"ghanaCard"`
		PoliceReport string                      `json:"policeReport"`
		Verified     bool                        `json:"verified"`
		Guarantor    *models.AgentGuarantorModel `json:"guarantor"`
		CreatedAt    time.Time                   `json:"created_at"`
		UpdatedAt    time.Time                   `json:"updated_at"`
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
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
		},
	)
}
func AgentsComplianceSuccessResponse(data *ent.Agent) *fiber.Map {
	return successResponse(
		AgentComplianceData{
			ID:           data.ID,
			Region:       *data.Region,
			District:     *data.District,
			City:         *data.City,
			GhanaCard:    data.Compliance.GhanaCard,
			PoliceReport: data.Compliance.PoliceReport,
			Verified:     data.Verified,
			Guarantor: &models.AgentGuarantorModel{
				GhanaCard:      data.Compliance.Guarantor.GhanaCard,
				LastName:       data.Compliance.Guarantor.LastName,
				OtherName:      data.Compliance.Guarantor.OtherName,
				Phone:          data.Compliance.Guarantor.Phone,
				OtherPhone:     data.Compliance.Guarantor.OtherPhone,
				Address:        data.Compliance.Guarantor.Address,
				DigitalAddress: data.Compliance.Guarantor.DigitalAddress,
				Relation:       data.Compliance.Guarantor.Relation,
				Occupation:     data.Compliance.Guarantor.Occupation,
			},
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	)
}
func AgentsSuccessResponse(data []*ent.Agent) *fiber.Map {
	var response []*Agent
	for _, v := range data {
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
				CreatedAt:      v.CreatedAt,
				UpdatedAt:      v.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}

func AgentErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
