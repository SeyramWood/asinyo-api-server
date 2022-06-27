package gateways

import (
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type (
	CustomerRepo interface {
		Insert(customer *models.Customer) (*ent.Customer, error)
		Read(id int) (*ent.Customer, error)
		ReadAll() ([]*ent.Customer, error)
		Update(customer *models.Customer) (*models.Customer, error)
		Delete(id string) error
	}
	AgentRepo interface {
		Insert(agent *models.AgentRequest) (*ent.Agent, error)
		Read(id int) (*ent.Agent, error)
		ReadAll() ([]*ent.Agent, error)
		Update(agent *models.Agent) (*models.Agent, error)
		Delete(id string) error
	}
	MerchantRepo interface {
		Insert(merchant *models.MerchantRequest) (*ent.Merchant, error)
		Read(id int) (*ent.Merchant, error)
		ReadAll() ([]*ent.Merchant, error)
		Update(merchant *models.Merchant) (*models.Merchant, error)
		Delete(id string) error
	}
	SupplierMerchantRepo interface {
		Insert(merchant *models.SupplierMerchant) (*ent.SupplierMerchant, error)
		Read(id int) (*ent.SupplierMerchant, error)
		ReadAll() ([]*ent.SupplierMerchant, error)
		Update(merchant *models.SupplierMerchant) (*models.SupplierMerchant, error)
		Delete(id string) error
	}
	RetailMerchantRepo interface {
		Insert(merchant *models.RetailMerchant) (*ent.RetailMerchant, error)
		Read(id int) (*ent.RetailMerchant, error)
		ReadAll() ([]*ent.RetailMerchant, error)
		Update(merchant *models.RetailMerchant) (*models.RetailMerchant, error)
		Delete(id string) error
	}
	MerchantStoreRepo interface {
		Insert(store *models.MerchantStore, merchantId int, logo string, images []string) (*ent.MerchantStore, error)
		UpdateAccount(store interface{}, storeId int, logo string) (*ent.MerchantStore, error)
		UpdateDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error)
		Read(id int) (*ent.MerchantStore, error)
		ReadByMerchant(merchantId int) (*ent.MerchantStore, error)
		ReadAll() ([]*ent.MerchantStore, error)
		ReadAllByMerchant(merchantType string) ([]*ent.MerchantStore, error)
		Update(store *models.MerchantStore) (*models.MerchantStore, error)
		Delete(id string) error
	}
	ProductRepo interface {
		Insert(merchant *models.Product, imageUrl string) (*ent.Product, error)
		Read(id int) (*ent.Product, error)
		ReadBySupplierMerchant(id int) (*ent.Product, error)
		ReadByRetailMerchant(id int) (*ent.Product, error)
		ReadAll() ([]*ent.Product, error)
		ReadBySlugRetailMerchantCategoryMajor(slug string) ([]*ent.ProductCategoryMajor, error)
		ReadBySlugRetailMerchantCategoryMinor(slug string) ([]*ent.ProductCategoryMinor, error)
		ReadAllRetailMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error)
		ReadAllSupplierMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error)
		ReadAllMajorByRetailer(majorId int) ([]*ent.Product, error)
		ReadAllBySupplierMerchant(merchant int) ([]*ent.Product, error)
		ReadAllByRetailMerchant(merchant int) ([]*ent.Product, error)
		ReadBestSellerBySupplierMerchant() ([]*ent.Product, error)
		ReadBestSellerRetailMerchant() ([]*ent.Product, error)
		Update(merchant *models.Product) (*models.Product, error)
		Delete(id string) error
	}
	ProductCatMajorRepo interface {
		Insert(merchant *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		Read(id int) (*ent.ProductCategoryMajor, error)
		ReadAll() ([]*ent.ProductCategoryMajor, error)
		Update(merchant *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error)
		Delete(id string) error
	}
	ProductCatMinorRepo interface {
		Insert(merchant *models.ProductCategoryMinor) (*ent.ProductCategoryMinor, error)
		Read(id int) (*ent.ProductCategoryMinor, error)
		ReadAll() ([]*ent.ProductCategoryMinor, error)
		Update(merchant *models.ProductCategoryMinor) (*models.ProductCategoryMinor, error)
		Delete(id string) error
	}
	AdminRepo interface {
		Insert(user *models.Admin) (*ent.Admin, error)
		Read(id int) (*ent.Admin, error)
		ReadAll() ([]*ent.Admin, error)
		Update(user *models.Admin) (*models.Admin, error)
		Delete(id string) error
	}
	AddressRepo interface {
		Insert(user *models.Address, userId int, userType string) (*ent.Address, error)
		Read(id int) (*ent.Address, error)
		ReadAll() ([]*ent.Address, error)
		ReadAllByUser(userId int, userType string) ([]*ent.Address, error)
		ReadByUser(userId int, userType string) (*ent.Address, error)
		Update(addressId int, address *models.Address) (*ent.Address, error)
		UpdateByUserDefaultAddress(userId, addressId int, userType string) ([]*ent.Address, error)
		Delete(id string) error
	}
	PickupStationRepo interface {
		Insert(station *models.PickupStation) (*ent.PickupStation, error)
		Read(id int) (*ent.PickupStation, error)
		ReadAll() ([]*ent.PickupStation, error)
		Update(stationId int, address *models.PickupStation) (*ent.PickupStation, error)
		Delete(id string) error
	}

	OrderRepo interface {
		Insert(order *models.OrderResponse) error
		Read(id int) (*ent.Order, error)
		ReadByUser(userType string, id int) (*ent.Order, error)
		ReadAll() ([]*ent.Order, error)
		ReadByAllUser(userType string, id int) ([]*ent.Order, error)
		Update(order *models.OrderResponse) (*ent.Order, error)
		Delete(id string) error
	}

	AuthRepo interface {
		ReadAdmin(username, field string) (*ent.Admin, error)
		ReadCustomer(username, field string) (*ent.Customer, error)
		ReadAgent(username, field string) (*ent.Agent, error)
		ReadMerchant(username, field string) (*ent.Merchant, error)
	}

	PaymentRepo interface {
		Insert(transaction *models.Transaction) (*ent.Order, error)
	}
)
