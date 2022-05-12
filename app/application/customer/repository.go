package customer

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"golang.org/x/crypto/bcrypt"
)

type repository struct {
	db *ent.Client
}

func NewCustomerRepo(db *database.Adapter) gateways.CustomerRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(customer *models.Customer) (*ent.Customer, error) {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 16)

	c, err := r.db.Customer.Create().
		SetFirstName(customer.FirstName).
		SetLastName(customer.LastName).
		SetPhone(customer.Phone).
		SetUsername(customer.Username).
		SetPassword(hashPassword).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return c, nil
}

func (r *repository) Read(id int) (*ent.Customer, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.Customer, error) {

	results, err := r.db.Customer.
		Query().
		All(context.Background())

	if err != nil {
		return nil, err
	}
	return results, nil
}

func (a *repository) Update(i *models.Customer) (*models.Customer, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

//DeleteBook is a mongo repository that helps to delete books
func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
