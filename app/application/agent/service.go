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

func (s *service) Fetch(id int) (*ent.Agent, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.Agent, error) {
	return s.repo.ReadAll()
}

func (s *service) FetchAllMerchant(agentId int) ([]*ent.MerchantStore, error) {
	return s.repo.ReadAllMerchant(agentId)
}

func (s *service) Update(user *models.Agent) (*models.Agent, error) {
	return s.repo.Update(user)
}

func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}

func (s *service) CreateCompliance(
	request *models.AgentComplianceRequest, id int, report string, personal []string, guarantor []string,
) (*ent.Agent, error) {
	return s.repo.CreateCompliance(request, id, report, personal, guarantor)
}
