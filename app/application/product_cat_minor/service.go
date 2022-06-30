package product_cat_minor

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.ProductCatMinorRepo
}

func NewProductCatMinorService(repo gateways.ProductCatMinorRepo) gateways.ProductCatMinorService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(cat *models.ProductCategoryMinor, image string) (*ent.ProductCategoryMinor, error) {

	return s.repo.Insert(cat, image)
}

func (s *service) Fetch(id int) (*ent.ProductCategoryMinor, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.ProductCategoryMinor, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(user *models.ProductCategoryMinor) (*models.ProductCategoryMinor, error) {
	return s.repo.Update(user)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
