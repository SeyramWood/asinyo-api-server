package admin

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/admin"
)

type repository struct {
	db *ent.Client
}

func NewAdminRepo(db *database.Adapter) gateways.AdminRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(user *models.Admin) (*ent.Admin, error) {
	ctx := context.Background()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 16)
	u, err := r.db.Admin.
		Create().
		SetUsername(user.Username).
		SetPassword(hashPassword).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return u, nil
}

func (r *repository) Read(id int) (*ent.Admin, error) {

	b, err := r.db.Admin.Query().Where(admin.ID(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repository) ReadAll() ([]*ent.Admin, error) {

	b, err := r.db.Admin.Query().
		Order(ent.Desc(admin.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (a *repository) Update(i *models.Admin) (*models.Admin, error) {
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
