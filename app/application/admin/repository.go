package admin

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/admin"
	"golang.org/x/crypto/bcrypt"
)

type repository struct {
	db *ent.Client
}

//NewRepo is the single instance repo that is being created.
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

//ReadUser is a mongo repository that helps to fetch books
func (r *repository) Read(id int) (*ent.Admin, error) {

	b, err := r.db.Admin.Query().Where(admin.ID(id)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

//ReadUser is a mongo repository that helps to fetch books
func (r *repository) ReadAll() ([]*ent.Admin, error) {

	b, err := r.db.Admin.Query().
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

//UpdateBook is a mongo repository that helps to update books
func (a *repository) Update(i *models.Admin) (*models.Admin, error) {
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
