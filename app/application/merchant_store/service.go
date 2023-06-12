package merchant_store

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

func (s service) SaveAccount(store any, storeId int, accountType string) (*ent.MerchantStore, error) {
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
func (s service) Update(request *models.MerchantStoreUpdate, storeId int) (*ent.MerchantStore, error) {
	return s.repo.Update(request, storeId)
}
func (s service) UpdateAddress(address *models.MerchantStoreAddress, storeId int) (*ent.MerchantStore, error) {
	return s.repo.UpdateAddress(address, storeId)
}
func (s service) UpdateBanner(storeId int, bannerPath string) (string, error) {
	return s.repo.UpdateBanner(storeId, bannerPath)
}

func (s service) UpdateImages(storeId int, newPath, oldPath string) ([]string, error) {
	return s.repo.UpdateImages(storeId, newPath, oldPath)
}
func (s service) AppendNewImages(storeId int, urls []string) ([]string, error) {
	return s.repo.AppendNewImages(storeId, urls)
}

func (s service) Remove(id string) error {
	// TODO implement me
	panic("implement me")
}

func (s service) saveFile(c *fiber.Ctx, file *multipart.FileHeader, directory string) (any, error) {
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
