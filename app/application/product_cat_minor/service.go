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

func (s *service) FetchAll(limit, offset int) ([]*ent.ProductCategoryMinor, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) Update(id int, request *models.ProductCategoryMinorUpdate) (*ent.ProductCategoryMinor, error) {
	return s.repo.Update(id, request)
}

func (s *service) UpdateImage(id int, imagePath string) (string, error) {
	return s.repo.UpdateImage(id, imagePath)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}
