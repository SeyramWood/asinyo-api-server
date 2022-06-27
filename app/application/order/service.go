package order

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.OrderRepo
}

func NewOrderService(repo gateways.OrderRepo) gateways.OrderService {
	return &service{repo: repo}
}

func (s service) Create(order *models.OrderResponse) error {
	return s.repo.Insert(order)
}

func (s service) FetchByAllUser(userType string, id int) ([]*ent.Order, error) {
	return s.repo.ReadByAllUser(userType, id)
}

func (s service) FetchByUser(userType string, id int) (*ent.Order, error) {
	return s.repo.ReadByUser(userType, id)
}
func (s service) FetchAll() ([]*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Fetch(id int) (*ent.Order, error) {
	return s.repo.Read(id)
}

func (s service) Update(order *models.OrderResponse) (*ent.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}
