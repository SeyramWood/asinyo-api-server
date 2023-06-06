package models

type (
	AdminUserUpdateRequest struct {
		LastName   string `json:"lastName" validate:"required|string"`
		OtherName  string `json:"otherName" validate:"required|string"`
		Phone      string `json:"phone" validate:"required|phone_no_code"`
		OtherPhone string `json:"otherPhone" validate:"phone_no_code"`
		Username   string `json:"username" validate:"required|email"`
		Roles      []int  `json:"roles" validate:"required"`
	}
	AdminUserRequest struct {
		LastName   string `json:"lastName" validate:"required|string"`
		OtherName  string `json:"otherName" validate:"required|string"`
		Phone      string `json:"phone" validate:"required|phone_no_code"`
		OtherPhone string `json:"otherPhone" validate:"phone_no_code"`
		Username   string `json:"username" validate:"required|email|unique:admins"`
		Roles      []int  `json:"roles" validate:"required"`
	}
	RoleRequest struct {
		Role        string `json:"role" validate:"required|string"`
		Permissions []int  `json:"permissions" validate:"required"`
	}
)
