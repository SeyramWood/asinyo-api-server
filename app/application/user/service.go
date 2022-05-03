package user

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.UserRepo
}

//NewService is used to create a single instance of the service
func NewUserService(repo gateways.UserRepo) gateways.UserService {
	return &service{
		repo: repo,
	}
}

//InsertBook is a service layer that helps insert book in BookShop
func (s *service) Create(user *models.User) (*ent.User, error) {

	return s.repo.Insert(user)
}

//FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) Fetch(id int) (*ent.User, error) {

	return s.repo.Read(id)
}
func (s *service) FetchAll() ([]*ent.User, error) {
	return s.repo.ReadAll()
}

//UpdateBook is a service layer that helps update books in BookShop
func (s *service) Update(user *models.User) (*models.User, error) {
	return s.repo.Update(user)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
