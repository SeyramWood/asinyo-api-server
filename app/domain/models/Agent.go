package models

type (
	Agent struct {
		GhanaCard       string `json:"ghanaCard" validate:"required|id_card"`
		LastName        string `json:"lastName" validate:"required|string"`
		OtherName       string `json:"firstName" validate:"required|string"`
		Phone           string `json:"phone" validate:"required|string|unique:agents"`
		OtherPhone      string `json:"otherPhone" validate:"required|string|unique:agents"`
		Address         string `json:"address" validate:"required|string"`
		DigitalAddress  string `json:"digitalAddress" validate:"required|string|digital_address"`
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"email" validate:"required|email_phone|unique:agents"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}
	AgentInfo struct {
		GhanaCard      string `json:"ghanaCard" validate:"required|id_card"`
		LastName       string `json:"lastName" validate:"required|string"`
		OtherName      string `json:"firstName" validate:"required|string"`
		Phone          string `json:"phone" validate:"required|string|unique:agents"`
		OtherPhone     string `json:"otherPhone" validate:"string|unique:agents"`
		Address        string `json:"address" validate:"required|string"`
		DigitalAddress string `json:"digitalAddress" validate:"required|string|digital_address"`
	}
	AgentCredentials struct {
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"email" validate:"required|email_phone|unique:agents"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}
)
