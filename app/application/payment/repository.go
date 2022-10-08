package payment

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewPaymentRepo(db *database.Adapter) gateways.PaymentRepo {
	return &repository{db.DB}
}

func (r repository) Insert(transaction *services.Transaction) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}
