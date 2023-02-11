package role_permission

import (
	"context"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/role"
)

type repository struct {
	db *ent.Client
}

func NewRoleAndPermissionRepo(db *database.Adapter) gateways.RoleAndPermissionRepo {
	return &repository{
		db: db.DB,
	}
}
func (r *repository) Insert(role *models.RoleRequest) (*ent.Role, error) {
	result, err := r.db.Role.Create().
		AddPermissionIDs(role.Permissions...).
		SetRole(role.Role).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return r.Read(result.ID)
}

func (r *repository) ReadAll(limit, offset int) (*presenters.ResponseWithTotalRecords, error) {
	ctx := context.Background()
	roleQuery := r.db.Role.Query()
	totalRecords, err := roleQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	results, err := roleQuery.
		Order(ent.Desc(role.FieldCreatedAt)).
		Limit(limit).Offset(offset).
		WithPermissions().
		All(ctx)

	if err != nil {
		return nil, err
	}
	return &presenters.ResponseWithTotalRecords{
		TotalRecords: totalRecords,
		Records:      results,
	}, nil
}

func (r *repository) ReadAllPermission() ([]*ent.Permission, error) {
	results, err := r.db.Permission.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) Read(id int) (*ent.Role, error) {
	result, err := r.db.Role.Query().Where(role.ID(id)).WithPermissions().
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) Update(id int, role *models.RoleRequest) (*ent.Role, error) {
	_, err := r.db.Role.UpdateOneID(id).
		ClearPermissions().
		AddPermissionIDs(role.Permissions...).
		SetRole(role.Role).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return r.Read(id)
}

func (r *repository) Delete(id int) error {
	err := r.db.Role.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
