package gateways

import (
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	CustomerService interface {
		Create(user *models.Customer) (*ent.Customer, error)
		FetchAll() ([]*ent.Customer, error)
		Fetch(id int) (*ent.Customer, error)
		Update(user *models.Customer) (*models.Customer, error)
		Remove(ID string) error
	}
	AgentService interface {
		Create(user *models.Agent) (*ent.Agent, error)
		FetchAll() ([]*ent.Agent, error)
		Fetch(id int) (*ent.Agent, error)
		Update(user *models.Agent) (*models.Agent, error)
		Remove(ID string) error
	}
	SupplierMerchantService interface {
		Create(user *models.SupplierMerchant) (*ent.SupplierMerchant, error)
		FetchAll() ([]*ent.SupplierMerchant, error)
		Fetch(id int) (*ent.SupplierMerchant, error)
		Update(user *models.SupplierMerchant) (*models.SupplierMerchant, error)
		Remove(ID string) error
	}
	RetailMerchantService interface {
		Create(user *models.RetailMerchant) (*ent.RetailMerchant, error)
		FetchAll() ([]*ent.RetailMerchant, error)
		Fetch(id int) (*ent.RetailMerchant, error)
		Update(user *models.RetailMerchant) (*models.RetailMerchant, error)
		Remove(ID string) error
	}

	AdminService interface {
		Create(book *models.Admin) (*ent.Admin, error)
		FetchAll() ([]*ent.Admin, error)
		Fetch(id int) (*ent.Admin, error)
		Update(book *models.Admin) (*models.Admin, error)
		Remove(ID string) error
	}

	UserService interface {
		Create(user *models.User) (*ent.User, error)
		FetchAll() ([]*ent.User, error)
		Fetch(id int) (*ent.User, error)
		Update(user *models.User) (*models.User, error)
		Remove(ID string) error
	}
	AuthService interface {
		Login(c *fiber.Ctx) error
		Logout(c *fiber.Ctx) error
	}
)
