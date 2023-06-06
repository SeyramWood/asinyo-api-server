package price_model

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.PriceModelRepo
}

func NewPriceModelService(repo gateways.PriceModelRepo) gateways.PriceModelService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(model *models.PriceModelRequest) (*ent.PriceModel, error) {
	return s.repo.Insert(model)
}

func (s *service) Fetch(id int) (*ent.PriceModel, error) {
	return s.repo.Read(id)
}

func (s *service) FetchAll(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) FetchAllPercentage(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllPercentage(limit, offset)
}

func (s *service) Update(id int, model *models.PriceModelRequest) (*ent.PriceModel, error) {
	return s.repo.Update(id, model)
}

func (s *service) UpdatePercentage(category, percentage int) (*ent.ProductCategoryMinor, error) {
	return s.repo.UpdatePercentage(category, percentage)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}

func (s *service) RemovePercentage(id int) error {
	return s.repo.DeletePercentage(id)
}
