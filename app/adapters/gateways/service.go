package gateways

import (
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	CustomerService interface {
		Create(customer *models.Customer) (*ent.Customer, error)
		FetchAll() ([]*ent.Customer, error)
		Fetch(id int) (*ent.Customer, error)
		Update(customer *models.Customer) (*models.Customer, error)
		Remove(ID string) error
	}
	AgentService interface {
		Create(agent *models.AgentRequest) (*ent.Agent, error)
		FetchAll() ([]*ent.Agent, error)
		Fetch(id int) (*ent.Agent, error)
		Update(agent *models.Agent) (*models.Agent, error)
		Remove(ID string) error
	}
	SupplierMerchantService interface {
		Create(merchant *models.SupplierMerchant) (*ent.SupplierMerchant, error)
		FetchAll() ([]*ent.SupplierMerchant, error)
		Fetch(id int) (*ent.SupplierMerchant, error)
		Update(merchant *models.SupplierMerchant) (*models.SupplierMerchant, error)
		Remove(ID string) error
	}
	RetailMerchantService interface {
		Create(merchant *models.RetailMerchant) (*ent.RetailMerchant, error)
		FetchAll() ([]*ent.RetailMerchant, error)
		Fetch(id int) (*ent.RetailMerchant, error)
		Update(merchant *models.RetailMerchant) (*models.RetailMerchant, error)
		Remove(ID string) error
	}
	MerchantService interface {
		Create(merchant *models.MerchantRequest) (*ent.Merchant, error)
		FetchAll() ([]*ent.Merchant, error)
		Fetch(id int) (*ent.Merchant, error)
		Update(merchant *models.Merchant) (*models.Merchant, error)
		Remove(ID string) error
	}

	ProductService interface {
		Create(merchant *models.Product) (*ent.Product, error)
		FetchAll() ([]*ent.Product, error)
		Fetch(id int) (*ent.Product, error)
		Update(merchant *models.Product) (*models.Product, error)
		Remove(ID string) error
	}
	ProductCatMajorService interface {
		Create(merchant *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		FetchAll() ([]*ent.ProductCategoryMajor, error)
		Fetch(id int) (*ent.ProductCategoryMajor, error)
		Update(merchant *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error)
		Remove(ID string) error
	}
	ProductCatMinorService interface {
		Create(merchant *models.ProductCategoryMinor) (*ent.ProductCategoryMinor, error)
		FetchAll() ([]*ent.ProductCategoryMinor, error)
		Fetch(id int) (*ent.ProductCategoryMinor, error)
		Update(merchant *models.ProductCategoryMinor) (*models.ProductCategoryMinor, error)
		Remove(ID string) error
	}
	AdminService interface {
		Create(user *models.Admin) (*ent.Admin, error)
		FetchAll() ([]*ent.Admin, error)
		Fetch(id int) (*ent.Admin, error)
		Update(user *models.Admin) (*models.Admin, error)
		Remove(ID string) error
	}

	AuthService interface {
		Login(c *fiber.Ctx) error
		Logout(c *fiber.Ctx) error
		FetcAuthUser(c *fiber.Ctx) error
	}
)
