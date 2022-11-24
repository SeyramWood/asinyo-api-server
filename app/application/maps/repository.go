package maps

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewMapRepo(db *database.Adapter) gateways.MapRepo {
	return &repository{
		db: db.DB,
	}
}

func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}

func (r repository) SaveCoordinate(coordinate *services.Coordinate, id int, model string) error {
	// if model == "address" {
	// 	r.db.Address.UpdateOneID(address.ID(id))
	// }
	return nil
}
