package customer

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/customer"
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

func (r *repository) Update(i *models.IndividualCustomer) (*ent.Customer, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
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
