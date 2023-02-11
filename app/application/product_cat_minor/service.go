package product_cat_minor

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
	repo gateways.ProductCatMinorRepo
}

func NewProductCatMinorService(repo gateways.ProductCatMinorRepo) gateways.ProductCatMinorService {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(cat *models.ProductCategoryMinor, image string) (*ent.ProductCategoryMinor, error) {

	return s.repo.Insert(cat, image)
}

func (s *service) Fetch(id int) (*ent.ProductCategoryMinor, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll(limit, offset int) ([]*ent.ProductCategoryMinor, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) Update(id int, request *models.ProductCategoryMinorUpdate) (*ent.ProductCategoryMinor, error) {
	return s.repo.Update(id, request)
}

func (s *service) UpdateImage(id int, imagePath string) (string, error) {
	return s.repo.UpdateImage(id, imagePath)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
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
