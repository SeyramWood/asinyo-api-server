package merchant

import (
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.MerchantRepo
	noti notification.NotificationService
}

func NewMerchantService(repo gateways.MerchantRepo, noti notification.NotificationService) gateways.MerchantService {
	return &service{
		repo: repo,
		noti: noti,
	}
}

func (s *service) Create(merchant *models.MerchantRequest) (*ent.Merchant, error) {
	return s.repo.Insert(merchant, false)
}

func (s *service) Onboard(merchant *models.OnboardMerchantFullRequest, agentId int, logo string, images []string) (
	*ent.Merchant, error,
) {
	password, _ := application.GeneratePassword(12)
	merc, err := s.repo.Onboard(merchant, agentId, logo, images, password)
	if err != nil {
		return nil, err
	}
	URL := fmt.Sprintf("%s/auth/merchant/sign-in?username=%s", config.App().AsinyoURL, merc.Username)
	if application.UsernameType(merc.Username, "phone") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.SMSPayload{
					Recipients: []string{merc.Username},
					Message: fmt.Sprintf(
						"Congratulations! Asinyo merchant account has been created for you. Sign in with the username and password provided below\n\nUsername: %s\nPassword: %s\n\n%s",
						merc.Username,
						password,
						URL,
					),
				},
			},
		)
	}
	if application.UsernameType(merc.Username, "email") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.MailerMessage{
					To:       merc.Username,
					Subject:  "ASINYO MERCHANT ACCOUNT",
					Template: "newuser",
					Data: struct {
						Username string
						Password string
						URL      string
						Message  string
					}{
						merc.Username,
						password,
						URL,
						"Asinyo merchant account has been created for you. Your new account give you access to Asinyo merchant dashboard.",
					},
				},
			},
		)
	}
	return merc, nil
}

func (s *service) Fetch(id int) (*ent.Merchant, error) {
	return s.repo.Read(id)
}

func (s *service) FetchStorefront(id int) (*ent.MerchantStore, error) {
	return s.repo.ReadStorefront(id)
}

func (s *service) FetchAll(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}

func (s *service) Update(id int, request any) (*ent.Merchant, error) {
	return s.repo.Update(id, request)
}

func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
