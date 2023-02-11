package gateways

import (
	"github.com/Jeffail/gabs"

	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent"
)

type (
	CustomerRepo interface {
		Insert(customer any, customerType string) (*ent.Customer, error)
		Read(id int) (*ent.Customer, error)
		ReadAll() ([]*ent.Customer, error)
		Update(id int, customer any) (*ent.Customer, error)
		UpdateLogo(customer int, logo string) (string, error)
		Delete(id string) error
	}
	AgentRepo interface {
		Insert(agent *models.AgentRequest) (*ent.Agent, error)
		CreateCompliance(
			request *models.AgentComplianceRequest, id int, report string, personal []string, guarantor []string,
		) (*ent.Agent, error)
		UpdateAccount(account any, agentId int, accountType string) (*ent.Agent, error)
		UpdateDefaultAccount(agentId int, accountType string) (*ent.Agent, error)
		Read(id int) (*ent.Agent, error)
		ReadAll() ([]*ent.Agent, error)
		ReadAllMerchant(agentId int) ([]*ent.MerchantStore, error)
		Update(id int, profile *models.AgentProfile) (*ent.Agent, error)
		UpdateGuarantor(id int, request *models.AgentGuarantorUpdate) (*ent.Agent, error)
		UpdateAgentComplianceCard(agentId int, newPath, oldPath string) ([]string, error)
		UpdateAgentPoliceReport(agentId int, filePath string) (string, error)
		UpdateGuarantorComplianceCard(agentId int, newPath, oldPath string) ([]string, error)
		Delete(id string) error
	}
	MerchantRepo interface {
		Insert(merchant *models.MerchantRequest, onboard bool) (*ent.Merchant, error)
		Onboard(
			merchant *models.OnboardMerchantFullRequest, agentId int, logo string, images []string, password string,
		) (*ent.Merchant, error)
		SaveCoordinate(coordinate *services.Coordinate, id int) error
		Read(id int) (*ent.Merchant, error)
		ReadAll() ([]*ent.Merchant, error)
		Update(id int, request any) (*ent.Merchant, error)
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
		Insert(store *models.MerchantStoreRequest, merchantId int, logo string, images []string) (
			*ent.MerchantStore, error,
		)
		UpdateAccount(store any, storeId int, logo string) (*ent.MerchantStore, error)
		UpdateDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error)
		UpdateAgentPermission(permission bool, storeId int) (*ent.MerchantStore, error)
		SaveCoordinate(coordinate *services.Coordinate, id int) error
		Read(id int) (*ent.MerchantStore, error)
		ReadByMerchant(merchantId int) (*ent.MerchantStore, error)
		ReadAgent(store int) (*ent.Agent, error)
		ReadAll() ([]*ent.MerchantStore, error)
		ReadAllByMerchant(merchantType string, limit, offset int) ([]*ent.MerchantStore, error)
		Update(request *models.MerchantStoreUpdate, storeId int) (*ent.MerchantStore, error)
		UpdateAddress(address *models.MerchantStoreAddress, storeId int) (*ent.MerchantStore, error)
		UpdateBanner(storeId int, bannerPath string) (string, error)
		UpdateImages(storeId int, newPath, oldPath string) ([]string, error)
		AppendNewImages(storeId int, urls []string) ([]string, error)
		Delete(id string) error
	}
	ProductRepo interface {
		Insert(merchant *models.Product, imageUrl string) (*ent.Product, error)
		Read(id int) (*ent.Product, error)
		ReadBySupplierMerchant(id int) (*ent.Product, error)
		ReadByRetailMerchant(id int) (*ent.Product, error)
		ReadAll(limit, offset int) ([]*ent.Product, error)
		ReadBySlugRetailMerchantCategoryMajor(slug string, limit, offset int) ([]*ent.ProductCategoryMajor, error)
		ReadBySlugRetailMerchantCategoryMinor(slug string, limit, offset int) ([]*ent.ProductCategoryMinor, error)
		ReadBySlugSupplierMerchantCategoryMajor(slug string, limit, offset int) ([]*ent.ProductCategoryMajor, error)
		ReadBySlugSupplierMerchantCategoryMinor(slug string, limit, offset int) ([]*ent.ProductCategoryMinor, error)
		ReadAllRetailMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error)
		ReadAllSupplierMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error)
		ReadAllBySupplierMerchant(merchant, limit, offset int) ([]*ent.Product, error)
		ReadAllByRetailMerchant(merchant, limit, offset int) ([]*ent.Product, error)
		ReadBestSellerBySupplierMerchant(limit, offset int) ([]*ent.Product, error)
		ReadBestSellerRetailMerchant(limit, offset int) ([]*ent.Product, error)

		ReadAllBySlugCategoryMajor(merchantType, slug string, limit, offset int) ([]*ent.Product, error)
		ReadAllBySlugCategoryMinor(merchantType, slug string, limit, offset int) ([]*ent.Product, error)
		ReadBestSellerByMerchant(id, limit, offset int) ([]*ent.Product, error)

		Update(id int, request *models.ProductUpdate) (*ent.Product, error)
		UpdateImage(id int, imagePath string) (string, error)

		Delete(id int) error
	}
	ProductCatMajorRepo interface {
		Insert(merchant *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		Read(id int) (*ent.ProductCategoryMajor, error)
		ReadAll() ([]*ent.ProductCategoryMajor, error)
		Update(id int, request *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		Delete(id int) error
	}
	ProductCatMinorRepo interface {
		Insert(merchant *models.ProductCategoryMinor, image string) (*ent.ProductCategoryMinor, error)
		Read(id int) (*ent.ProductCategoryMinor, error)
		ReadAll(limit, offset int) ([]*ent.ProductCategoryMinor, error)
		Update(id int, request *models.ProductCategoryMinorUpdate) (*ent.ProductCategoryMinor, error)
		UpdateImage(id int, imagePath string) (string, error)
		Delete(id int) error
	}
	AdminRepo interface {
		Insert(user *models.AdminUserRequest, password string) (*ent.Admin, error)
		Read(id int) (*ent.Admin, error)
		ReadAll(limit, offset int) (*presenters.ResponseWithTotalRecords, error)
		Update(id int, user *models.AdminUserRequest) (*ent.Admin, error)
		Delete(id int) error
	}
	RoleAndPermissionRepo interface {
		Insert(role *models.RoleRequest) (*ent.Role, error)
		ReadAll(limit, offset int) (*presenters.ResponseWithTotalRecords, error)
		ReadAllPermission() ([]*ent.Permission, error)
		Read(id int) (*ent.Role, error)
		Update(id int, role *models.RoleRequest) (*ent.Role, error)
		Delete(id int) error
	}
	AddressRepo interface {
		Insert(user *models.Address, userId int, userType string) (*ent.Address, error)
		Read(id int) (*ent.Address, error)
		ReadAll() ([]*ent.Address, error)
		ReadAllByUser(userId int, userType string) ([]*ent.Address, error)
		ReadByUser(userId int, userType string) (*ent.Address, error)
		Update(addressId int, address *models.Address) (*ent.Address, error)
		UpdateByUserDefaultAddress(userId, addressId int, userType string) ([]*ent.Address, error)
		SaveCoordinate(coordinate *services.Coordinate, id int) error
		Delete(id string) error
	}
	PickupStationRepo interface {
		Insert(station *models.PickupStation) (*ent.PickupStation, error)
		Read(id int) (*ent.PickupStation, error)
		ReadAll() ([]*ent.PickupStation, error)
		Update(stationId int, address *models.PickupStation) (*ent.PickupStation, error)
		Delete(id string) error
	}
	LogisticRepo interface {
		InsertResponse(
			orderNum string, storeId int, response *models.TookanPickupAndDeliveryTaskResponse,
		) (*ent.Logistic, error)
		UpdateResponse(response *models.TookanPickupAndDeliveryTaskResponse) (*ent.Logistic, error)
		UpdateOrderStatus(token, status string) error
		UpdateOrderDeliveryTask(orderNum string, storeId int) error
		Delete(id string) error
	}
	MapRepo interface {
		Delete(id string) error
		SaveCoordinate(coordinate *services.Coordinate, id int, model string) error
	}

	OrderRepo interface {
		Insert(order *models.OrderPayload) (*ent.Order, error)
		Read(id int) (*ent.Order, error)
		ReadByUser(userType string, id int) (*ent.Order, error)
		ReadAll() ([]*ent.Order, error)
		ReadAllByUser(userType string, id int) ([]*ent.Order, error)
		ReadAllByStore(merchantId int) ([]*ent.Order, error)
		ReadAllByAgentStore(agentId int) ([]*ent.Order, error)
		ReadByStore(id, merchantId int) (*ent.Order, error)
		ReadByAgentStore(id, agentId int) (*ent.Order, error)
		ReadOrderStoreMerchants(orderId int) (*ent.Order, error)
		Update(order *services.PaystackResponse) (*ent.Order, error)
		Delete(id string) error
		UpdateOrderDetailStatus(requests map[string]*gabs.Container) (*ent.Order, error)
		ReadByStoreOrderDetail(orderId int) ([]*ent.OrderDetail, error)
	}

	AuthRepo interface {
		ReadAdmin(username, field string) (*ent.Admin, error)
		ReadCustomer(username, field string) (*ent.Customer, error)
		ReadAgent(username, field string) (*ent.Agent, error)
		ReadMerchant(username, field string) (*ent.Merchant, error)
		UpdatePassword(id int, password string, userType string, isOTP bool) (bool, error)
		ResetPassword(id int, password, userType string) (bool, error)
	}

	PaymentRepo interface {
		Insert(transaction *services.Transaction) (*ent.Order, error)
	}
)
