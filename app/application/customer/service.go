package customer

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.CustomerRepo
}

func NewCustomerService(repo gateways.CustomerRepo) gateways.CustomerService {
	return &service{
		repo: repo,
	}
}

func (s *service) CreatePurchaseRequest(customerId int, request *models.PurchaseOrderRequest) (
	*ent.PurchaseRequest, error,
) {
	return s.repo.InsertPurchaseRequest(customerId, request)
}

func (s *service) Create(customer any, customerType string) (*ent.Customer, error) {

	return s.repo.Insert(customer, customerType)
}

func (s *service) Fetch(id int) (*ent.Customer, error) {

	return s.repo.Read(id)
}

func (s *service) FetchPurchaseRequest(id int) (*ent.PurchaseRequest, error) {
	return s.repo.ReadPurchaseRequest(id)
}
func (s *service) FetchAllPurchaseRequestByCustomer(customerId, limit, offset int) (
	*presenters.PaginationResponse, error,
) {
	return s.repo.ReadAllPurchaseRequestByCustomer(customerId, limit, offset)
}
func (s *service) FetchAll(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) Update(id int, customer any) (*ent.Customer, error) {
	return s.repo.Update(id, customer)
}

func (s *service) UpdateLogo(customer int, logo string) (string, error) {
	return s.repo.UpdateLogo(customer, logo)
}

func (s *service) UpdatePurchaseRequest(id int, request *models.PurchaseOrderRequest) (*ent.PurchaseRequest, error) {
	return s.repo.UpdatePurchaseRequest(id, request)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

func (s *service) RemovePurchaseRequest(id int) error {
	// TODO implement me
	panic("implement me")
}
