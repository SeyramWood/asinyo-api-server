package merchant

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent/merchant"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type repository struct {
	db *ent.Client
}

func NewMerchantRepo(db *database.Adapter) gateways.MerchantRepo {
	return &repository{
		db: db.DB,
	}
}

func (r *repository) Insert(mc *models.MerchantRequest, onboard bool) (*ent.Merchant, error) {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(mc.Credentials.Password), 16)

	ctx := context.Background()
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}

	if onboard {
		merc, err := tx.Merchant.Create().
			SetType(mc.Info.MerchantType).
			SetUsername(mc.Credentials.Username).
			SetPassword(hashPassword).
			SetOtp(true).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating merchant: %w", err))
		}
		if mc.Info.MerchantType == "supplier" {
			_, err := tx.SupplierMerchant.Create().
				SetMerchant(merc).
				SetGhanaCard(mc.Info.GhanaCard).
				SetLastName(mc.Info.LastName).
				SetOtherName(mc.Info.OtherName).
				SetPhone(mc.Info.Phone).
				SetOtherPhone(mc.Info.OtherPhone).
				Save(ctx)
			if err != nil {
				return nil, application.Rollback(tx, fmt.Errorf("failed creating supplier merchant: %w", err))
			}
		} else if mc.Info.MerchantType == "retailer" {
			_, err := tx.RetailMerchant.Create().
				SetMerchant(merc).
				SetGhanaCard(mc.Info.GhanaCard).
				SetLastName(mc.Info.LastName).
				SetOtherName(mc.Info.OtherName).
				SetPhone(mc.Info.Phone).
				SetOtherPhone(mc.Info.OtherPhone).
				Save(ctx)
			if err != nil {
				return nil, application.Rollback(tx, fmt.Errorf("failed creating retail merchant: %w", err))
			}
		}
		if err = tx.Commit(); err != nil {
			return nil, fmt.Errorf("failed commiting merchant transaction: %w", err)
		}
		return merc.Unwrap(), nil

	}

	merc, err := tx.Merchant.Create().
		SetType(mc.Info.MerchantType).
		SetUsername(mc.Credentials.Username).
		SetPassword(hashPassword).
		Save(ctx)

	if err != nil {
		return nil, application.Rollback(tx, fmt.Errorf("failed creating merchant: %w", err))
	}

	if mc.Info.MerchantType == "supplier" {
		_, err := tx.SupplierMerchant.Create().
			SetMerchant(merc).
			SetGhanaCard(mc.Info.GhanaCard).
			SetLastName(mc.Info.LastName).
			SetOtherName(mc.Info.OtherName).
			SetPhone(mc.Info.Phone).
			SetOtherPhone(mc.Info.OtherPhone).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating supplier merchant: %w", err))
		}
	} else if mc.Info.MerchantType == "retailer" {
		_, err := tx.RetailMerchant.Create().
			SetMerchant(merc).
			SetGhanaCard(mc.Info.GhanaCard).
			SetLastName(mc.Info.LastName).
			SetOtherName(mc.Info.OtherName).
			SetPhone(mc.Info.Phone).
			SetOtherPhone(mc.Info.OtherPhone).
			Save(ctx)
		if err != nil {
			return nil, application.Rollback(tx, fmt.Errorf("failed creating retail merchant: %w", err))
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed commiting merchant transaction: %w", err)
	}
	return merc.Unwrap(), nil

}

func (r *repository) Onboard(
	merc *models.OnboardMerchantFullRequest, agentId int, logo string, images []string, password string,
) (
	*ent.Merchant, error,
) {

	mr := &models.MerchantRequest{
		Info: models.RetailMerchantRequestInfo{
			MerchantType: merc.PersonalInfo.MerchantType,
			GhanaCard:    merc.PersonalInfo.GhanaCard,
			LastName:     merc.PersonalInfo.LastName,
			OtherName:    merc.PersonalInfo.OtherName,
			Phone:        merc.PersonalInfo.Phone,
			OtherPhone:   merc.PersonalInfo.OtherPhone,
		},
		Credentials: models.MerchantRequestCredentials{
			Username:        merc.PersonalInfo.Username,
			Password:        password,
			ConfirmPassword: password,
			Terms:           true,
		},
	}
	m, merr := r.Insert(mr, true)
	if merr != nil {
		return nil, merr
	}
	ctx := context.Background()
	mq := r.db.Merchant.Query().Where(merchant.ID(m.ID)).OnlyX(ctx)
	_, err := r.db.MerchantStore.Create().SetMerchant(mq).
		SetAgentID(agentId).
		SetName(merc.StoreInfo.BusinessName).
		SetSlogan(merc.StoreInfo.BusinessSlogan).
		SetAbout(merc.StoreInfo.About).
		SetDescription(merc.StoreInfo.Description).
		SetLogo(logo).
		SetImages(images).
		SetMerchantType(merc.PersonalInfo.MerchantType).
		SetAddress(merc.Address).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating merchant store: %w", err)
	}
	return m, nil
}
func (r repository) SaveCoordinate(coordinate *services.Coordinate, id int) error {
	_, err := r.db.MerchantStore.UpdateOneID(id).SetCoordinate(coordinate).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
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

func (r *repository) Update(i *models.Merchant) (*models.Merchant, error) {
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
