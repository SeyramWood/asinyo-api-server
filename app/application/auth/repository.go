package auth

import (
	"context"
	"strconv"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/admin"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/merchant"
)

type repository struct {
	db *ent.Client
}

//NewRepo is the single instance repo that is being created.
func NewAuthRepo(db *database.Adapter) gateways.AuthRepo {
	return &repository{db.DB}
}

//ReadUser is a mongo repository that helps to fetch books
func (r *repository) ReadAdmin(username, field string) (*ent.Admin, error) {
	if field == "id" {
		id, _ := strconv.Atoi(username)
		user, err := r.db.Admin.Query().Where(admin.ID(id)).First(context.Background())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := r.db.Admin.Query().Where(admin.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *repository) ReadCustomer(username, field string) (*ent.Customer, error) {

	if field == "id" {
		id, _ := strconv.Atoi(username)
		user, err := r.db.Customer.Query().Where(customer.ID(id)).First(context.Background())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := r.db.Customer.Query().Where(customer.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *repository) ReadAgent(username, field string) (*ent.Agent, error) {
	if field == "id" {
		id, _ := strconv.Atoi(username)
		user, err := r.db.Agent.Query().Where(agent.ID(id)).First(context.Background())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := r.db.Agent.Query().Where(agent.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *repository) ReadMerchant(username, field string) (*ent.Merchant, error) {
	if field == "id" {
		id, _ := strconv.Atoi(username)
		user, err := r.db.Merchant.Query().Where(merchant.ID(id)).
			First(context.Background())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := r.db.Merchant.Query().Where(merchant.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
