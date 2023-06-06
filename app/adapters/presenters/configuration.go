package presenters

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	LogisticConfiguration struct {
		Logistics []string `json:"logistics"`
		Current   string   `json:"current"`
	}
	LogisticConfigurations struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Data any    `json:"data"`
	}
)

func ConfigurationsResponse(data []*ent.Configuration) *fiber.Map {
	var response []*LogisticConfigurations
	for _, v := range data {
		response = append(response, &LogisticConfigurations{
			ID:   v.ID,
			Name: v.Name,
			Data: v.Data.Data,
		})
	}
	return successResponse(response)
}

func ConfigurationResponse(data *ent.Configuration) *fiber.Map {
	return successResponse(&LogisticConfigurations{
		ID:   data.ID,
		Name: data.Name,
		Data: data.Data.Data,
	})
}

func ConfigurationErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
