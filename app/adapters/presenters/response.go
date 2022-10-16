package presenters

import "github.com/gofiber/fiber/v2"

func errorResponse(e error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  e.Error(),
	}
}
func successResponse(data any) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
	}
}
func EmptySuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   nil,
	}
}
func EmptySuccessResponseMany() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   []string{},
	}
}
