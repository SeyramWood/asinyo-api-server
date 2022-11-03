package merchant_store

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
)

type repository struct {
	db *ent.Client
}

func NewMerchantStoreRepo(db *database.Adapter) gateways.MerchantStoreRepo {
	return &repository{db.DB}
}

func (r repository) Insert(
	store *models.MerchantStore, merchantId int, logo string, images []string,
) (*ent.MerchantStore, error) {
	ctx := context.Background()
	mq := r.db.Merchant.Query().Where(merchant.ID(merchantId)).OnlyX(ctx)
	storeResult, err := r.db.MerchantStore.Create().SetMerchant(mq).
		SetName(store.BusinessName).
		SetAbout(store.About).
		SetDescTitle(store.DescriptionTitle).
		SetDescription(store.Description).
		SetLogo(logo).
		SetImages(images).
		SetRegion(store.Region).
		SetDistrict(store.District).
		SetCity(store.City).
		SetMerchantType(store.MerchantType).
		SetRegion(store.Region).
		SetDistrict(store.District).
		SetCity(store.City).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant store: %w", err)
	}

	result, err := r.Read(storeResult.ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r repository) UpdateAccount(store interface{}, storeId int, accountType string) (*ent.MerchantStore, error) {
	ctx := context.Background()
	if accountType == "bank" {
		account := store.(*models.MerchantBankAccountRequest)
		if account.DefaultAccount {
			result, err := r.db.MerchantStore.UpdateOneID(storeId).
				SetDefaultAccount("momo").
				SetBankAccount(
					&models.MerchantBankAccount{
						Name:   account.AccountName,
						Number: account.AccountNumber,
						Bank:   account.Bank,
						Branch: account.Branch,
					},
				).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to update merchant momo account : %w", err)
			}
			return result, nil
		}
		result, err := r.db.MerchantStore.UpdateOneID(storeId).
			SetBankAccount(
				&models.MerchantBankAccount{
					Name:   account.AccountName,
					Number: account.AccountNumber,
					Bank:   account.Bank,
					Branch: account.Branch,
				},
			).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed update merchant momo account : %w", err)
		}
		return result, nil
	}

	account := store.(*models.MerchantMomoAccountRequest)
	if account.DefaultAccount {
		result, err := r.db.MerchantStore.UpdateOneID(storeId).
			SetDefaultAccount("momo").
			SetMomoAccount(
				&models.MerchantMomoAccount{
					Name:     account.AccountName,
					Number:   account.PhoneNumber,
					Provider: account.Provider,
				},
			).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed update merchant momo account : %w", err)
		}
		return result, nil
	}
	result, err := r.db.MerchantStore.UpdateOneID(storeId).
		SetMomoAccount(
			&models.MerchantMomoAccount{
				Name:     account.AccountName,
				Number:   account.PhoneNumber,
				Provider: account.Provider,
			},
		).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to update merchant bank account : %w", err)
	}
	return result, nil

}

func (r repository) UpdateAgentPermission(permission bool, storeId int) (*ent.MerchantStore, error) {

	result, err := r.db.MerchantStore.UpdateOneID(storeId).
		SetPermitAgent(permission).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to update agent store permission : %w", err)
	}
	return result, nil
}

func (r repository) UpdateDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error) {
	ctx := context.Background()
	if accountType == "bank" {
		result, err := r.db.MerchantStore.UpdateOneID(storeId).
			SetDefaultAccount(merchantstore.DefaultAccountBank).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed update merchant default account : %w", err)
		}
		return result, nil
	}
	result, err := r.db.MerchantStore.UpdateOneID(storeId).
		SetDefaultAccount(merchantstore.DefaultAccountMomo).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed update merchant default account : %w", err)
	}
	return result, nil

}

func (r repository) Read(id int) (*ent.MerchantStore, error) {
	result, err := r.db.MerchantStore.Query().
		Where(merchantstore.ID(id)).
		WithMerchant().
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r repository) ReadByMerchant(merchantId int) (*ent.MerchantStore, error) {
	result, err := r.db.Merchant.Query().
		Where(merchant.ID(merchantId)).
		QueryStore().
		WithMerchant().
		WithAgent().
		Only(context.Background())
	if err != nil {
		return nil, nil
	}
	return result, nil
}
func (r repository) ReadAgent(store int) (*ent.Agent, error) {
	result, err := r.db.MerchantStore.Query().
		Where(merchantstore.ID(store)).
		QueryAgent().
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r repository) ReadAll() ([]*ent.MerchantStore, error) {
	// TODO implement me
	panic("implement me")
}
func (r repository) ReadAllByMerchant(merchantType string, limit, offset int) ([]*ent.MerchantStore, error) {
	if merchantType == "supplier" {
		results, err := r.db.MerchantStore.Query().
			Limit(limit).
			Offset(offset).
			Order(ent.Desc(merchantstore.FieldCreatedAt)).
			Where(merchantstore.MerchantType(merchantType)).
			WithMerchant(
				func(query *ent.MerchantQuery) {
					query.WithSupplier()
				},
			).
			All(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed fetching supplier stores : %w", err)
		}
		return results, nil
	}
	results, err := r.db.MerchantStore.Query().
		Order(ent.Desc(merchantstore.FieldCreatedAt)).
		Where(merchantstore.MerchantType(merchantType)).
		WithMerchant(
			func(query *ent.MerchantQuery) {
				query.WithRetailer()
			},
		).
		All(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed fetching retailer stores : %w", err)
	}
	return results, nil
}

func (r repository) Update(store *models.MerchantStore) (*models.MerchantStore, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}
