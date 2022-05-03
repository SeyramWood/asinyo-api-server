package models

type (
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Auth struct {
		Username string `json:"username" validate:"required|min:3"`
		Password string `json:"password" validate:"required|min:6|max:10"`
	}
)
