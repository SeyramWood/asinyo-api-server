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

func (s *service) FetchAll(limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) FetchBySlugRetailMerchantCategoryMajor(slug string, limit, offset int) (
	[]*ent.ProductCategoryMajor, error,
) {
	return s.repo.ReadBySlugRetailMerchantCategoryMajor(slug, limit, offset)
}

func (s *service) FetchBySlugRetailMerchantCategoryMinor(slug string, limit, offset int) (
	[]*ent.ProductCategoryMinor, error,
) {
	return s.repo.ReadBySlugRetailMerchantCategoryMinor(slug, limit, offset)
}
func (s *service) FetchBySlugSupplierMerchantCategoryMajor(slug string, limit, offset int) (
	[]*ent.ProductCategoryMajor, error,
) {
	return s.repo.ReadBySlugSupplierMerchantCategoryMajor(slug, limit, offset)
}

func (s *service) FetchBySlugSupplierMerchantCategoryMinor(slug string, limit, offset int) (
	[]*ent.ProductCategoryMinor, error,
) {
	return s.repo.ReadBySlugSupplierMerchantCategoryMinor(slug, limit, offset)
}

func (s *service) FetchAllBySlugCategoryMajor(
	merchantType, slug string, limit, offset int,
) ([]*ent.Product, error) {
	return s.repo.ReadAllBySlugCategoryMajor(merchantType, slug, limit, offset)
}

func (s *service) FetchAllBySlugCategoryMinor(merchantType, slug string, limit, offset int) (
	[]*ent.Product, error,
) {
	return s.repo.ReadAllBySlugCategoryMinor(merchantType, slug, limit, offset)

}

func (s *service) FetchAllRetailMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadAllRetailMerchantCategoryMajor(limit, offset)
}

func (s *service) FetchAllSupplierMerchantCategoryMajor(limit, offset int) ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadAllSupplierMerchantCategoryMajor(limit, offset)
}

func (s *service) FetchAllBySupplier(supplier, limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadAllBySupplierMerchant(supplier, limit, offset)
}

func (s *service) FetchAllByRetailer(retailer, limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadAllByRetailMerchant(retailer, limit, offset)
}

func (s *service) FetchBestSellerBySupplier(limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadBestSellerBySupplierMerchant(limit, offset)
}

func (s *service) FetchBestSellerByRetailer(limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadBestSellerRetailMerchant(limit, offset)
}
func (s *service) FetchBestSellerByMerchant(id, limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadBestSellerByMerchant(id, limit, offset)
}

func (s *service) Update(id int, request *models.ProductUpdate) (*ent.Product, error) {
	return s.repo.Update(id, request)
}

func (s *service) UpdateImage(id int, imagePath string) (string, error) {
	return s.repo.UpdateImage(id, imagePath)
}
func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}
