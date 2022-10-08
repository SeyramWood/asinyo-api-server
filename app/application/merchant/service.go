package merchant

import (
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/sms"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.MerchantRepo
	sms  gateways.SMSService
	mail gateways.EmailService
}

func NewMerchantService(repo gateways.MerchantRepo, mail gateways.EmailService) gateways.MerchantService {
	smsService := sms.NewSMSService()
	return &service{
		repo: repo,
		sms:  smsService,
		mail: mail,
	}
}

func (s *service) Create(merchant *models.MerchantRequest) (*ent.Merchant, error) {

	return s.repo.Insert(merchant, false)
}

func (s *service) Onboard(merchant *models.StoreFinalRequest, agentId int, logo string, images []string) (
	*ent.Merchant, error,
) {
	password, _ := application.GenerateOTP(12)
	merc, err := s.repo.Onboard(merchant, agentId, logo, images, password)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf(
		"Congratulations for joining Asinyo! Please enter the OTP (password) to continue with Asinyo. Make sure to change your password when you successfuly sign in.%s",
		password,
	)
	if application.UsernameType(merchant.Username, "phone") {
		_, err := s.sms.Send(
			&services.SMSPayload{
				Recipients: []string{merchant.Username},
				Message:    msg,
			},
		)
		if err != nil {
			return nil, err
		}
	}
	if application.UsernameType(merchant.Username, "email") {
		s.mail.Send(
			&services.Message{
				To:      merchant.Username,
				Subject: "ASINYO MERCHANT OTP",
				Data:    msg,
			},
		)
	}
	return merc, nil
}

func (s *service) Fetch(id int) (*ent.Merchant, error) {

	return s.repo.Read(id)
}

func (s *service) FetchAll() ([]*ent.Merchant, error) {
	return s.repo.ReadAll()
}

func (s *service) Update(user *models.Merchant) (*models.Merchant, error) {
	return s.repo.Update(user)
}

func (s *service) Remove(ID string) error {
	return s.repo.Delete(ID)
}
