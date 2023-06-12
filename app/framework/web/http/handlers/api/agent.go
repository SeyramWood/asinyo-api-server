package api

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application/agent"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
)

type AgentHandler struct {
	service    gateways.AgentService
	storageSrv gateways.StorageService
}

func NewAgentHandler(db *database.Adapter, storageSrv gateways.StorageService) *AgentHandler {
	repo := agent.NewAgentRepo(db)
	service := agent.NewAgentService(repo)

	return &AgentHandler{
		service:    service,
		storageSrv: storageSrv,
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
func (h *AgentHandler) FetchComplianceByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("agent")
		result, err := h.service.Fetch(id)

		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(err))
		}
		return c.JSON(presenters.AgentsComplianceSuccessResponse(result))
	}
}
func (h *AgentHandler) Fetch() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		results, err := h.service.FetchAll(limit, offset)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AgentsSuccessResponse(results))
	}
}
func (h *AgentHandler) FetchWithPaginate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		results, err := h.service.FetchAll(limit, offset)
		if err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AgentsPaginationResponse(results))
	}
}

func (h *AgentHandler) FetchAllMerchant() fiber.Handler {
	return func(c *fiber.Ctx) error {
		agentId, _ := c.ParamsInt("agent")
		results, err := h.service.FetchAllMerchant(agentId)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenters.AgentMerchantStorefrontsSuccessResponse(results))

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
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(errors.New("error creating agent")))
		}
		return c.JSON(presenters.EmptySuccessResponse())
	}
}
func (h *AgentHandler) CreateCompliance() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.AgentComplianceRequest

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
		}
		agentId, _ := c.ParamsInt("agent")
		file, err := c.FormFile("policeReport")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		report, err := h.storageSrv.Disk("uploadcare").UploadFile("agent_police_report", file)
		if err != nil {
			h.storageSrv.Disk("uploadcare").ExecuteTask(report, "delete_file")
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": err,
				},
			)
		}
		personal, err := h.storageSrv.Disk("uploadcare").UploadFiles("agent_id", form.File["ghanaCardPersonal"])
		if err != nil {
			h.storageSrv.Disk("uploadcare").ExecuteTask(report, "delete_file")
			h.storageSrv.Disk("uploadcare").ExecuteTask(personal, "delete_files")
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": err,
				},
			)
		}
		guarantor, err := h.storageSrv.Disk("uploadcare").UploadFiles("guarantor_id", form.File["ghanaCard"])
		if err != nil {
			h.storageSrv.Disk("uploadcare").ExecuteTask(report, "delete_file")
			h.storageSrv.Disk("uploadcare").ExecuteTask(personal, "delete_files")
			h.storageSrv.Disk("uploadcare").ExecuteTask(guarantor, "delete_files")
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": err,
				},
			)
		}
		result, err := h.service.CreateCompliance(&request, agentId, report, personal, guarantor)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(errors.New("error creating agent")))
		}
		return c.JSON(presenters.AgentsComplianceSuccessResponse(result))
	}
}

func (h *AgentHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.AgentProfile

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
		}
		agentId, _ := c.ParamsInt("id")

		result, err := h.service.Update(agentId, &request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}
		return c.JSON(presenters.AgentSuccessResponse(result))
	}

}

func (h *AgentHandler) UpdateGuarantor() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.AgentGuarantorUpdate

		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
		}
		agentId, _ := c.ParamsInt("id")

		result, err := h.service.UpdateGuarantor(agentId, &request)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.CustomerErrorResponse(err))
		}

		return c.JSON(presenters.AgentsComplianceSuccessResponse(result))
	}

}

func (h *AgentHandler) UpdateAgentComplianceCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cardFile, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		newPath, err := h.storageSrv.Disk("uploadcare").UploadFile("agent_id", cardFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		agentId, _ := strconv.Atoi(c.Query("id"))
		oldPath := c.Query("file", "")
		result, err := h.service.UpdateAgentComplianceCard(agentId, newPath, oldPath)
		if err != nil {
			if oldPath != "" {
				h.storageSrv.Disk("uploadcare").ExecuteTask(oldPath, "delete_file")
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		if oldPath != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(oldPath, "delete_file")
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"data":   result,
			},
		)
	}

}
func (h *AgentHandler) UpdateAgentPoliceReport() fiber.Handler {
	return func(c *fiber.Ctx) error {
		reportFile, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		filePath, err := h.storageSrv.Disk("uploadcare").UploadFile("agent_police_report", reportFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		agentId, _ := strconv.Atoi(c.Query("id"))
		prevUrl := c.Query("file", "")
		result, err := h.service.UpdateAgentPoliceReport(agentId, filePath)
		if err != nil {
			if prevUrl != "" {
				h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		if prevUrl != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(prevUrl, "delete_file")
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"data":   result,
			},
		)
	}

}
func (h *AgentHandler) UpdateGuarantorComplianceCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cardFile, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.ProductErrorResponse(err))
		}
		newPath, err := h.storageSrv.Disk("uploadcare").UploadFile("guarantor_id", cardFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		agentId, _ := strconv.Atoi(c.Query("id"))
		oldPath := c.Query("file", "")
		result, err := h.service.UpdateGuarantorComplianceCard(agentId, newPath, oldPath)
		if err != nil {
			if oldPath != "" {
				h.storageSrv.Disk("uploadcare").ExecuteTask(oldPath, "delete_file")
			}
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"msg": "Upload error",
				},
			)
		}
		if oldPath != "" {
			h.storageSrv.Disk("uploadcare").ExecuteTask(oldPath, "delete_file")
		}
		return c.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"status": true,
				"data":   result,
			},
		)
	}

}
func (h *AgentHandler) ApproveAgent() fiber.Handler {
	return func(c *fiber.Ctx) error {
		agentId, _ := c.ParamsInt("agent")
		complianceStatus, _ := strconv.ParseBool(c.Params("status"))
		_, err := h.service.ApproveAgent(agentId, complianceStatus)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"status":  true,
				"message": "Updated",
			},
		)

	}

}

func (h *AgentHandler) SaveAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var momoRequest models.AgentMomoAccountRequest
		var bankRequest models.AgentBankAccountRequest

		agentId, _ := c.ParamsInt("id")
		accountType := c.Params("accountType")
		if accountType == "bank" {
			err := c.BodyParser(&bankRequest)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
			}
			result, err := h.service.SaveAccount(&bankRequest, agentId, accountType)
			if err != nil {

				return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(err))
			}
			return c.JSON(presenters.AgentSuccessResponse(result))
		}

		err := c.BodyParser(&momoRequest)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.AgentErrorResponse(err))
		}
		result, err := h.service.SaveAccount(&momoRequest, agentId, accountType)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(err))
		}

		return c.JSON(presenters.AgentSuccessResponse(result))

	}

}
func (h *AgentHandler) SaveDefaultAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {

		agentId, _ := c.ParamsInt("id")
		accountType := c.Params("accountType")

		result, err := h.service.SaveDefaultAccount(agentId, accountType)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.AgentErrorResponse(err))
		}

		return c.JSON(presenters.AgentSuccessResponse(result))

	}

}

func (h *AgentHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.service.Remove(c.Params("id")); err != nil {
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
