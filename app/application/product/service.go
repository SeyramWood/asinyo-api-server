package product

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.ProductRepo
}

func NewProductService(repo gateways.ProductRepo) gateways.ProductService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(merchant *models.Product, imageUrl string) (*ent.Product, error) {

	return s.repo.Insert(merchant, imageUrl)
}

func (s *service) Fetch(id int) (*ent.Product, error) {

	return s.repo.Read(id)
}

func (s *service) FetchBySupplierMerchant(id int) (*ent.Product, error) {
	return s.repo.ReadBySupplierMerchant(id)
}
func (s *service) FetchByRetailMerchant(id int) (*ent.Product, error) {
	return s.repo.ReadByRetailMerchant(id)
}

func (s *service) FetchAll() ([]*ent.Product, error) {
	return s.repo.ReadAll()
}

func (s *service) FetchBySlugRetailMerchantCategoryMajor(slug string) ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadBySlugRetailMerchantCategoryMajor(slug)
}

func (s *service) FetchBySlugRetailMerchantCategoryMinor(slug string) ([]*ent.ProductCategoryMinor, error) {
	return s.repo.ReadBySlugRetailMerchantCategoryMinor(slug)
}

func (s *service) FetchAllRetailMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadAllRetailMerchantCategoryMajor()
}

func (s *service) FetchAllSupplierMerchantCategoryMajor() ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadAllSupplierMerchantCategoryMajor()
}

func (s *service) FetchAllBySupplier(supplier int) ([]*ent.Product, error) {
	return s.repo.ReadAllBySupplierMerchant(supplier)
}

func (s *service) FetchAllByRetailer(retailer int) ([]*ent.Product, error) {
	return s.repo.ReadAllByRetailMerchant(retailer)
}

func (s *service) FetchBestSellerBySupplier() ([]*ent.Product, error) {
	return s.repo.ReadBestSellerBySupplierMerchant()
}

func (s *service) FetchBestSellerByRetailer() ([]*ent.Product, error) {
	return s.repo.ReadBestSellerRetailMerchant()
}

func (s *service) FetchAllMajorByRetailer(majorId int) ([]*ent.Product, error) {
	return s.repo.ReadAllMajorByRetailer(majorId)
}

func (s *service) Update(user *models.Product) (*models.Product, error) {
	return s.repo.Update(user)
}

func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
