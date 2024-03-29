package models

type (
	User struct {
		Username string `json:"username" validate:"required|email_phone"`
		Password string `json:"password" validate:"required|min:8"`
	}
	Customer struct {
		CustomerType string `json:"customerType" validate:"required"`
		Username     string `json:"username" validate:"required|email_phone"`
		Password     string `json:"password" validate:"required|min:8"`
	}
	UserMerchant struct {
		MerchantType string `json:"merchantType" validate:"required"`
		Username     string `json:"username" validate:"required|email_phone"`
		Password     string `json:"password" validate:"required|min:8"`
	}
	ChangePassword struct {
		CurrentPassword string `json:"currentPassword" validate:"required|min:8"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword,omitempty" validate:"required|min:8|match:Password"`
	}
	ResetPassword struct {
		NewPassword        string `json:"newPassword" validate:"required|min:8"`
		ConfirmNewPassword string `json:"confirmNewPassword" validate:"required|min:8|match:NewPassword"`
	}
)
