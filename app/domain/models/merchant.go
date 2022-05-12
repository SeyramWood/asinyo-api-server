package models

type (
	MerchantRequestInfo struct {
		MerchantType   string `json:"merchantType" validate:"required|string"`
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"otherName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:retail_merchants"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:retail_merchants"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	MerchantRequestCredentials struct {
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:retail_merchants"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}

	MerchantRequest struct {
		Info        MerchantRequestInfo
		Credentials MerchantRequestCredentials
	}

	Merchant struct {
		MerchantRequestInfo
		MerchantRequestCredentials
	}
)
