package models

type (
	Address struct {
		LastName     string `json:"lastName" validate:"required|string"`
		OtherName    string `json:"otherName" validate:"required|string"`
		Phone        string `json:"phone" validate:"required|string"`
		OtherPhone   string `json:"otherPhone" validate:"string"`
		Address      string `json:"postalAddress" validate:"required"`
		Country      string `json:"country,omitempty" validate:"required|string"`
		Region       string `json:"region" validate:"required|string"`
		City         string `json:"city" validate:"required|string"`
		District     string `json:"district" validate:"required|string"`
		StreetName   string `json:"streetName" validate:"required|string"`
		StreetNumber string `json:"streetNumber" validate:"string"`
		Default      bool   `json:"default"`
	}
)
