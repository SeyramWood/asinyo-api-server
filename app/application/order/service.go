package order

import (
	"github.com/Jeffail/gabs"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.OrderRepo
}

func NewOrderService(repo gateways.OrderRepo) gateways.OrderService {
	return &service{repo: repo}
}

func (s service) Create(order *models.OrderPayload) (*ent.Order, error) {

	return s.repo.Insert(order)

}

func (s service) FetchAllByUser(userType string, id int) ([]*ent.Order, error) {
	return s.repo.ReadAllByUser(userType, id)
}
func (s service) FetchAllByStore(merchantId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByStore(merchantId)
}
func (s service) FetchAllByAgentStore(agentId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByAgentStore(agentId)
}
func (s service) FetchByStore(id, userId int, userType string) (*ent.Order, error) {
	if userType == "agent" {
		return s.repo.ReadByAgentStore(id, userId)
	}
	return s.repo.ReadByStore(id, userId)
}

func (s service) FetchByUser(userType string, id int) (*ent.Order, error) {
	return s.repo.ReadByUser(userType, id)
}
func (s service) FetchAll() ([]*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s service) Fetch(id int) (*ent.Order, error) {
	return s.repo.Read(id)
}

func (s service) Update(order *services.PaystackResponse) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s service) Remove(id string) error {
	// TODO implement me
	panic("implement me")
}

func (s service) UpdateOrderDetailStatus(request []byte) (*ent.Order, error) {
	resBody, err := gabs.ParseJSON(request)
	if err != nil {
		return nil, err
	}
	statuses, errr := resBody.ChildrenMap()

	if errr != nil {
		return nil, errr
	}

	return s.repo.UpdateOrderDetailStatus(statuses)
}
