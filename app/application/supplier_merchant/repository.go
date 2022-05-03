package supplier_merchant

import (
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewSupplierMerchantRepo(db *database.Adapter) gateways.SupplierMerchantRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(customer *models.SupplierMerchant) (*ent.SupplierMerchant, error) {

	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte(customer.Password), 16)

	// r.db.Customer.Create().
	// 	SetFirstName(customer.FirstName).
	// 	SetLastName(customer.LastName).
	// 	SetPhone(customer.Phone).
	// 	SetEmail(customer.Email).
	// 	SetPassword(hashPassword).
	// 	Save(context.Background())

	// hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	// u, err := r.db.User.
	// 	Create().
	// 	SetUsername(user.Username).
	// 	SetPassword(hashPassword).
	// 	Save(ctx)

	// if err != nil {
	// 	return nil, fmt.Errorf("failed creating user: %w", err)
	// }

	return nil, nil
}

func (r *repository) Read(id int) (*ent.SupplierMerchant, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.SupplierMerchant, error) {

	// b, err := r.db.User.Query().
	// 	All(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (a *repository) Update(i *models.SupplierMerchant) (*models.SupplierMerchant, error) {
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
