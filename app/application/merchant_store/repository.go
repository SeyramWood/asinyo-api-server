package merchant_store

import (
	"context"
	"fmt"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/merchant"
)

type repository struct {
	db *ent.Client
}

func NewMerchantStoreRepo(db *database.Adapter) gateways.MerchantStoreRepo {
	return &repository{db.DB}
}

func (r repository) Insert(store *models.MerchantStore, merchantId int, logo string, images []string) (*ent.Merchant, error) {
	ctx := context.Background()
	mq := r.db.Merchant.Query().Where(merchant.ID(merchantId)).OnlyX(ctx)
	_, err := r.db.MerchantStore.Create().SetMerchant(mq).
		SetName(store.BusinessName).
		SetAbout(store.About).
		SetDescTitle(store.DescriptionTitle).
		SetDescription(store.Description).
		SetLogo(logo).
		SetImages(images).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant store: %w", err)
	}

	result, err := r.Read(merchantId)
	//result.Edges.Store.Update()
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
				SetBankAccount(&models.MerchantBankAccount{
					Name:   account.AccountName,
					Number: account.AccountNumber,
					Bank:   account.Bank,
					Branch: account.Branch,
				}).
				Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed update merchant momo account : %w", err)
			}
			return result, nil
		}
		result, err := r.db.MerchantStore.UpdateOneID(storeId).
			SetBankAccount(&models.MerchantBankAccount{
				Name:   account.AccountName,
				Number: account.AccountNumber,
				Bank:   account.Bank,
				Branch: account.Branch,
			}).
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
			SetMomoAccount(&models.MerchantMomoAccount{
				Name:     account.AccountName,
				Number:   account.PhoneNumber,
				Provider: account.Provider,
			}).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed update merchant momo account : %w", err)
		}
		return result, nil
	}
	result, err := r.db.MerchantStore.UpdateOneID(storeId).
		SetMomoAccount(&models.MerchantMomoAccount{
			Name:     account.AccountName,
			Number:   account.PhoneNumber,
			Provider: account.Provider,
		}).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed update merchant bank account : %w", err)
	}
	return result, nil

}

func (r repository) Read(id int) (*ent.Merchant, error) {
	result, err := r.db.Merchant.Query().
		Where(merchant.ID(id)).
		WithStore().
		Only(context.Background())

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) ReadAll() ([]*ent.MerchantStore, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(store *models.MerchantStore) (*models.MerchantStore, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
