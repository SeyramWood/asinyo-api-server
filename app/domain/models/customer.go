package models

type (
	BusinessCustomer struct {
		BusinessName    string `json:"businessName" validate:"required"`
		Phone           string `json:"phone" validate:"required|string"`
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:customers"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|min:8|match:Password"`
	}
	IndividualCustomer struct {
		LastName        string `json:"lastName" validate:"required|string"`
		OtherName       string `json:"otherName" validate:"required|string"`
		Phone           string `json:"phone" validate:"required|string"`
		Terms           bool   `json:"terms" validate:"required|bool"`
		Username        string `json:"username" validate:"required|email_phone|unique:customers"`
		Password        string `json:"password" validate:"required|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|min:8|match:Password"`
	}
	BusinessCustomerContact struct {
		Name     string `json:"businessName" validate:"required"`
		Position string `json:"position" validate:"required|string"`
		Phone    string `json:"phone" validate:"required|string|phone"`
		Email    string `json:"username" validate:"required|email"`
	}
)
