package gateways

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent"
)

type (
	CustomerService interface {
		Create(customer any, customerType string) (*ent.Customer, error)
		FetchAll() ([]*ent.Customer, error)
		Fetch(id int) (*ent.Customer, error)
		Update(customer *models.IndividualCustomer) (*ent.Customer, error)
		UpdateLogo(customer int, logo string) (string, error)
		Remove(id string) error
	}
	AgentService interface {
		Create(agent *models.AgentRequest) (*ent.Agent, error)
		FetchAll() ([]*ent.Agent, error)
		FetchAllMerchant(agentId int) ([]*ent.MerchantStore, error)
		Fetch(id int) (*ent.Agent, error)
		Update(agent *models.Agent) (*models.Agent, error)
		Remove(id string) error
		CreateCompliance(
			request *models.AgentComplianceRequest, id int, report string, personal []string, guarantor []string,
		) (*ent.Agent, error)
	}
	MerchantService interface {
		Create(merchant *models.MerchantRequest) (*ent.Merchant, error)
		Onboard(merchant *models.OnboardMerchantFullRequest, agentId int, logo string, images []string) (
			*ent.Merchant, error,
		)
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
		Create(store *models.MerchantStoreRequest, merchantId int, logo string, images []string) (
			*ent.MerchantStore, error,
		)
		SaveAccount(store interface{}, storeId int, logo string) (*ent.MerchantStore, error)
		SaveDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error)
		SaveAgentPermission(request bool, storeId int) (*ent.MerchantStore, error)
		SaveLogo(c *fiber.Ctx, field, directory string) (interface{}, error)
		SavePhotos(c *fiber.Ctx, field, directory string) (interface{}, error)
		FetchAll() ([]*ent.MerchantStore, error)
		FetchAllByMerchant(merchantType string, limit, offset int) ([]*ent.MerchantStore, error)
		Fetch(id int) (*ent.MerchantStore, error)
		FetchAgent(store int) (*ent.Agent, error)
		FetchByMerchant(merchantId int) (*ent.MerchantStore, error)
		Update(store *models.MerchantStore, storeId int) (*ent.MerchantStore, error)
		UpdateAddress(address *models.MerchantStoreAddress, storeId int) (*ent.MerchantStore, error)
		Remove(id string) error
	}

	ProductService interface {
		Create(merchant *models.Product, imageUrl string) (*ent.Product, error)
		FetchAll(limit, offset int) ([]*ent.Product, error)
		FetchBySlugRetailMerchantCategoryMajor(slug string, limit, offset int) ([]*ent.ProductCategoryMajor, error)
		FetchBySlugRetailMerchantCategoryMinor(slug string, limit, offset int) ([]*ent.ProductCategoryMinor, error)
		FetchBySlugSupplierMerchantCategoryMajor(slug string, limit, offset int) ([]*ent.ProductCategoryMajor, error)
		FetchBySlugSupplierMerchantCategoryMinor(slug string, limit, offset int) ([]*ent.ProductCategoryMinor, error)
		FetchAllRetailMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error)
		FetchAllSupplierMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error)
		FetchAllBySupplier(supplier, limit, offset int) ([]*ent.Product, error)
		FetchAllByRetailer(retailer, limit, offset int) ([]*ent.Product, error)
		FetchBestSellerBySupplier(limit, offset int) ([]*ent.Product, error)
		FetchBestSellerByRetailer(limit, offset int) ([]*ent.Product, error)
		Fetch(id int) (*ent.Product, error)
		FetchBySupplierMerchant(id int) (*ent.Product, error)
		FetchByRetailMerchant(id int) (*ent.Product, error)

		FetchAllBySlugCategoryMajor(merchantType, slug string, limit, offset int) ([]*ent.Product, error)
		FetchAllBySlugCategoryMinor(merchantType, slug string, limit, offset int) ([]*ent.Product, error)
		FetchBestSellerByMerchant(id, limit, offset int) ([]*ent.Product, error)

		Update(merchant *models.Product) (*models.Product, error)
		Remove(id string) error
		SaveImage(c *fiber.Ctx, field, directory string) (map[string]string, error)
	}
	ProductCatMajorService interface {
		Create(merchant *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		FetchAll() ([]*ent.ProductCategoryMajor, error)
		Fetch(id int) (*ent.ProductCategoryMajor, error)
		Update(merchant *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error)
		Remove(id string) error
	}
	ProductCatMinorService interface {
		Create(merchant *models.ProductCategoryMinor, image string) (*ent.ProductCategoryMinor, error)
		SaveImage(c *fiber.Ctx, field, directory string) (map[string]string, error)
		FetchAll(limit, offset int) ([]*ent.ProductCategoryMinor, error)
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
		Create(order *models.OrderPayload) (*ent.Order, error)
		TesCreate(orderId int) (*ent.Order, error)
		FetchAll() ([]*ent.Order, error)
		FetchAllByUser(userType string, id int) ([]*ent.Order, error)
		FetchByStore(id, userId int, userType string) (*ent.Order, error)
		FetchAllByStore(merchantId int) ([]*ent.Order, error)
		FetchAllByAgentStore(agentId int) ([]*ent.Order, error)
		Fetch(id int) (*ent.Order, error)
		FetchByUser(userType string, id int) (*ent.Order, error)
		Update(order *services.PaystackResponse) (*ent.Order, error)
		Remove(id string) error
		UpdateOrderDetailStatus(request []byte) (*ent.Order, error)
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
		UpdatePassword(id string, request any, userType string, isOTP bool) (bool, error)
		ResetPassword(request *models.ResetPassword, username, userType string) (bool, error)
		SendUserVerificationCode(username string) (string, error)
		SendPasswordResetCode(username, userType string) (string, error)
	}

	PaymentService interface {
		Pay(request any) (any, error)
		Verify(reference string) (any, error)
		FormatPayload(request any) (*models.OrderPayload, error)
	}
	SMSService interface {
		Send(request *services.SMSPayload) (any, error)
	}

	EmailService interface {
		Listen()
		Send(msg *services.Message)
		Done()
		CloseChannels()
	}
	LogisticService interface {
		New(repo OrderRepo) LogisticService
		ExecuteTask(order *ent.Order, deliveryType string)
		ExecuteWebhook(response any)
		FareEstimate(coordinates *models.OrderFareEstimateRequest) ([]*services.FareEstimateResponseData, error)
		Listen()
		Done()
		CloseChannels()
	}
	MapService interface {
		SetRepo(repo MapRepo) MapService
		SetAddressRepo(repo AddressRepo) MapService
		SetMerchantRepo(repo MerchantRepo) MapService
		SetMerchantStoreRepo(repo MerchantStoreRepo) MapService
		ExecuteTask(data any, taskType, repoType string)
		Listen()
		Done()
		CloseChannels()
	}
	StorageService interface {
		UploadFile(dir string, f *multipart.FileHeader) (string, error)
		UploadFiles(dir string, f []*multipart.FileHeader) ([]*string, error)
		Disk(disk string) StorageService
		ExecuteTask(data any, taskType string)
		Listen()
		Done()
		CloseChannels()
	}
)
