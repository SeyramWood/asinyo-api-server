package supplier_merchant

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/suppliermerchant"
)

type repository struct {
	db *ent.Client
}

func NewSupplierMerchantRepo(db *database.Adapter) gateways.SupplierMerchantRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(mercthant *models.SupplierMerchant) (*ent.SupplierMerchant, error) {

	result, err := r.db.SupplierMerchant.Create().
		SetLastName(mercthant.LastName).
		SetOtherName(mercthant.OtherName).
		SetPhone(mercthant.Phone).
		SetOtherPhone(mercthant.OtherPhone).
		SetAddress(mercthant.Address).
		SetDigitalAddress(mercthant.DigitalAddress).
		SetGhanaCard(mercthant.GhanaCard).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating agent: %w", err)
	}

	return result, nil
}

func (r *repository) Read(id int) (*ent.SupplierMerchant, error) {

	merchant, err := r.db.SupplierMerchant.Query().Where(suppliermerchant.ID(id)).WithMerchant().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return merchant, nil
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
