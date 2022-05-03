package handlers

import (
	"errors"

	"github.com/SeyramWood/app/adapters"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/user"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service gateways.UserService
}

func NewUserHandler(db *database.Adapter) adapters.Handler {
	repo := user.NewUserRepo(db)
	service := user.NewUserService(repo)

	return &userHandler{
		service: service,
	}
}

func (h *userHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.UserErrorResponse(err))
		}
		return c.JSON(presenters.UserSuccessResponse(result))
	}
}
func (h *userHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.UserErrorResponse(err))
		}
		return c.JSON(presenters.UsersSuccessResponse(result))
	}
}

func (h *userHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.User

		err := c.BodyParser(&request)

		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(presenters.UserErrorResponse(err))
		}

		if request.Username == "" || request.Password == "" {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.UserErrorResponse(errors.New(
				"please specify title and sbn")))
		}

		result, err := h.service.Create(&request)

		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(presenters.UserErrorResponse(err))
		}
		return c.JSON(presenters.UserSuccessResponse(result))
	}
}

func (h *userHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.UserErrorResponse(err))
		}
		return c.JSON(presenters.UsersSuccessResponse(result))
	}
}
func (h *userHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.UserErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
