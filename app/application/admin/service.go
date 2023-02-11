package admin

import (
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.AdminRepo
}

func NewAdminService(repo gateways.AdminRepo) gateways.AdminService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(user *models.AdminUserRequest) (*ent.Admin, error) {
	// TODO generate password and send email
	password, _ := application.GeneratePassword(12)
	fmt.Println(password)

	return s.repo.Insert(user, "password")
}

func (s *service) Fetch(id int) (*ent.Admin, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll(limit, offset int) (*presenters.ResponseWithTotalRecords, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) Update(id int, user *models.AdminUserRequest) (*ent.Admin, error) {
	return s.repo.Update(id, user)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}
