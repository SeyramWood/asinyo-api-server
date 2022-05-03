package customer

import (
	"github.com/SeyramWood/app/adapters/gateways"
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

func (s *service) Create(customer *models.Customer) (*ent.Customer, error) {

	return s.repo.Insert(customer)
}

func (s *service) Fetch(id int) (*ent.Customer, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.Customer, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(user *models.Customer) (*models.Customer, error) {
	return s.repo.Update(user)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
