package address

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.AddressRepo
}

func NewAddressService(repo gateways.AddressRepo) gateways.AddressService {
	return &service{
		repo: repo,
	}
}

func (s service) Create(address *models.Address, userId int, userType string) (*ent.Address, error) {
	return s.repo.Insert(address, userId, userType)
}

func (s service) FetchAll() ([]*ent.Address, error) {
	//TODO implement me
	panic("implement me")
}
func (s service) FetchAllByUser(userId int, userType string) ([]*ent.Address, error) {
	return s.repo.ReadAllByUser(userId, userType)
}
func (s service) FetchByUser(userId int, userType string) (*ent.Address, error) {
	return s.repo.ReadByUser(userId, userType)
}
func (s service) Fetch(id int) (*ent.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(addressId int, address *models.Address) (*ent.Address, error) {
	return s.repo.Update(addressId, address)
}

func (s service) SaveDefaultAddress(userId, addressId int, userType string) ([]*ent.Address, error) {
	return s.repo.UpdateByUserDefaultAddress(userId, addressId, userType)
}

func (s service) Remove(id string) error {
	//TODO implement me
	panic("implement me")
}
