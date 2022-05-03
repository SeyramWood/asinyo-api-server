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
		Delete(ID string) error
	}
	AgentRepo interface {
		Insert(customer *models.Agent) (*ent.Agent, error)
		Read(id int) (*ent.Agent, error)
		ReadAll() ([]*ent.Agent, error)
		Update(customer *models.Agent) (*models.Agent, error)
		Delete(ID string) error
	}
	SupplierMerchantRepo interface {
		Insert(customer *models.SupplierMerchant) (*ent.SupplierMerchant, error)
		Read(id int) (*ent.SupplierMerchant, error)
		ReadAll() ([]*ent.SupplierMerchant, error)
		Update(customer *models.SupplierMerchant) (*models.SupplierMerchant, error)
		Delete(ID string) error
	}
	RetailMerchantRepo interface {
		Insert(customer *models.RetailMerchant) (*ent.RetailMerchant, error)
		Read(id int) (*ent.RetailMerchant, error)
		ReadAll() ([]*ent.RetailMerchant, error)
		Update(customer *models.RetailMerchant) (*models.RetailMerchant, error)
		Delete(ID string) error
	}

	AdminRepo interface {
		Insert(book *models.Admin) (*ent.Admin, error)
		Read(id int) (*ent.Admin, error)
		ReadAll() ([]*ent.Admin, error)
		Update(book *models.Admin) (*models.Admin, error)
		Delete(ID string) error
	}
	UserRepo interface {
		Insert(book *models.User) (*ent.User, error)
		Read(id int) (*ent.User, error)
		ReadAll() ([]*ent.User, error)
		Update(book *models.User) (*models.User, error)
		Delete(ID string) error
	}
	AuthRepo interface {
		Read(username string) (*ent.User, error)
	}
)
