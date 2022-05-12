package retail_merchant

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewRetailMerchantRepo(db *database.Adapter) gateways.RetailMerchantRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(mercthant *models.RetailMerchant) (*ent.RetailMerchant, error) {
	result, err := r.db.RetailMerchant.Create().
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

func (r *repository) Read(id int) (*ent.RetailMerchant, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.RetailMerchant, error) {

	// b, err := r.db.User.Query().
	// 	All(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (a *repository) Update(i *models.RetailMerchant) (*models.RetailMerchant, error) {
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
