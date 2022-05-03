package models

type (
	RetailMerchant struct {
		GhanaCard       string `json:"ghanaCard" validate:"required|id_card"`
		LastName        string `json:"lastName" validate:"required|string"`
		OtherName       string `json:"firstName" validate:"required|string"`
		Phone           string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone      string `json:"otherPhone" validate:"required|string|unique:retail_merchants"`
		Address         string `json:"address" validate:"required|string"`
		DigitalAddress  string `json:"digitalAddress" validate:"required|string|digital_address"`
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"email" validate:"required|email_phone|unique:retail_merchants"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}
	RetailMerchantInfo struct {
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"firstName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"required|string|unique:retail_merchants"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	RetailMerchantCredentials struct {
		MerchantType    bool   `json:"merchantType" validate:"required|string"`
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"email" validate:"required|email_phone|unique:retail_merchants"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}
)
