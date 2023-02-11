package admin

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
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

func (r *repository) Insert(user *models.AdminUserRequest, password string) (*ent.Admin, error) {
	ctx := context.Background()
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 16)
	u, err := r.db.Admin.
		Create().
		AddRoleIDs(user.Roles...).
		SetUsername(user.Username).
		SetPassword(hashPassword).
		SetLastName(user.LastName).
		SetOtherName(user.OtherName).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating admin user: %w", err)
	}
	return r.Read(u.ID)
}

func (r *repository) Read(id int) (*ent.Admin, error) {

	b, err := r.db.Admin.Query().Where(admin.ID(id)).WithRoles().Only(context.Background())
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repository) ReadAll(limit, offset int) (*presenters.ResponseWithTotalRecords, error) {
	ctx := context.Background()
	adminQuery := r.db.Admin.Query()
	totalRecords, err := adminQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := adminQuery.
		Order(ent.Desc(admin.FieldCreatedAt)).
		Limit(limit).Offset(offset).WithRoles().All(ctx)
	if err != nil {
		return nil, err
	}
	return &presenters.ResponseWithTotalRecords{
		TotalRecords: totalRecords,
		Records:      results,
	}, nil
}

func (r *repository) Update(id int, user *models.AdminUserRequest) (*ent.Admin, error) {
	_, err := r.db.Admin.UpdateOneID(id).
		ClearRoles().
		AddRoleIDs(user.Roles...).
		SetUsername(user.Username).
		SetLastName(user.LastName).
		SetOtherName(user.OtherName).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

func (r *repository) Delete(id int) error {
	err := r.db.Admin.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
