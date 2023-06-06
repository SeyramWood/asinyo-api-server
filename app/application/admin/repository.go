package admin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/admin"
	"github.com/SeyramWood/ent/configuration"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/permission"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategorymajor"
	"github.com/SeyramWood/ent/productcategoryminor"
	"github.com/SeyramWood/ent/role"
)

type repository struct {
	db *ent.Client
}

func NewAdminRepo(db *database.Adapter) gateways.AdminRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(user *models.AdminUserRequest, password string) (*ent.Admin, error) {
	ctx := context.Background()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 16)
	u, err := r.db.Admin.
		Create().
		AddRoleIDs(user.Roles...).
		SetUsername(user.Username).
		SetPassword(hashPassword).
		SetLastName(user.LastName).
		SetOtherName(user.OtherName).
		SetPhone(user.Phone).
		SetOtherPhone(user.OtherPhone).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating admin user: %w", err)
	}
	return r.Read(u.ID)
}
func (r *repository) OnboardNewCustomer(manager int, password string, business *models.BusinessCustomerOnboardRequest) (
	*ent.Customer, error,
) {
	ctx := context.Background()
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a transaction: %w", err)
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 16)
	c, err := tx.Customer.Create().
		SetAdminID(manager).
		SetUsername(business.Detail.Username).
		SetPassword(hashPassword).
		SetType("business").
		Save(ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating customer: %w", err))
	}
	_, err = tx.BusinessCustomer.Create().
		SetCustomer(c).
		SetName(business.Detail.BusinessName).
		SetPhone(business.Detail.BusinessPhone).
		SetOtherPhone(business.Detail.OtherPhone).
		SetContact(
			&models.BusinessCustomerContact{
				Name:     business.Contact.Name,
				Position: business.Contact.Position,
				Phone:    business.Contact.Phone,
				Email:    business.Contact.Email,
			},
		).Save(ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating business customer: %w", err))
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed commiting customer transaction: %w", err)
	}
	result, err := r.db.Customer.Query().Where(customer.ID(c.Unwrap().ID)).WithBusiness().Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Read(id int) (*ent.Admin, error) {
	b, err := r.db.Admin.Query().Where(admin.ID(id)).WithRoles().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repository) ReadAll(limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	adminQuery := r.db.Admin.Query()
	totalRecords, err := adminQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := adminQuery.
		Order(ent.Desc(admin.FieldCreatedAt)).
		Limit(limit).Offset(offset).WithRoles().All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  results,
	}, nil
}
func (r *repository) ReadMyClients(manager, limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	customerQuery := r.db.Admin.Query().Where(admin.ID(manager)).QueryCustomers()
	totalRecords, err := customerQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := customerQuery.
		Order(ent.Desc(customer.FieldCreatedAt)).
		Limit(limit).Offset(offset).
		WithBusiness().
		WithIndividual().
		WithPurchaseRequest().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  results,
	}, nil
}
func (r *repository) ReadMyClientsPurchaseRequest(manager int) ([]*ent.Customer, error) {
	results, err := r.db.Admin.Query().Where(admin.ID(manager)).QueryCustomers().Where(customer.HasPurchaseRequest()).
		Order(ent.Desc(customer.FieldCreatedAt)).
		WithBusiness().
		WithIndividual().
		WithPurchaseRequest().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *repository) ReadMyClientOrders(manager, limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	customerQuery := r.db.Admin.Query().Where(admin.ID(manager)).QueryCustomers().QueryOrders()
	totalRecords, err := customerQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := customerQuery.
		Order(ent.Desc(customer.FieldCreatedAt)).
		Limit(limit).Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  results,
	}, nil

}
func (r *repository) ReadAllOrders(limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	orderQuery := r.db.Order.Query()
	totalRecords, err := orderQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := orderQuery.
		Order(ent.Desc(customer.FieldCreatedAt)).
		Limit(limit).Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  results,
	}, nil
}
func (r *repository) ReadProducts(major, minor string, limit, offset int) (*presenters.PaginationResponse, error) {

	ctx := context.Background()
	productQuery := r.readProductsQuery("supplier", major, minor)
	totalRecords, err := productQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	products, err := productQuery.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		WithMajor().
		WithMinor().
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithRetailer()
				mq.WithStore(
					func(query *ent.MerchantStoreQuery) {
						query.Select("id", "name", "coordinate")
					},
				)
			},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  products,
	}, nil
}
func (r *repository) ReadAdminProducts(limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	productQuery := r.db.Product.Query()
	totalRecords, err := productQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	products, err := productQuery.
		Limit(limit).
		Offset(offset).
		Order(ent.Desc(product.FieldCreatedAt)).
		WithMajor().
		WithMinor().
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.WithSupplier()
				mq.WithRetailer()
				mq.WithStore()
			},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: totalRecords,
		Data:  products,
	}, nil
}
func (r *repository) ReadAllByPermissions(permissions []string) ([]*ent.Admin, error) {
	results, err := r.db.Admin.Query().
		Order(ent.Desc(admin.FieldCreatedAt)).
		Where(
			admin.HasRolesWith(role.HasPermissionsWith(permission.SlugIn(permissions...))),
		).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) ReadCounts(span string) (*presenters.DashboardRecordCount, error) {
	ctx := context.Background()
	// cQuery := r.db.Customer.Query()

	spanQuery := fmt.Sprintf("created_at > now() - interval %s day", span)
	return &presenters.DashboardRecordCount{
		Customers: &presenters.CustomerRecordCount{
			Business: &presenters.RecordCount{
				Total: r.db.Customer.Query().Where(customer.TypeIn("business")).CountX(ctx),
				Recent: r.db.Customer.Query().Where(customer.TypeIn("business")).Where(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(spanQuery))
					},
				).CountX(ctx),
			},
			Individual: &presenters.RecordCount{
				Total: r.db.Customer.Query().Where(customer.TypeIn("individual")).CountX(ctx),
				Recent: r.db.Customer.Query().Where(customer.TypeIn("individual")).Where(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(spanQuery))
					},
				).CountX(ctx),
			},
		},
		Merchants: &presenters.MerchantRecordCount{
			Supplier: &presenters.RecordCount{
				Total: r.db.Merchant.Query().Where(merchant.TypeIn("supplier")).CountX(ctx),
				Recent: r.db.Merchant.Query().Where(merchant.TypeIn("supplier")).Where(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(spanQuery))
					},
				).CountX(ctx),
			},
			Retailer: &presenters.RecordCount{
				Total: r.db.Merchant.Query().Where(merchant.TypeIn("retailer")).CountX(ctx),
				Recent: r.db.Merchant.Query().Where(merchant.TypeIn("retailer")).Where(
					func(s *sql.Selector) {
						s.Where(sql.ExprP(spanQuery))
					},
				).CountX(ctx),
			},
		},
		Agents: &presenters.RecordCount{
			Total: r.db.Agent.Query().CountX(ctx),
			Recent: r.db.Agent.Query().Where(
				func(s *sql.Selector) {
					s.Where(sql.ExprP(spanQuery))
				},
			).CountX(ctx),
		},
		Orders: &presenters.RecordCount{
			Total: r.db.Order.Query().CountX(ctx),
			Recent: r.db.Order.Query().Where(
				func(s *sql.Selector) {
					s.Where(sql.ExprP(spanQuery))
				},
			).CountX(ctx),
		},
	}, nil
}

func (r *repository) ReadConfigurations() ([]*ent.Configuration, error) {
	results, err := r.db.Configuration.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) ReadConfigurationByIdOrName(slug any) (*ent.Configuration, error) {
	var result *ent.Configuration
	if id, ok := slug.(int); ok {
		res, err := r.db.Configuration.Query().Where(configuration.ID(id)).Only(context.Background())
		if err != nil {
			return nil, err
		}
		result = res
	} else if name, ok := slug.(string); ok {
		res, err := r.db.Configuration.Query().Where(configuration.Name(name)).Only(context.Background())
		if err != nil {
			return nil, err
		}
		result = res
	}
	return result, nil
}

func (r *repository) Update(id int, user *models.AdminUserRequest) (*ent.Admin, error) {
	_, err := r.db.Admin.UpdateOneID(id).
		ClearRoles().
		AddRoleIDs(user.Roles...).
		SetUsername(user.Username).
		SetLastName(user.LastName).
		SetOtherName(user.OtherName).
		SetPhone(user.Phone).
		SetOtherPhone(user.OtherPhone).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

func (r *repository) AssignAccountManager(manager, customer int) (*ent.Customer, error) {
	result, err := r.db.Customer.UpdateOneID(customer).SetAdminID(manager).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateCurrentConfiguration(id int, configType, configValue string) (*ent.Configuration, error) {
	switch configType {
	case "logistic":
		return r.updateCurrentLogisticConfiguration(id, configValue)
	case "pricing-model":
		return r.updateCurrentPricingModelConfiguration(id, configValue)
	default:
		return nil, nil
	}
}

func (r *repository) Delete(id int) error {
	err := r.db.Admin.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteOrder(id int) error {
	err := r.db.Order.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) updateCurrentLogisticConfiguration(id int, configValue string) (*ent.Configuration, error) {
	conf := r.db.Configuration.Query().Where(configuration.ID(id)).OnlyX(context.Background())
	oldConf := conf.Data.Data.(map[string]any)
	_, err := conf.Update().
		SetData(
			&struct {
				Data any `json:"data"`
			}{
				Data: &models.LogisticConfiguration{
					Logistics: append([]any{}, oldConf["logistics"].([]any)...),
					Current:   configValue,
				},
			},
		).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *repository) updateCurrentPricingModelConfiguration(id int, configValue string) (*ent.Configuration, error) {
	conf := r.db.Configuration.Query().Where(configuration.ID(id)).OnlyX(context.Background())
	oldConf := conf.Data.Data.(map[string]any)
	_, err := conf.Update().
		SetData(
			&struct {
				Data any `json:"data"`
			}{
				Data: &models.PricingModelConfiguration{
					Models:  append([]any{}, oldConf["models"].([]any)...),
					Current: configValue,
				},
			},
		).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *repository) readProductsQuery(merchantType, major, minor string) *ent.ProductQuery {
	productQuery := r.db.Product.Query()
	if major != "all" && minor != "all" {
		productQuery.Where(
			product.HasMerchantWith(merchant.Type(merchantType)),
			product.And(
				product.HasMajorWith(productcategorymajor.Slug(major)),
				product.HasMinorWith(productcategoryminor.Slug(minor)),
			),
		)
		return productQuery
	}
	if major == "all" && minor != "all" {
		productQuery.Where(
			product.HasMerchantWith(merchant.Type(merchantType)),
			product.And(
				product.HasMinorWith(productcategoryminor.Slug(minor)),
			),
		)
		return productQuery
	}
	if major != "all" && minor == "all" {
		productQuery.Where(
			product.HasMerchantWith(merchant.Type(merchantType)),
			product.And(
				product.HasMajorWith(productcategorymajor.Slug(major)),
			),
		)
		return productQuery
	}

	productQuery.Where(
		product.HasMerchantWith(merchant.Type(merchantType)),
	)
	return productQuery
}
