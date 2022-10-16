package product

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/pkg/storage"
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
func (s *service) FetchBySlugSupplierMerchantCategoryMajor(slug string) ([]*ent.ProductCategoryMajor, error) {
	return s.repo.ReadBySlugSupplierMerchantCategoryMajor(slug)
}

func (s *service) FetchBySlugSupplierMerchantCategoryMinor(slug string) ([]*ent.ProductCategoryMinor, error) {
	return s.repo.ReadBySlugSupplierMerchantCategoryMinor(slug)
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

func (s *service) FetchBestSellerBySupplier(limit, offset int) ([]*ent.Product, error) {
	return s.repo.ReadBestSellerBySupplierMerchant(limit, offset)
}

func (s *service) FetchBestSellerByRetailer() ([]*ent.Product, error) {
	return s.repo.ReadBestSellerRetailMerchant()
}

func (s *service) Update(user *models.Product) (*models.Product, error) {
	return s.repo.Update(user)
}

func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}

func (s service) SaveImage(c *fiber.Ctx, field, directory string) (map[string]string, error) {
	file, _ := c.FormFile(field)
	pubDisk := storage.NewStorage().Disk("public")
	if !pubDisk.Exist(directory) {
		if err := pubDisk.MakeDirectory(directory); err != nil {
			if err != nil {
				return nil, err
			}
		}
	}
	dir, err := pubDisk.GetPath(directory)
	if err != nil {
		return nil, err
	}
	filename, err := s.saveFile(c, file, dir)
	if err != nil {
		return nil, err
	}
	return map[string]string{
		"file": filepath.Join(dir, filename.(string)),
		"url":  fmt.Sprintf("%s/%s", c.BaseURL(), filepath.Join(directory, filename.(string))),
	}, nil
}

func (s service) saveFile(c *fiber.Ctx, file *multipart.FileHeader, directory string) (interface{}, error) {
	buffer, err := file.Open()
	if err != nil {
		return nil, err
	}
	head := make([]byte, 512)
	buffer.Read(head)
	buffer.Close()

	mtype := mimetype.Detect(head)

	filename := fmt.Sprintf("asinyo_%s%s", uuid.New(), mtype.Extension())

	if err := c.SaveFile(file, filepath.Join("./", directory, filename)); err != nil {
		return nil, err
	}
	return filename, nil
}
