package api

import (
	"errors"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/agent"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/gofiber/fiber/v2"
)

type AgentHandler struct {
	service gateways.AgentService
}

func NewAgentHandler(db *database.Adapter) *AgentHandler {
	repo := agent.NewAgentRepo(db)
	service := agent.NewAgentService(repo)

	return &AgentHandler{
		service: service,
	}
}

func (h *AgentHandler) FetchByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(err))
		}
		return c.JSON(presenters.AgentSuccessResponse(result))
	}
}
func (h *AgentHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AgentsSuccessResponse(result))
	}
}

func (h *AgentHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.AgentRequest

		err := c.BodyParser(&request)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
		}

		_, err = h.service.Create(&request)

		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(errors.New("error creating agent")))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}

func (h *AgentHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		result, err := h.service.FetchAll()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AgentsSuccessResponse(result))
	}

}
func (h *AgentHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": true,
			"error":  nil,
		})
	}
}
