package retail_merchant

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.RetailMerchantRepo
}

func NewRetailMerchantService(repo gateways.RetailMerchantRepo) gateways.RetailMerchantService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(customer *models.RetailMerchant) (*ent.RetailMerchant, error) {

	return s.repo.Insert(customer)
}

func (s *service) Fetch(id int) (*ent.RetailMerchant, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.RetailMerchant, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(user *models.RetailMerchant) (*models.RetailMerchant, error) {
	return s.repo.Update(user)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
