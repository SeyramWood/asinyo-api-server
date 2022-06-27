package pickup_station

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.PickupStationRepo
}

func NewPickupStationService(repo gateways.PickupStationRepo) gateways.PickupStationService {
	return &service{
		repo: repo,
	}
}

func (s service) Create(station *models.PickupStation) (*ent.PickupStation, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) FetchAll() ([]*ent.PickupStation, error) {
	return s.repo.ReadAll()
}

func (s service) Fetch(id int) (*ent.PickupStation, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(station *models.PickupStation) (*ent.PickupStation, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}
