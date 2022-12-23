package agent

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.AgentRepo
}

func NewAgentService(repo gateways.AgentRepo) gateways.AgentService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(agent *models.AgentRequest) (*ent.Agent, error) {

	return s.repo.Insert(agent)
}

func (s *service) CreateCompliance(
	request *models.AgentComplianceRequest, id int, report string, personal []string, guarantor []string,
) (*ent.Agent, error) {
	return s.repo.CreateCompliance(request, id, report, personal, guarantor)
}

func (s *service) Fetch(id int) (*ent.Agent, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.Agent, error) {
	return s.repo.ReadAll()
}

func (s *service) FetchAllMerchant(agentId int) ([]*ent.MerchantStore, error) {
	return s.repo.ReadAllMerchant(agentId)
}

func (s *service) Update(id int, profile *models.AgentProfile) (*ent.Agent, error) {
	return s.repo.Update(id, profile)
}
func (s *service) UpdateGuarantor(id int, request *models.AgentGuarantorUpdate) (*ent.Agent, error) {
	return s.repo.UpdateGuarantor(id, request)
}

func (s *service) UpdateAgentComplianceCard(agentId int, newPath, oldPath string) ([]string, error) {
	return s.repo.UpdateAgentComplianceCard(agentId, newPath, oldPath)
}

func (s *service) UpdateAgentPoliceReport(agentId int, filePath string) (string, error) {
	return s.repo.UpdateAgentPoliceReport(agentId, filePath)
}

func (s *service) UpdateGuarantorComplianceCard(agentId int, newPath, oldPath string) ([]string, error) {
	return s.repo.UpdateGuarantorComplianceCard(agentId, newPath, oldPath)
}

func (s *service) SaveAccount(account any, agentId int, accountType string) (*ent.Agent, error) {
	return s.repo.UpdateAccount(account, agentId, accountType)
}

func (s *service) SaveDefaultAccount(agentId int, accountType string) (*ent.Agent, error) {
	return s.repo.UpdateDefaultAccount(agentId, accountType)
}

func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
