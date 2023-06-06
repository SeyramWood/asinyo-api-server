package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/admin"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type AdminHandler struct {
	service gateways.AdminService
}

func NewAdminHandler(db *database.Adapter, noti notification.NotificationService) *AdminHandler {
	repo := admin.NewAdminRepo(db)
	service := admin.NewAdminService(repo, noti)

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

func (h *AdminHandler) FetchCounts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		span := c.Query("span", "30")
		result, err := h.service.FetchCounts(span)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.DashboardCountSuccessResponse(result))
	}
}
func (h *AdminHandler) FetchAccountManagers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		perm := c.Query("permission", "view-clients")
		result, err := h.service.FetchAccountManagers(perm)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.AccountManagersSuccessResponse(result))
	}
}

func (h *AdminHandler) FetchConfigurations() fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := h.service.FetchConfigurations()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.ConfigurationsResponse(result))
	}
}

func (h *AdminHandler) FetchConfigurationByIdOrName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var slug any
		if s, err := c.ParamsInt("slug"); err != nil {
			slug = c.Params("slug")
		} else {
			slug = s
		}
		result, err := h.service.FetchConfigurationByIdOrName(slug)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.ConfigurationResponse(result))
	}
}
func (h *AdminHandler) FetchMyClients() fiber.Handler {
	return func(c *fiber.Ctx) error {

		manager, _ := strconv.Atoi(c.Query("manager", "0"))
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		results, err := h.service.FetchMyClients(manager, limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.MyClientsSuccessResponse(results))
	}
}
func (h *AdminHandler) FetchMyClientsPurchaseRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {

		manager, _ := strconv.Atoi(c.Query("manager", "0"))
		results, err := h.service.FetchMyClientsPurchaseRequest(manager)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.MyClientsPurchaseRequestSuccessResponse(results))
	}
}
func (h *AdminHandler) FetchMyClientOrders() fiber.Handler {
	return func(c *fiber.Ctx) error {

		manager, _ := strconv.Atoi(c.Query("manager", "0"))
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))

		results, err := h.service.FetchMyClientOrders(manager, limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.PaginateOrdersSuccessResponse(results))
	}
}
func (h *AdminHandler) FetchProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		major := c.Query("major")
		minor := c.Query("minor")
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		results, err := h.service.FetchProducts(major, minor, limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AdminProductsResponse(results))
	}
}
func (h *AdminHandler) FetchAdminProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		results, err := h.service.FetchAdminProducts(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AdminProductsResponse(results))
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
func (h *AdminHandler) FetchOrders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		results, err := h.service.FetchAllOrders(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.PaginateOrdersSuccessResponse(results))
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
func (h *AdminHandler) OnboardNewCustomer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.BusinessCustomerOnboardRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		manager, _ := strconv.Atoi(c.Query("manager"))
		result, err := h.service.OnboardNewCustomer(manager, &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.JSON(presenters.MyClientSuccessResponse(result))
	}
}

func (h *AdminHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.AdminUserRequest
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AdminErrorResponse(err))
		}
		id, _ := c.ParamsInt("id")
		result, err := h.service.Update(id, &request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.JSON(presenters.AdminSuccessResponse(result))
	}
}

func (h *AdminHandler) AssignAccountManager() fiber.Handler {
	return func(c *fiber.Ctx) error {
		manager, _ := strconv.Atoi(c.Query("manager"))
		customer, _ := strconv.Atoi(c.Query("customer"))
		_, err := h.service.AssignAccountManager(manager, customer)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.RolePermissionErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
	}
}

func (h *AdminHandler) UpdateCurrentConfiguration() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		configurationType := c.Params("configuration")
		configValue := c.Query("configValue")
		_, err := h.service.UpdateCurrentConfiguration(id, configurationType, configValue)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AdminErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(
			&fiber.Map{
				"status": true,
				"error":  nil,
			},
		)
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
func (h *AdminHandler) DeleteOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		if err := h.service.RemoveOrder(id); err != nil {
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
