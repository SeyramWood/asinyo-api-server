package models

type (
	Customer struct {
		FirstName       string `json:"firstName" validate:"required|string"`
		LastName        string `json:"lastName" validate:"required|string"`
		Phone           string `json:"phone" validate:"required|string"`
		Terms           bool   `json:"terms" validate:"required|bool"`
		Email           string `json:"email" validate:"required|email|unique:customers"`
		Password        string `json:"password" validate:"required|string|min:8"`
		ConfirmPassword string `json:"confirmPassword" validate:"required|string|min:8|match:password"`
	}
)
