package address

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/merchant"
)

type repository struct {
	db *ent.Client
}

func NewAddressRepo(db *database.Adapter) gateways.AddressRepo {
	return &repository{db.DB}
}

func (r repository) Insert(address *models.Address, userId int, userType string) (*ent.Address, error) {
	switch userType {
	case "retailer", "supplier":
		return r.insertMerchantAddress(address, userId)
	case "agent":
		return r.insertAgentAddress(address, userId)
	default:
		return r.insertCustomerAddress(address, userId)
	}
}

func (r repository) Read(id int) (*ent.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) ReadAll() ([]*ent.Address, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) ReadAllByUser(userId int, userType string) ([]*ent.Address, error) {
	ctx := context.Background()
	switch userType {
	case "retailer", "supplier":
		result, err := r.db.Address.Query().Where(address.HasMerchantWith(merchant.ID(userId))).Order(ent.Desc(address.FieldDefault)).All(ctx)
		if err != nil {
			return nil, nil
		}
		return result, nil
	case "agent":
		result, err := r.db.Address.Query().Where(address.HasAgentWith(agent.ID(userId))).Order(ent.Desc(address.FieldDefault)).All(ctx)
		if err != nil {
			return nil, nil
		}
		return result, nil
	default:
		result, err := r.db.Address.Query().Where(address.HasCustomerWith(customer.ID(userId))).Order(ent.Desc(address.FieldDefault)).All(ctx)
		if err != nil {
			return nil, nil
		}
		return result, nil
	}
}

func (r repository) ReadByUser(userId int, userType string) (*ent.Address, error) {
	ctx := context.Background()
	switch userType {
	case "retailer", "supplier":
		result, err := r.db.Address.Query().Where(
			address.HasMerchantWith(merchant.ID(userId)),
			address.Default(true),
		).
			Only(ctx)
		if err != nil {
			return nil, nil
		}
		return result, nil
	case "agent":
		result, err := r.db.Address.Query().Where(
			address.HasAgentWith(agent.ID(userId)),
			address.Default(true),
		).
			Only(ctx)
		if err != nil {
			return nil, nil
		}
		return result, nil
	default:
		result, err := r.db.Address.Query().Where(
			address.HasCustomerWith(customer.ID(userId)),
			address.Default(true),
		).
			Only(ctx)
		if err != nil {
			return nil, nil
		}
		return result, nil
	}
}

func (r repository) Update(addressId int, addr *models.Address) (*ent.Address, error) {
	result, err := r.db.Address.UpdateOneID(addressId).
		SetAddress(addr.Address).
		SetRegion(addr.Region).
		SetDistrict(addr.District).
		SetCity(addr.City).
		SetStreetName(addr.StreetName).
		SetStreetNumber(addr.StreetNumber).
		SetLastName(addr.LastName).
		SetOtherName(addr.OtherName).
		SetPhone(addr.Phone).
		SetOtherPhone(addr.OtherPhone).
		SetDefault(addr.Default).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil

}
func (r repository) UpdateByUserDefaultAddress(userId, addressId int, userType string) ([]*ent.Address, error) {
	ctx := context.Background()
	switch userType {
	case "retailer", "supplier":
		_, allErr := r.db.Address.Update().Where(address.HasMerchantWith(merchant.ID(userId))).SetDefault(false).Save(ctx)
		if allErr != nil {
			return nil, allErr
		}
		_, oneErr := r.db.Address.UpdateOneID(addressId).SetDefault(true).Save(ctx)
		if oneErr != nil {
			return nil, oneErr
		}
		return r.ReadAllByUser(userId, userType)
	case "agent":
		_, allErr := r.db.Address.Update().Where(address.HasAgentWith(agent.ID(userId))).SetDefault(false).Save(ctx)
		if allErr != nil {
			return nil, allErr
		}
		_, oneErr := r.db.Address.UpdateOneID(addressId).SetDefault(true).Save(ctx)
		if oneErr != nil {
			return nil, oneErr
		}
		return r.ReadAllByUser(userId, userType)
	default:
		_, allErr := r.db.Address.Update().Where(address.HasCustomerWith(customer.ID(userId))).SetDefault(false).Save(ctx)
		if allErr != nil {
			return nil, allErr
		}
		_, oneErr := r.db.Address.UpdateOneID(addressId).SetDefault(true).Save(ctx)
		if oneErr != nil {
			return nil, oneErr
		}
		return r.ReadAllByUser(userId, userType)
	}
}
func (r repository) SaveCoordinate(coordinate *services.Coordinate, id int) error {
	_, err := r.db.Address.UpdateOneID(id).SetCoordinate(coordinate).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}
func (r repository) insertMerchantAddress(addr *models.Address, userId int) (*ent.Address, error) {
	ctx := context.Background()
	mq := r.db.Merchant.Query().Where(merchant.ID(userId)).OnlyX(ctx)
	result, err := r.db.Address.Create().
		SetMerchant(mq).
		SetAddress(addr.Address).
		SetRegion(addr.Region).
		SetDistrict(addr.District).
		SetCity(addr.City).
		SetStreetName(addr.StreetName).
		SetStreetNumber(addr.StreetNumber).
		SetLastName(addr.LastName).
		SetOtherName(addr.OtherName).
		SetPhone(addr.Phone).
		SetOtherPhone(addr.OtherPhone).
		SetDefault(addr.Default).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant address: %w", err)
	}
	return result, nil
}
func (r repository) insertAgentAddress(addr *models.Address, userId int) (*ent.Address, error) {
	ctx := context.Background()
	mq := r.db.Agent.Query().Where(agent.ID(userId)).OnlyX(ctx)
	result, err := r.db.Address.Create().
		SetAgent(mq).
		SetAddress(addr.Address).
		SetRegion(addr.Region).
		SetDistrict(addr.District).
		SetCity(addr.City).
		SetStreetName(addr.StreetName).
		SetStreetNumber(addr.StreetNumber).
		SetLastName(addr.LastName).
		SetOtherName(addr.OtherName).
		SetPhone(addr.Phone).
		SetOtherPhone(addr.OtherPhone).
		SetDefault(addr.Default).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant address: %w", err)
	}
	return result, nil
}
func (r repository) insertCustomerAddress(addr *models.Address, userId int) (*ent.Address, error) {
	ctx := context.Background()
	mq := r.db.Customer.Query().Where(customer.ID(userId)).OnlyX(ctx)
	result, err := r.db.Address.Create().
		SetCustomer(mq).
		SetAddress(addr.Address).
		SetRegion(addr.Region).
		SetDistrict(addr.District).
		SetCity(addr.City).
		SetStreetName(addr.StreetName).
		SetStreetNumber(addr.StreetNumber).
		SetLastName(addr.LastName).
		SetOtherName(addr.OtherName).
		SetPhone(addr.Phone).
		SetOtherPhone(addr.OtherPhone).
		SetDefault(addr.Default).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant address: %w", err)
	}
	return result, nil
}
