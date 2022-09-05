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

func (r *repository) Insert(merchant *models.SupplierMerchant) (*ent.SupplierMerchant, error) {

	result, err := r.db.SupplierMerchant.Create().
		SetLastName(merchant.LastName).
		SetOtherName(merchant.OtherName).
		SetPhone(merchant.Phone).
		SetOtherPhone(merchant.OtherPhone).
		SetAddress(merchant.Address).
		SetDigitalAddress(merchant.DigitalAddress).
		SetGhanaCard(merchant.GhanaCard).
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

func (r *repository) Update(i *models.SupplierMerchant) (*models.SupplierMerchant, error) {
	// book.UpdatedAt = time.Now()
	// _, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": book.ID}, bson.M{"$set": book})
	// if err != nil {
	// 	return nil, err
	// }
	return i, nil
}

func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
