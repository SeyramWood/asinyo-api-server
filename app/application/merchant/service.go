package merchant

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.MerchantRepo
}

func NewMerchantService(repo gateways.MerchantRepo) gateways.MerchantService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(merchant *models.MerchantRequest) (*ent.Merchant, error) {

	return s.repo.Insert(merchant)
}

func (s *service) Fetch(id int) (*ent.Merchant, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.Merchant, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(user *models.Merchant) (*models.Merchant, error) {
	return s.repo.Update(user)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
