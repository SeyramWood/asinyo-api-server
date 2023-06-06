package merchant

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/retailmerchant"
	"github.com/SeyramWood/ent/suppliermerchant"
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
	_, err := r.db.MerchantStore.Create().SetMerchantID(m.ID).
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
	result, err := r.db.Merchant.Query().Where(merchant.ID(m.ID)).WithStore().Only(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r repository) SaveCoordinate(coordinate *services.Coordinate, id int) error {
	_, err := r.db.MerchantStore.UpdateOneID(id).SetCoordinate(coordinate).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Read(id int) (*ent.Merchant, error) {
	m, err := r.db.Merchant.Query().Where(merchant.ID(id)).WithSupplier().WithRetailer().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *repository) ReadStorefront(id int) (*ent.MerchantStore, error) {
	store, err := r.db.Merchant.Query().Where(merchant.ID(id)).
		QueryStore().
		WithAgent().
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (r *repository) ReadAll(limit, offset int) (*presenters.PaginationResponse, error) {
	ctx := context.Background()
	query := r.db.Merchant.Query()
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := query.
		Limit(limit).Offset(offset).
		Order(ent.Desc(merchant.FieldCreatedAt)).
		WithSupplier().
		WithRetailer().
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.PaginationResponse{
		Count: count,
		Data:  results,
	}, nil
}

func (r *repository) Update(id int, request any) (*ent.Merchant, error) {
	ctx := context.Background()
	if req, ok := request.(*models.RetailerProfileUpdate); ok {
		_, err := r.db.RetailMerchant.Update().Where(
			retailmerchant.HasMerchantWith(
				func(rm *sql.Selector) {
					rm.Where(sql.InInts(merchant.RetailerColumn, id))
				},
			),
		).
			SetLastName(req.LastName).
			SetOtherName(req.OtherName).
			SetPhone(req.Phone).
			SetOtherPhone(req.OtherPhone).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	if req, ok := request.(*models.SupplierProfileUpdate); ok {
		_, err := r.db.SupplierMerchant.Update().Where(
			suppliermerchant.HasMerchantWith(
				func(sm *sql.Selector) {
					sm.Where(sql.InInts(merchant.SupplierColumn, id))
				},
			),
		).
			SetLastName(req.LastName).
			SetOtherName(req.OtherName).
			SetPhone(req.Phone).
			SetOtherPhone(req.OtherPhone).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	result, err := r.db.Merchant.Query().Where(merchant.ID(id)).WithSupplier().WithRetailer().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Delete(ID string) error {
	return fmt.Errorf("failed creating book")
	// return r.Delete(ID).Error
}
