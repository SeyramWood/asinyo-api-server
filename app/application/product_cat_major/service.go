package product_cat_major

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.ProductCatMajorRepo
}

func NewProductCatMajorService(repo gateways.ProductCatMajorRepo) gateways.ProductCatMajorService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(cat *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error) {
	return s.repo.Insert(cat)
}

func (s *service) Fetch(id int) (*ent.ProductCategoryMajor, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(id int, request *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error) {
	return s.repo.Update(id, request)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}
