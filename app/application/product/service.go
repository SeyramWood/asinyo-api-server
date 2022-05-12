package product

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.ProductRepo
}

func NewProductService(repo gateways.ProductRepo) gateways.ProductService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(merchant *models.Product) (*ent.Product, error) {

	return s.repo.Insert(merchant)
}

func (s *service) Fetch(id int) (*ent.Product, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.Product, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(user *models.Product) (*models.Product, error) {
	return s.repo.Update(user)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
