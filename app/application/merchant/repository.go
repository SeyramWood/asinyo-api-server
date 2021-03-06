package merchant

import (
	"context"
	"fmt"
	"github.com/SeyramWood/app/application"

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

func (r *repository) Insert(mc *models.MerchantRequest) (*ent.Merchant, error) {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(mc.Credentials.Password), 16)

	ctx := context.Background()
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	merchant, err := tx.Merchant.Create().
		SetType(mc.Info.MerchantType).
		SetUsername(mc.Credentials.Username).
		SetPassword(hashPassword).
		Save(ctx)
	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating merchant: %w", err))
	}
	if mc.Info.MerchantType == "supplier" {
		_, err := tx.SupplierMerchant.Create().
			SetMerchant(merchant).
			SetGhanaCard(mc.Info.GhanaCard).
			SetLastName(mc.Info.LastName).
			SetOtherName(mc.Info.OtherName).
			SetPhone(mc.Info.Phone).
			SetOtherPhone(mc.Info.OtherPhone).
			SetAddress(mc.Info.Address).
			SetDigitalAddress(mc.Info.DigitalAddress).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating supplier merchant: %w", err))
		}
	} else if mc.Info.MerchantType == "retailer" {
		_, err := tx.RetailMerchant.Create().
			SetMerchant(merchant).
			SetGhanaCard(mc.Info.GhanaCard).
			SetLastName(mc.Info.LastName).
			SetOtherName(mc.Info.OtherName).
			SetPhone(mc.Info.Phone).
			SetOtherPhone(mc.Info.OtherPhone).
			SetAddress(mc.Info.Address).
			SetDigitalAddress(mc.Info.DigitalAddress).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating retail merchant: %w", err))
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed commiting merchant transaction: %w", err)
	}
	return merchant.Unwrap(), nil

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
