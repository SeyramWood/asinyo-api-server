package models

type (
	PickupStation struct {
		Region  string `json:"region" validate:"required|string"`
		City    string `json:"city" validate:"required|string"`
		Name    string `json:"name" validate:"required|string"`
		Address string `json:"address" validate:"required|string"`
	}
)
