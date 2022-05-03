package presenters

import (
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

// Book is the presenter object which will be passed in the response by Handler

type (
	AuthUser struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
)

// UserSuccessResponse is the singular SuccessResponse that will be passed in the response by
//Handler
func AuthSuccessResponse(data *ent.User) *fiber.Map {
	return successResponse(AuthUser{
		ID:       data.ID,
		Username: data.Username,
	})
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func AuthErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err,
	}
}
