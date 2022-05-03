package auth

import (
	"context"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/user"
)

type repository struct {
	db *ent.Client
}

//NewRepo is the single instance repo that is being created.
func NewAuthRepo(db *database.Adapter) gateways.AuthRepo {
	return &repository{db.DB}
}

//ReadUser is a mongo repository that helps to fetch books
func (r *repository) Read(username string) (*ent.User, error) {
	b, err := r.db.User.Query().Where(user.Username(username)).First(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}
