package merchant_store

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.MerchantStoreRepo
}

func NewMerchantStoreService(repo gateways.MerchantStoreRepo) gateways.MerchantStoreService {
	return &service{
		repo: repo,
	}
}

func (s service) Create(store *models.MerchantStore, merchantId int, logo string, images []string) (*ent.MerchantStore, error) {
	return s.repo.Insert(store, merchantId, logo, images)
}

func (s service) SaveAccount(store interface{}, storeId int, accountType string) (*ent.MerchantStore, error) {
	return s.repo.UpdateAccount(store, storeId, accountType)
}
func (s service) SaveDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error) {
	return s.repo.UpdateDefaultAccount(storeId, accountType)
}

func (s service) FetchAll() ([]*ent.MerchantStore, error) {
	return s.repo.ReadAll()
}
func (s service) FetchAllByMerchant(merchantType string) ([]*ent.MerchantStore, error) {
	return s.repo.ReadAllByMerchant(merchantType)
}

func (s service) Fetch(id int) (*ent.MerchantStore, error) {
	return s.repo.Read(id)
}
func (s service) FetchByMerchant(merchantId int) (*ent.MerchantStore, error) {
	return s.repo.ReadByMerchant(merchantId)
}

func (s service) Update(store *models.MerchantStore) (*models.MerchantStore, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}
