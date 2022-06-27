package models

type Address struct {
	LastName     string `json:"lastName" validate:"required|string"`
	OtherName    string `json:"otherName" validate:"required|string"`
	Phone        string `json:"phone" validate:"required|string"`
	OtherPhone   string `json:"otherPhone" validate:"string"`
	Address      string `json:"address" validate:"required|string"`
	OtherAddress string `json:"otherAddress" validate:"string"`
	Region       string `json:"region" validate:"required|string"`
	City         string `json:"city" validate:"required|string"`
	Default      bool   `json:"default"`
}
