package merchant_store

import (
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"
	"sync"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/pkg/storage"
)

type service struct {
	repo gateways.MerchantStoreRepo
}

func NewMerchantStoreService(repo gateways.MerchantStoreRepo) gateways.MerchantStoreService {
	return &service{
		repo: repo,
	}
}

func (s service) Create(store *models.MerchantStoreRequest, merchantId int, logo string, images []string) (
	*ent.MerchantStore, error,
) {
	return s.repo.Insert(store, merchantId, logo, images)
}

func (s service) SaveAccount(store interface{}, storeId int, accountType string) (*ent.MerchantStore, error) {
	return s.repo.UpdateAccount(store, storeId, accountType)
}

func (s service) SaveAgentPermission(permission bool, storeId int) (*ent.MerchantStore, error) {
	return s.repo.UpdateAgentPermission(permission, storeId)
}

func (s service) SaveDefaultAccount(storeId int, accountType string) (*ent.MerchantStore, error) {
	return s.repo.UpdateDefaultAccount(storeId, accountType)
}

func (s service) FetchAll() ([]*ent.MerchantStore, error) {
	return s.repo.ReadAll()
}
func (s service) FetchAllByMerchant(merchantType string, limit, offset int) ([]*ent.MerchantStore, error) {
	return s.repo.ReadAllByMerchant(merchantType, limit, offset)
}

func (s service) Fetch(id int) (*ent.MerchantStore, error) {
	return s.repo.Read(id)
}
func (s service) FetchByMerchant(merchantId int) (*ent.MerchantStore, error) {
	return s.repo.ReadByMerchant(merchantId)
}
func (s service) FetchAgent(store int) (*ent.Agent, error) {
	return s.repo.ReadAgent(store)
}
func (s service) Update(store *models.MerchantStore, storeId int) (*ent.MerchantStore, error) {
	// TODO implement me
	panic("implement me")
}
func (s service) UpdateAddress(address *models.MerchantStoreAddress, storeId int) (*ent.MerchantStore, error) {
	return s.repo.UpdateAddress(address, storeId)
}
func (s service) Remove(id string) error {
	// TODO implement me
	panic("implement me")
}

func (s service) SaveLogo(c *fiber.Ctx, field, directory string) (interface{}, error) {
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
	return fmt.Sprintf("%s/%s", c.BaseURL(), filepath.Join(directory, filename.(string))), nil
}
func (s service) SavePhotos(c *fiber.Ctx, field, directory string) (interface{}, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File[field]

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

	var urls []string
	wg := sync.WaitGroup{}
	for _, file := range files {
		wg.Add(1)
		go func(f *multipart.FileHeader) {
			defer wg.Done()
			filename, err := s.saveFile(c, f, dir)
			if err != nil {
				log.Fatalln(fmt.Sprintf("error saving [%s]\n[error]: %s", filename, err))
			}
			urls = append(urls, fmt.Sprintf("%s/%s/%s", c.BaseURL(), directory, filename))
		}(file)
	}
	wg.Wait()

	return urls, nil
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
