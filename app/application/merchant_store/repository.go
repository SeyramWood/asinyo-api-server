package merchant_store

import (
	"context"
	"fmt"

	"github.com/samber/lo"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
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
	store *models.MerchantStoreRequest, merchantId int, logo string, images []string,
) (*ent.MerchantStore, error) {
	ctx := context.Background()
	mq := r.db.Merchant.Query().Where(merchant.ID(merchantId)).OnlyX(ctx)
	storeResult, err := r.db.MerchantStore.Create().SetMerchant(mq).
		SetName(store.Info.BusinessName).
		SetSlogan(store.Info.BusinessSlogan).
		SetAbout(store.Info.About).
		SetDescription(store.Info.Description).
		SetLogo(logo).
		SetImages(images).
		SetMerchantType(store.Info.MerchantType).
		SetAddress(store.Address).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant store: %w", err)
	}

	result, errr := r.Read(storeResult.ID)

	if errr != nil {
		return nil, err
	}

	return result, nil
}
func (r repository) SaveCoordinate(coordinate *services.Coordinate, id int) error {
	_, err := r.db.MerchantStore.UpdateOneID(id).SetCoordinate(coordinate).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r repository) UpdateAccount(account any, storeId int, accountType string) (*ent.MerchantStore, error) {
	ctx := context.Background()
	if accountType == "bank" {
		request := account.(*models.MerchantBankAccountRequest)
		if request.DefaultAccount {
			result, err := r.db.MerchantStore.UpdateOneID(storeId).
				SetDefaultAccount("bank").
				SetBankAccount(
					&models.MerchantBankAccount{
						Name:   request.AccountName,
						Number: request.AccountNumber,
						Bank:   request.Bank,
						Branch: request.Branch,
					},
				).
				Save(ctx)
			if err != nil {
				return nil, err
			}
			return result, nil
		}

		result, err := r.db.MerchantStore.UpdateOneID(storeId).
			SetBankAccount(
				&models.MerchantBankAccount{
					Name:   request.AccountName,
					Number: request.AccountNumber,
					Bank:   request.Bank,
					Branch: request.Branch,
				},
			).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	request := account.(*models.MerchantMomoAccountRequest)
	if request.DefaultAccount {
		result, err := r.db.MerchantStore.UpdateOneID(storeId).
			SetDefaultAccount("momo").
			SetMomoAccount(
				&models.MerchantMomoAccount{
					Name:     request.AccountName,
					Number:   request.PhoneNumber,
					Provider: request.Provider,
				},
			).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	result, err := r.db.MerchantStore.UpdateOneID(storeId).
		SetMomoAccount(
			&models.MerchantMomoAccount{
				Name:     request.AccountName,
				Number:   request.PhoneNumber,
				Provider: request.Provider,
			},
		).
		Save(ctx)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("failed to update merchant default account: %w", err)
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

func (r repository) Update(request *models.MerchantStoreUpdate, storeId int) (*ent.MerchantStore, error) {
	result, err := r.db.MerchantStore.UpdateOneID(storeId).
		SetName(request.BusinessName).
		SetSlogan(request.BusinessSlogan).
		SetAbout(request.About).
		SetDescription(request.Description).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed to update merchant storefront: %w", err)
	}
	return result, nil
}
func (r repository) UpdateAddress(address *models.MerchantStoreAddress, storeId int) (*ent.MerchantStore, error) {

	result, err := r.db.MerchantStore.UpdateOneID(storeId).SetAddress(address).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r repository) UpdateBanner(storeId int, bannerPath string) (string, error) {
	ctx := context.Background()
	_, err := r.db.MerchantStore.UpdateOneID(storeId).SetLogo(bannerPath).Save(ctx)
	if err != nil {
		return "", err
	}
	return bannerPath, nil
}

func (r repository) UpdateImages(storeId int, newPath, oldPath string) ([]string, error) {
	ctx := context.Background()
	old := r.db.MerchantStore.Query().Where(merchantstore.ID(storeId)).OnlyX(ctx)
	newImages := lo.Map[string](
		old.Images, func(path string, index int) string {
			if path == oldPath {
				return newPath
			}
			return path
		},
	)
	_, err := old.Update().
		SetImages(newImages).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newImages, nil
}

func (r repository) AppendNewImages(storeId int, urls []string) ([]string, error) {
	ctx := context.Background()
	old := r.db.MerchantStore.Query().Where(merchantstore.ID(storeId)).OnlyX(ctx)
	var newImages []string
	newImages = append(newImages, old.Images...)
	newImages = append(newImages, urls...)

	_, err := old.Update().
		SetImages(newImages).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return newImages, nil
}
func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}
