package auth

import (
	"context"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

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

func NewAuthRepo(db *database.Adapter) gateways.AuthRepo {
	return &repository{db.DB}
}

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
		user, err := r.db.Customer.Query().Where(customer.ID(id)).
			WithBusiness().
			WithIndividual().
			Only(context.Background())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := r.db.Customer.Query().Where(customer.Username(username)).
		WithBusiness().
		WithIndividual().
		Only(context.Background())
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
		user, err := r.db.Merchant.Query().Where(merchant.ID(id)).WithSupplier().WithRetailer().
			Only(context.Background())
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	user, err := r.db.Merchant.Query().Where(merchant.Username(username)).WithSupplier().WithRetailer().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *repository) UpdatePassword(id int, password string, userType string, isOTP bool) (bool, error) {
	ctx := context.Background()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 16)
	switch userType {
	case "business", "individual":
		_, err := r.db.Customer.UpdateOneID(id).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	case "agent":
		_, err := r.db.Agent.UpdateOneID(id).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	case "supplier", "retailer":
		if isOTP {
			_, err := r.db.Merchant.UpdateOneID(id).
				SetPassword(hashPassword).
				SetOtp(false).
				Save(ctx)
			if err != nil {
				return false, fmt.Errorf("failed to update password")
			}
			return true, nil
		}
		_, err := r.db.Merchant.UpdateOneID(id).
			SetPassword(hashPassword).
			Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	case "asinyo":
		_, err := r.db.Admin.UpdateOneID(id).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	default:
		return false, nil
	}
}
func (r *repository) ResetPassword(id int, password, userType string) (bool, error) {
	ctx := context.Background()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 16)
	switch userType {
	case "customer":
		_, err := r.db.Customer.UpdateOneID(id).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	case "agent":
		_, err := r.db.Agent.UpdateOneID(id).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	case "merchant":
		_, err := r.db.Merchant.UpdateOneID(id).
			SetPassword(hashPassword).
			Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	case "asinyo":
		_, err := r.db.Admin.UpdateOneID(id).SetPassword(hashPassword).Save(ctx)
		if err != nil {
			return false, fmt.Errorf("failed to update password")
		}
		return true, nil
	default:
		return false, nil
	}
}
