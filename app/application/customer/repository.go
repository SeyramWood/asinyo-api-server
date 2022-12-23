package customer

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/businesscustomer"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/individualcustomer"
)

type repository struct {
	db *ent.Client
}

func NewCustomerRepo(db *database.Adapter) gateways.CustomerRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(customer any, customerType string) (*ent.Customer, error) {
	ctx := context.Background()
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	switch customerType {
	case "individual":
		request := customer.(*models.IndividualCustomer)
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 16)
		c, err := tx.Customer.Create().
			SetUsername(request.Username).
			SetPassword(hashPassword).
			SetType(customerType).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating customer: %w", err))
		}
		_, err = tx.IndividualCustomer.Create().
			SetCustomer(c).
			SetLastName(request.LastName).
			SetOtherName(request.OtherName).
			SetPhone(request.Phone).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating individual customer: %w", err))
		}
		if err = tx.Commit(); err != nil {
			return nil, fmt.Errorf("failed commiting customer transaction: %w", err)
		}
		return c.Unwrap(), nil
	default:
		request := customer.(*models.BusinessCustomer)
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 16)
		c, err := tx.Customer.Create().
			SetUsername(request.Username).
			SetPassword(hashPassword).
			SetType(customerType).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating customer: %w", err))
		}
		_, err = tx.BusinessCustomer.Create().
			SetCustomer(c).
			SetName(request.BusinessName).
			SetPhone(request.Phone).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating business customer: %w", err))
		}
		if err = tx.Commit(); err != nil {
			return nil, fmt.Errorf("failed commiting customer transaction: %w", err)
		}
		return c.Unwrap(), nil

	}

}

func (r *repository) Read(id int) (*ent.Customer, error) {

	c, err := r.db.Customer.Query().Where(customer.ID(id)).WithBusiness().WithIndividual().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *repository) ReadAll() ([]*ent.Customer, error) {

	results, err := r.db.Customer.Query().
		Order(ent.Desc(customer.FieldCreatedAt)).
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) Update(id int, c any) (*ent.Customer, error) {
	ctx := context.Background()
	if request, ok := c.(*models.IndividualCustomerUpdate); ok {
		_, err := r.db.IndividualCustomer.Update().Where(
			individualcustomer.HasCustomerWith(
				func(bc *sql.Selector) {
					bc.Where(sql.InInts(customer.IndividualColumn, id))
				},
			),
		).
			SetLastName(request.LastName).
			SetOtherName(request.OtherName).
			SetPhone(request.Phone).
			SetOtherPhone(request.OtherPhone).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	if request, ok := c.(*models.BusinessCustomerUpdate); ok {
		_, err := r.db.BusinessCustomer.Update().Where(
			businesscustomer.HasCustomerWith(
				func(bc *sql.Selector) {
					bc.Where(sql.InInts(customer.BusinessColumn, id))
				},
			),
		).
			SetName(request.BusinessName).
			SetPhone(request.BusinessPhone).
			SetOtherPhone(request.OtherPhone).
			SetContact(
				&models.BusinessCustomerContact{
					Name:     request.ContactName,
					Position: request.ContactPosition,
					Phone:    request.ContactPhone,
					Email:    request.ContactEmail,
				},
			).Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	result, err := r.db.Customer.Query().Where(customer.ID(id)).WithBusiness().WithIndividual().Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) UpdateLogo(c int, logo string) (string, error) {
	ctx := context.Background()
	_, err := r.db.BusinessCustomer.UpdateOneID(r.db.Customer.Query().Where(customer.ID(c)).QueryBusiness().OnlyIDX(ctx)).SetLogo(logo).Save(ctx)
	if err != nil {
		return "", err
	}
	return logo, nil
}

func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
