package models

type (
	RetailMerchant struct {
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:retail_merchants"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
)
