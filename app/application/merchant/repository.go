package merchant

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

func NewMerchantRepo(db *database.Adapter) gateways.MerchantRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(mercthant *models.MerchantRequest) (*ent.Merchant, error) {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(mercthant.Credentials.Password), 16)

	merchant, err := r.db.Merchant.Create().
		SetType(mercthant.Info.MerchantType).
		SetUsername(mercthant.Credentials.Username).
		SetPassword(hashPassword).
		Save(context.Background())

	if err != nil {
		return nil, fmt.Errorf("failed creating merchant: %w", err)
	}

	if mercthant.Info.MerchantType == "supplier" {
		_, err := r.db.SupplierMerchant.Create().
			SetMerchant(merchant).
			SetGhanaCard(mercthant.Info.GhanaCard).
			SetLastName(mercthant.Info.LastName).
			SetOtherName(mercthant.Info.OtherName).
			SetPhone(mercthant.Info.Phone).
			SetOtherPhone(mercthant.Info.OtherPhone).
			SetAddress(mercthant.Info.Address).
			SetDigitalAddress(mercthant.Info.DigitalAddress).
			Save(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed creating merchant: %w", err)
		}
	} else if mercthant.Info.MerchantType == "retailer" {
		_, err := r.db.RetailMerchant.Create().
			SetMerchant(merchant).
			SetGhanaCard(mercthant.Info.GhanaCard).
			SetLastName(mercthant.Info.LastName).
			SetOtherName(mercthant.Info.OtherName).
			SetPhone(mercthant.Info.Phone).
			SetOtherPhone(mercthant.Info.OtherPhone).
			SetAddress(mercthant.Info.Address).
			SetDigitalAddress(mercthant.Info.DigitalAddress).
			Save(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed creating merchant: %w", err)
		}
	}
	return merchant, nil

}

func (r *repository) Read(id int) (*ent.Merchant, error) {

	// b, err := r.db.User.Query().Where(user.ID(id)).First(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (r *repository) ReadAll() ([]*ent.Merchant, error) {

	// b, err := r.db.User.Query().
	// 	All(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}

func (a *repository) Update(i *models.Merchant) (*models.Merchant, error) {
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
