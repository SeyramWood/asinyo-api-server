package gateways

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent"
)

type (
	CustomerService interface {
		Create(customer any, customerType string) (*ent.Customer, error)
		CreatePurchaseRequest(customerId int, request *models.PurchaseOrderRequest) (*ent.PurchaseRequest, error)
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllPurchaseRequestByCustomer(customerId, limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Customer, error)
		FetchPurchaseRequest(id int) (*ent.PurchaseRequest, error)
		Update(id int, customer any) (*ent.Customer, error)
		UpdateLogo(customer int, logo string) (string, error)
		UpdatePurchaseRequest(id int, request *models.PurchaseOrderRequest) (*ent.PurchaseRequest, error)
		Remove(id int) error
		RemovePurchaseRequest(id int) error
	}
	AgentService interface {
		Create(agent *models.AgentRequest) (*ent.Agent, error)
		SaveAccount(account any, agentId int, accountType string) (*ent.Agent, error)
		SaveDefaultAccount(agentId int, accountType string) (*ent.Agent, error)
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllMerchant(agentId int) ([]*ent.MerchantStore, error)
		Fetch(id int) (*ent.Agent, error)
		Update(id int, profile *models.AgentProfile) (*ent.Agent, error)
		UpdateGuarantor(id int, request *models.AgentGuarantorUpdate) (*ent.Agent, error)
		UpdateAgentComplianceCard(agentId int, newPath, oldPath string) ([]string, error)
		UpdateAgentPoliceReport(agentId int, filePath string) (string, error)
		UpdateGuarantorComplianceCard(agentId int, newPath, oldPath string) ([]string, error)
		ApproveAgent(agentId int, complianceStatus bool) (*ent.Agent, error)
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
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		Fetch(id int) (*ent.Merchant, error)
		FetchStorefront(id int) (*ent.MerchantStore, error)
		Update(id int, request any) (*ent.Merchant, error)
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
		SaveAccount(store any, storeId int, accountType string) (*ent.MerchantStore, error)
		SaveDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error)
		SaveAgentPermission(request bool, storeId int) (*ent.MerchantStore, error)
		FetchAll() ([]*ent.MerchantStore, error)
		FetchAllByMerchant(merchantType string, limit, offset int) ([]*ent.MerchantStore, error)
		Fetch(id int) (*ent.MerchantStore, error)
		FetchAgent(store int) (*ent.Agent, error)
		FetchByMerchant(merchantId int) (*ent.MerchantStore, error)
		Update(request *models.MerchantStoreUpdate, storeId int) (*ent.MerchantStore, error)
		UpdateAddress(address *models.MerchantStoreAddress, storeId int) (*ent.MerchantStore, error)
		UpdateBanner(storeId int, bannerPath string) (string, error)
		UpdateImages(storeId int, newPath, oldPath string) ([]string, error)
		AppendNewImages(storeId int, urls []string) ([]string, error)
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

		Update(id int, request *models.ProductUpdate) (*ent.Product, error)
		UpdateImage(id int, imagePath string) (string, error)

		Remove(id int) error
	}
	ProductCatMajorService interface {
		Create(request *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		FetchAll() ([]*ent.ProductCategoryMajor, error)
		Fetch(id int) (*ent.ProductCategoryMajor, error)
		Update(id int, request *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		Remove(id int) error
	}
	ProductCatMinorService interface {
		Create(merchant *models.ProductCategoryMinor, image string) (*ent.ProductCategoryMinor, error)
		FetchAll(limit, offset int) ([]*ent.ProductCategoryMinor, error)
		Fetch(id int) (*ent.ProductCategoryMinor, error)
		Update(id int, request *models.ProductCategoryMinorUpdate) (*ent.ProductCategoryMinor, error)
		UpdateImage(id int, imagePath string) (string, error)
		Remove(id int) error
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
		Create(order *models.OrderPayload, params ...int) (*ent.Order, error)
		SaveOrderUpdate(id int, order *models.OrderPayload) (*ent.Order, error)
		TestCreate(orderId int) (*ent.Order, error)
		FetchAll() ([]*ent.Order, error)
		FetchAllByUser(userType string, id int) ([]*ent.Order, error)
		FetchByStore(id, userId int, userType string) (*ent.Order, error)
		FetchAllByStore(merchantId int) ([]*ent.Order, error)
		FetchAllByAgentStore(agentId int) ([]*ent.Order, error)
		Fetch(id int) (*ent.Order, error)
		FetchByUser(userType string, id int) (*ent.Order, error)
		Update(order *services.PaystackResponse) (*ent.Order, error)
		Remove(id string) error
		UpdateOrderDetailStatus(request []byte, logisticType string) (*ent.Order, error)
		UpdateOrderApprovalStatus(orderId int, status string) (*ent.Order, error)
	}
	AdminService interface {
		Create(user *models.AdminUserRequest) (*ent.Admin, error)
		OnboardNewCustomer(manager int, business *models.BusinessCustomerOnboardRequest) (*ent.Customer, error)
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllOrders(limit, offset int) (*presenters.PaginationResponse, error)
		FetchMyClients(manager, limit, offset int) (*presenters.PaginationResponse, error)
		FetchMyClientOrders(manager, limit, offset int) (*presenters.PaginationResponse, error)
		FetchMyClientsPurchaseRequest(manager int) ([]*ent.Customer, error)
		FetchProducts(major, minor string, limit, offset int) (*presenters.PaginationResponse, error)
		FetchAdminProducts(limit, offset int) (*presenters.PaginationResponse, error)
		FetchCounts(span string) (*presenters.DashboardRecordCount, error)
		FetchAccountManagers(perm ...string) ([]*ent.Admin, error)
		FetchConfigurations() ([]*ent.Configuration, error)
		FetchConfigurationByIdOrName(slug any) (*ent.Configuration, error)
		Fetch(id int) (*ent.Admin, error)
		Update(id int, user *models.AdminUserRequest) (*ent.Admin, error)
		AssignAccountManager(manager, customer int) (*ent.Customer, error)
		UpdateCurrentConfiguration(id int, configType, configValue string) (*ent.Configuration, error)
		Remove(id int) error
		RemoveOrder(id int) error
	}

	AuthService interface {
		Login(c *fiber.Ctx) error
		Logout(c *fiber.Ctx) error
		RefreshToken(c *fiber.Ctx) error
		FetchAuthUser(c *fiber.Ctx) error
		UpdatePassword(id string, request any, userType string, isOTP bool) (bool, error)
		ResetPassword(request *models.ResetPassword, username, userType string) (bool, error)
		SendUserVerificationCode(username string) (string, error)
		SendPasswordResetCode(username, userType string) (string, error)
		GenerateNewTokens(token string) (map[string]any, error)
	}
	RoleAndPermissionService interface {
		Create(role *models.RoleRequest) (*ent.Role, error)
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllPermission() ([]*ent.Permission, error)
		Fetch(id int) (*ent.Role, error)
		Update(id int, role *models.RoleRequest) (*ent.Role, error)
		Remove(id int) error
	}

	PaymentService interface {
		Pay(request any) (any, error)
		Verify(reference string) (any, error)
		FormatPayload(request any) (*models.OrderPayload, error)
	}
	SMSService interface {
		Listen()
		Send(msg *services.SMSPayload)
		Done()
		CloseChannels()
	}

	EmailService interface {
		Listen()
		Send(msg *services.MailerMessage)
		Done()
		CloseChannels()
	}

	DBNotificationService interface {
		Listen()
		Send(msg *services.DBNotificationMessage)
		Done()
		CloseChannels()
		Broker(action string, clientId string, params *services.NotificationFilters) *services.NotificationClient
		FetchNotifications(limit, offset int) ([]*ent.Notification, error)
		FetchUserNotifications(params *services.NotificationFilters) ([]*ent.Notification, error)
		MarkAsRead(userId, notificationId int, userType, timestamp string) (*ent.Notification, error)
		MarkSelectedAsRead(userId int, notificationIds []int, userType, timestamp string) error
		RemoveSelected(id []int) error
		Remove(id int) error
	}

	LogisticService interface {
		OrderRepo(repo OrderRepo) LogisticService
		ExecuteTask(order *ent.Order, task ...any)
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
		UploadFiles(dir string, files []*multipart.FileHeader) ([]string, error)
		Disk(disk string) StorageService
		ExecuteTask(data any, taskType string)
		Listen()
		Done()
		CloseChannels()
	}
	PriceModelService interface {
		Create(model *models.PriceModelRequest) (*ent.PriceModel, error)
		Fetch(id int) (*ent.PriceModel, error)
		FetchAll(limit, offset int) (*presenters.PaginationResponse, error)
		FetchAllPercentage(limit, offset int) (*presenters.PaginationResponse, error)
		Update(id int, model *models.PriceModelRequest) (*ent.PriceModel, error)
		UpdatePercentage(category, percentage int) (*ent.ProductCategoryMinor, error)
		Remove(id int) error
		RemovePercentage(id int) error
	}
)
