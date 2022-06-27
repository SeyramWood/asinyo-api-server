package gateways

import (
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	CustomerService interface {
		Create(customer *models.Customer) (*ent.Customer, error)
		FetchAll() ([]*ent.Customer, error)
		Fetch(id int) (*ent.Customer, error)
		Update(customer *models.Customer) (*models.Customer, error)
		Remove(id string) error
	}
	AgentService interface {
		Create(agent *models.AgentRequest) (*ent.Agent, error)
		FetchAll() ([]*ent.Agent, error)
		Fetch(id int) (*ent.Agent, error)
		Update(agent *models.Agent) (*models.Agent, error)
		Remove(id string) error
	}
	MerchantService interface {
		Create(merchant *models.MerchantRequest) (*ent.Merchant, error)
		FetchAll() ([]*ent.Merchant, error)
		Fetch(id int) (*ent.Merchant, error)
		Update(merchant *models.Merchant) (*models.Merchant, error)
		Remove(id string) error
	}
	SupplierMerchantService interface {
		Create(merchant *models.SupplierMerchant) (*ent.SupplierMerchant, error)
		FetchAll() ([]*ent.SupplierMerchant, error)
		Fetch(id int) (*ent.SupplierMerchant, error)
		Update(merchant *models.SupplierMerchant) (*models.SupplierMerchant, error)
		Remove(id string) error
	}
	RetailMerchantService interface {
		Create(merchant *models.RetailMerchant) (*ent.RetailMerchant, error)
		FetchAll() ([]*ent.RetailMerchant, error)
		Fetch(id int) (*ent.RetailMerchant, error)
		Update(merchant *models.RetailMerchant) (*models.RetailMerchant, error)
		Remove(id string) error
	}
	MerchantStoreService interface {
		Create(store *models.MerchantStore, merchantId int, logo string, images []string) (*ent.MerchantStore, error)
		SaveAccount(store interface{}, storeId int, logo string) (*ent.MerchantStore, error)
		SaveDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error)
		FetchAll() ([]*ent.MerchantStore, error)
		FetchAllByMerchant(merchantType string) ([]*ent.MerchantStore, error)
		Fetch(id int) (*ent.MerchantStore, error)
		FetchByMerchant(merchantId int) (*ent.MerchantStore, error)
		Update(store *models.MerchantStore) (*models.MerchantStore, error)
		Remove(id string) error
	}

	ProductService interface {
		Create(merchant *models.Product, imageUrl string) (*ent.Product, error)
		FetchAll() ([]*ent.Product, error)
		FetchBySlugRetailMerchantCategoryMajor(slug string) ([]*ent.ProductCategoryMajor, error)
		FetchBySlugRetailMerchantCategoryMinor(slug string) ([]*ent.ProductCategoryMinor, error)
		FetchAllRetailMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error)
		FetchAllSupplierMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error)
		FetchAllMajorByRetailer(majorId int) ([]*ent.Product, error)
		FetchAllBySupplier(supplier int) ([]*ent.Product, error)
		FetchAllByRetailer(retailer int) ([]*ent.Product, error)
		FetchBestSellerBySupplier() ([]*ent.Product, error)
		FetchBestSellerByRetailer() ([]*ent.Product, error)
		Fetch(id int) (*ent.Product, error)
		FetchBySupplierMerchant(id int) (*ent.Product, error)
		FetchByRetailMerchant(id int) (*ent.Product, error)
		Update(merchant *models.Product) (*models.Product, error)
		Remove(id string) error
	}
	ProductCatMajorService interface {
		Create(merchant *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		FetchAll() ([]*ent.ProductCategoryMajor, error)
		Fetch(id int) (*ent.ProductCategoryMajor, error)
		Update(merchant *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error)
		Remove(id string) error
	}
	ProductCatMinorService interface {
		Create(merchant *models.ProductCategoryMinor) (*ent.ProductCategoryMinor, error)
		FetchAll() ([]*ent.ProductCategoryMinor, error)
		Fetch(id int) (*ent.ProductCategoryMinor, error)
		Update(merchant *models.ProductCategoryMinor) (*models.ProductCategoryMinor, error)
		Remove(id string) error
	}
	AddressService interface {
		Create(address *models.Address, userId int, userType string) (*ent.Address, error)
		FetchAll() ([]*ent.Address, error)
		FetchAllByUser(userId int, userType string) ([]*ent.Address, error)
		FetchByUser(userId int, userType string) (*ent.Address, error)
		Fetch(id int) (*ent.Address, error)
		Update(addressId int, address *models.Address) (*ent.Address, error)
		SaveDefaultAddress(userId, addressId int, userType string) ([]*ent.Address, error)
		Remove(id string) error
	}
	PickupStationService interface {
		Create(station *models.PickupStation) (*ent.PickupStation, error)
		FetchAll() ([]*ent.PickupStation, error)
		Fetch(id int) (*ent.PickupStation, error)
		Update(station *models.PickupStation) (*ent.PickupStation, error)
		Remove(id string) error
	}
	OrderService interface {
		Create(order *models.OrderResponse) error
		FetchAll() ([]*ent.Order, error)
		FetchByAllUser(userType string, id int) ([]*ent.Order, error)
		Fetch(id int) (*ent.Order, error)
		FetchByUser(userType string, id int) (*ent.Order, error)
		Update(order *models.OrderResponse) (*ent.Order, error)
		Remove(id string) error
	}
	AdminService interface {
		Create(user *models.Admin) (*ent.Admin, error)
		FetchAll() ([]*ent.Admin, error)
		Fetch(id int) (*ent.Admin, error)
		Update(user *models.Admin) (*models.Admin, error)
		Remove(id string) error
	}

	AuthService interface {
		Login(c *fiber.Ctx) error
		Logout(c *fiber.Ctx) error
		FetchAuthUser(c *fiber.Ctx) error
	}

	PaymentService interface {
		Pay(request interface{}) (*http.Response, error)
		Verify(reference string) (*http.Response, error)
	}
)
