package models

type (
	User struct {
		Username string `json:"username" validate:"required|email_phone"`
		Password string `json:"password" validate:"required|min:8"`
	}
	UserMerchant struct {
		MerchantType string `json:"merchantType" validate:"required"`
		Username     string `json:"username" validate:"required|email_phone"`
		Password     string `json:"password" validate:"required|min:8"`
	}
)
