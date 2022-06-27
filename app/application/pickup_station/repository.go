package pickup_station

import (
	"context"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewPickupStationRepo(db *database.Adapter) gateways.PickupStationRepo {
	return &repository{db.DB}
}

func (r repository) Insert(station *models.PickupStation) (*ent.PickupStation, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Read(id int) (*ent.PickupStation, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) ReadAll() ([]*ent.PickupStation, error) {
	b, err := r.db.PickupStation.Query().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r repository) Update(stationId int, address *models.PickupStation) (*ent.PickupStation, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
