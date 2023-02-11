package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/admin"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type AdminHandler struct {
	service gateways.AdminService
}

func NewAdminHandler(db *database.Adapter) *AdminHandler {
	repo := admin.NewAdminRepo(db)
	service := admin.NewAdminService(repo)

	return &AdminHandler{
		service: service,
	}
}

func (h *AdminHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.AdminSuccessResponse(result))
	}
}
func (h *AdminHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		result, err := h.service.FetchAll(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.AdminsSuccessResponse(result))
	}
}

func (h *AdminHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.AdminUserRequest

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		result, err := h.service.Create(&request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.AdminSuccessResponse(result))
	}
}

func (h *AdminHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.AdminUserRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.RolePermissionErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(presenters.AdminSuccessResponse(result))
	}

}
func (h *AdminHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.Remove(id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}
