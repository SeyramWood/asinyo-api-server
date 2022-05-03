package models

type (
	Admin struct {
		Username string `json:"firstName" validate:"required|unique:admins"`
		Password string `json:"password" validate:"required|string|min:8"`
	}
)
