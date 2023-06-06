package api

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/role_permission"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type roleAndPermissionHandler struct {
	service gateways.RoleAndPermissionService
}

func NewRoleAndPermissionHandler(db *database.Adapter) *roleAndPermissionHandler {
	return &roleAndPermissionHandler{
		service: role_permission.NewRoleAndPermissionService(role_permission.NewRoleAndPermissionRepo(db)),
	}
}

func (h *roleAndPermissionHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		result, err := h.service.FetchAll(limit, offset)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(presenters.RoleSuccessResponses(result))
	}
}
func (h *roleAndPermissionHandler) FetchPermissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.service.FetchAllPermission()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(presenters.PermissionSuccessResponses(result))
	}
}
func (h *roleAndPermissionHandler) CreateRole() fiber.Handler {

	return func(c *fiber.Ctx) error {
		var request models.RoleRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.RolePermissionErrorResponse(err))
		}
		result, err := h.service.Create(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(presenters.RoleSuccessResponse(result))
	}

}
func (h *roleAndPermissionHandler) UpdateRole() fiber.Handler {

	return func(c *fiber.Ctx) error {
		var request models.RoleRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.RolePermissionErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(presenters.RoleSuccessResponse(result))
	}

}
func (h *roleAndPermissionHandler) DeleteRole() fiber.Handler {

	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		err := h.service.Remove(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"status":  true,
				"message": "Deleted",
			},
		)
	}

}
