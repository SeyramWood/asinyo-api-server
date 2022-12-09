package order

import (
	"fmt"

	"github.com/Jeffail/gabs"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/sms"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo     gateways.OrderRepo
	logistic gateways.LogisticService
	sms      gateways.SMSService
	mail     gateways.EmailService
}

func NewOrderService(
	repo gateways.OrderRepo, logistic gateways.LogisticService, mail gateways.EmailService,
) gateways.OrderService {
	smsService := sms.NewSMSService()
	return &service{
		repo:     repo,
		logistic: logistic.New(repo),
		sms:      smsService,
		mail:     mail,
	}
}

func (s service) Create(order *models.OrderPayload) (*ent.Order, error) {
	result, err := s.repo.Insert(order)
	if err != nil {
		return nil, err
	}
	storeMerchant, err := s.repo.ReadOrderStoreMerchants(result.ID)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf(
		"New Order Recieved! Please visit your portal to start the order process. %s",
		config.App().AppURL,
	)
	for _, store := range storeMerchant.Edges.Stores {
		if application.UsernameType(store.Edges.Merchant.Username, "phone") {
			_, _ = s.sms.Send(
				&services.SMSPayload{
					Recipients: []string{store.Edges.Merchant.Username},
					Message:    msg,
				},
			)
		}
		if application.UsernameType(store.Edges.Merchant.Username, "email") {
			s.mail.Send(
				&services.Message{
					To:      store.Edges.Merchant.Username,
					Subject: "ASINYO ORDER NOTIFICATION",
					Data:    msg,
				},
			)
		}
	}
	return result, nil
}

func (s service) TesCreate(orderId int) (*ent.Order, error) {
	// TODO Send SMS OR Email to store
	storeMerchant, err := s.repo.ReadOrderStoreMerchants(orderId)
	if err != nil {
		return nil, err
	}
	msg := fmt.Sprintf(
		"New Order Recieved! Please visit your portal to start the order process. %s",
		config.App().AppURL,
	)
	for _, store := range storeMerchant.Edges.Stores {
		if application.UsernameType(store.Edges.Merchant.Username, "phone") {
			_, _ = s.sms.Send(
				&services.SMSPayload{
					Recipients: []string{store.Edges.Merchant.Username},
					Message:    msg,
				},
			)
		}
		if application.UsernameType(store.Edges.Merchant.Username, "email") {
			s.mail.Send(
				&services.Message{
					To:      store.Edges.Merchant.Username,
					Subject: "ASINYO ORDER NOTIFICATION",
					Data:    msg,
				},
			)
		}
	}
	return storeMerchant, nil
}

func (s service) FetchAllByUser(userType string, id int) ([]*ent.Order, error) {
	return s.repo.ReadAllByUser(userType, id)
}
func (s service) FetchAllByStore(merchantId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByStore(merchantId)
}
func (s service) FetchAllByAgentStore(agentId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByAgentStore(agentId)
}
func (s service) FetchByStore(id, userId int, userType string) (*ent.Order, error) {
	if userType == "agent" {
		return s.repo.ReadByAgentStore(id, userId)
	}
	return s.repo.ReadByStore(id, userId)
}

func (s service) FetchByUser(userType string, id int) (*ent.Order, error) {
	return s.repo.ReadByUser(userType, id)
}
func (s service) FetchAll() ([]*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s service) Fetch(id int) (*ent.Order, error) {
	return s.repo.Read(id)
}

func (s service) Update(order *services.PaystackResponse) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s service) Remove(id string) error {
	// TODO implement me
	panic("implement me")
}

func (s service) UpdateOrderDetailStatus(request []byte) (*ent.Order, error) {
	resBody, err := gabs.ParseJSON(request)
	if err != nil {
		return nil, err
	}
	statuses, errr := resBody.ChildrenMap()

	if errr != nil {
		return nil, errr
	}

	if result, err := s.repo.UpdateOrderDetailStatus(statuses); err != nil {
		return nil, err
	} else {
		if result.Status == "in_progress" {
			s.logistic.ExecuteTask(result, "pickup_delivery_tasks")
			// if result.DeliveryTask {
			// 	s.logistic.DoTask(result, "edit_multiple_tasks")
			// } else {
			// 	s.logistic.DoTask(result, "create_multiple_tasks")
			// }
		}
		return result, nil
	}
}
