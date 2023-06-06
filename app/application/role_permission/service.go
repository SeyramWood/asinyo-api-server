package role_permission

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.RoleAndPermissionRepo
}

func NewRoleAndPermissionService(repo gateways.RoleAndPermissionRepo) gateways.RoleAndPermissionService {
	return &service{
		repo: repo,
	}
}
func (s *service) Create(role *models.RoleRequest) (*ent.Role, error) {
	return s.repo.Insert(role)
}

func (s *service) FetchAll(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) FetchAllPermission() ([]*ent.Permission, error) {
	return s.repo.ReadAllPermission()
}

func (s *service) Fetch(id int) (*ent.Role, error) {
	return s.repo.Read(id)
}

func (s *service) Update(id int, role *models.RoleRequest) (*ent.Role, error) {
	return s.repo.Update(id, role)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}
