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
		Insert(agent *models.AgentRequest) (*ent.Agent, error)
		Read(id int) (*ent.Agent, error)
		ReadAll() ([]*ent.Agent, error)
		Update(agent *models.Agent) (*models.Agent, error)
		Delete(ID string) error
	}
	SupplierMerchantRepo interface {
		Insert(merchant *models.SupplierMerchant) (*ent.SupplierMerchant, error)
		Read(id int) (*ent.SupplierMerchant, error)
		ReadAll() ([]*ent.SupplierMerchant, error)
		Update(merchant *models.SupplierMerchant) (*models.SupplierMerchant, error)
		Delete(ID string) error
	}
	RetailMerchantRepo interface {
		Insert(merchant *models.RetailMerchant) (*ent.RetailMerchant, error)
		Read(id int) (*ent.RetailMerchant, error)
		ReadAll() ([]*ent.RetailMerchant, error)
		Update(merchant *models.RetailMerchant) (*models.RetailMerchant, error)
		Delete(ID string) error
	}
	MerchantRepo interface {
		Insert(merchant *models.MerchantRequest) (*ent.Merchant, error)
		Read(id int) (*ent.Merchant, error)
		ReadAll() ([]*ent.Merchant, error)
		Update(merchant *models.Merchant) (*models.Merchant, error)
		Delete(ID string) error
	}
	ProductRepo interface {
		Insert(merchant *models.Product) (*ent.Product, error)
		Read(id int) (*ent.Product, error)
		ReadAll() ([]*ent.Product, error)
		Update(merchant *models.Product) (*models.Product, error)
		Delete(ID string) error
	}
	ProductCatMajorRepo interface {
		Insert(merchant *models.ProductCategoryMajor) (*ent.ProductCategoryMajor, error)
		Read(id int) (*ent.ProductCategoryMajor, error)
		ReadAll() ([]*ent.ProductCategoryMajor, error)
		Update(merchant *models.ProductCategoryMajor) (*models.ProductCategoryMajor, error)
		Delete(ID string) error
	}
	ProductCatMinorRepo interface {
		Insert(merchant *models.ProductCategoryMinor) (*ent.ProductCategoryMinor, error)
		Read(id int) (*ent.ProductCategoryMinor, error)
		ReadAll() ([]*ent.ProductCategoryMinor, error)
		Update(merchant *models.ProductCategoryMinor) (*models.ProductCategoryMinor, error)
		Delete(ID string) error
	}
	AdminRepo interface {
		Insert(user *models.Admin) (*ent.Admin, error)
		Read(id int) (*ent.Admin, error)
		ReadAll() ([]*ent.Admin, error)
		Update(user *models.Admin) (*models.Admin, error)
		Delete(ID string) error
	}

	AuthRepo interface {
		ReadAdmin(username, field string) (*ent.Admin, error)
		ReadCustomer(username, field string) (*ent.Customer, error)
		ReadAgent(username, field string) (*ent.Agent, error)
		ReadMerchant(username, field string) (*ent.Merchant, error)
	}
)
