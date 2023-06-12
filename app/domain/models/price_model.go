package models

type (
	PriceModelRequest struct {
		Name     string `json:"name" validate:"required|string"`
		Initials string `json:"initials" validate:"required|string"`
		Formula  string `json:"formula" validate:"required"`
	}
)
