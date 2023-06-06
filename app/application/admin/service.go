package admin

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
	repo gateways.AdminRepo
	noti notification.NotificationService
}

func NewAdminService(repo gateways.AdminRepo, noti notification.NotificationService) gateways.AdminService {
	return &service{
		repo: repo,
		noti: noti,
	}
}

func (s *service) Create(user *models.AdminUserRequest) (*ent.Admin, error) {
	password, _ := application.GeneratePassword(12)
	newAdmin, err := s.repo.Insert(user, password)
	if err != nil {
		return nil, err
	}
	s.noti.Broadcast(
		&notification.Message{
			Data: services.MailerMessage{
				To:       user.Username,
				Subject:  "ASINYO ADMIN ACCOUNT",
				Template: "newadmin",
				Data: struct {
					Username string
					Password string
					URL      string
				}{
					user.Username,
					password,
					fmt.Sprintf("%s?username=%s", config.App().AsinyoAdminURL, user.Username),
				},
			},
		},
		&notification.Message{
			Data: services.DBNotificationMessage{
				AdminIDs:    []int{newAdmin.ID},
				SubjectType: "Admin",
				SubjectId:   newAdmin.ID,
				CreatorType: "Admin",
				CreatorId:   85899345920,
				Event:       "NewAdmin",
				Activity:    "New admin user created",
				Description: fmt.Sprintf(
					"New admin user account created for %s. Account creator: Seyram", user.Username,
				),
				Data: newAdmin,
			},
		},
	)
	return newAdmin, nil
}

func (s *service) OnboardNewCustomer(manager int, business *models.BusinessCustomerOnboardRequest) (
	*ent.Customer, error,
) {

	password, _ := application.GeneratePassword(12)
	newCustomer, err := s.repo.OnboardNewCustomer(manager, password, business)
	if err != nil {
		return nil, err
	}
	URL := fmt.Sprintf(
		"%s/auth/customer/sign-in?username=%s", config.App().AsinyoURL, business.Detail.Username,
	)
	creator, _ := s.repo.Read(manager)
	if application.UsernameType(business.Detail.Username, "phone") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.DBNotificationMessage{
					AdminIDs:    []int{manager},
					CustomerIDs: []int{newCustomer.ID},
					SubjectType: "Customer",
					SubjectId:   newCustomer.ID,
					CreatorType: "Admin",
					CreatorId:   manager,
					Event:       "NewCustomer",
					Activity:    "New business account created",
					Description: fmt.Sprintf(
						"New bussiness account created for %s. \nAccount creator: %s", business.Detail.Username,
						creator.Username,
					),
					Data: newCustomer,
				},
			},
			&notification.Message{
				Data: services.SMSPayload{
					Message: fmt.Sprintf(
						"Congratulations! Asinyo business account has been created for you. Sign in with the username and password provided below\n\nUsername: %s\nPassword: %s\n\n%s",
						business.Detail.Username,
						password,
						URL,
					),
					Recipients: []string{business.Detail.Username},
				},
			},
		)
	}
	if application.UsernameType(business.Detail.Username, "email") {
		s.noti.Broadcast(
			&notification.Message{
				Data: services.DBNotificationMessage{
					AdminIDs:    []int{manager},
					CustomerIDs: []int{newCustomer.ID},
					SubjectType: "Customer",
					SubjectId:   newCustomer.ID,
					CreatorType: "Admin",
					CreatorId:   manager,
					Event:       "NewCustomer",
					Activity:    "New business account created",
					Description: fmt.Sprintf(
						"New bussiness account created for %s. \nAccount creator: %s", business.Detail.Username,
						creator.Username,
					),
					Data: newCustomer,
				},
			},
			&notification.Message{
				Data: services.MailerMessage{
					To:       business.Detail.Username,
					Subject:  "ASINYO BUSINESS ACCOUNT",
					Template: "newuser",
					Data: struct {
						Username string
						Password string
						URL      string
						Message  string
					}{
						business.Detail.Username,
						password,
						URL,
						"Asinyo business account has been created for you. Your new account give you access to Asinyo Marketplace.",
					},
				},
			},
		)
	}

	return newCustomer, nil
}

func (s *service) Fetch(id int) (*ent.Admin, error) {
	return s.repo.Read(id)
}

func (s *service) FetchAll(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAll(limit, offset)
}
func (s *service) FetchMyClients(manager, limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadMyClients(manager, limit, offset)
}
func (s *service) FetchMyClientsPurchaseRequest(manager int) ([]*ent.Customer, error) {
	return s.repo.ReadMyClientsPurchaseRequest(manager)
}
func (s *service) FetchMyClientOrders(manager, limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadMyClientOrders(manager, limit, offset)
}
func (s *service) FetchAllOrders(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAllOrders(limit, offset)
}
func (s *service) FetchAdminProducts(limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadAdminProducts(limit, offset)
}
func (s *service) FetchProducts(major, minor string, limit, offset int) (*presenters.PaginationResponse, error) {
	return s.repo.ReadProducts(major, minor, limit, offset)
}

func (s *service) FetchAccountManagers(permissions ...string) ([]*ent.Admin, error) {
	return s.repo.ReadAllByPermissions(permissions)
}

func (s *service) FetchCounts(span string) (*presenters.DashboardRecordCount, error) {
	return s.repo.ReadCounts(span)
}

func (s *service) FetchConfigurations() ([]*ent.Configuration, error) {
	return s.repo.ReadConfigurations()
}

func (s *service) FetchConfigurationByIdOrName(slug any) (*ent.Configuration, error) {
	return s.repo.ReadConfigurationByIdOrName(slug)
}

func (s *service) Update(id int, user *models.AdminUserRequest) (*ent.Admin, error) {
	return s.repo.Update(id, user)
}

func (s *service) AssignAccountManager(manager, customer int) (*ent.Customer, error) {
	return s.repo.AssignAccountManager(manager, customer)
}

func (s *service) UpdateCurrentConfiguration(id int, configType, configValue string) (*ent.Configuration, error) {
	return s.repo.UpdateCurrentConfiguration(id, configType, configValue)
}

func (s *service) Remove(id int) error {
	return s.repo.Delete(id)
}
func (s *service) RemoveOrder(id int) error {
	return s.repo.DeleteOrder(id)
}
