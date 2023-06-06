package order

import (
	"fmt"

	"github.com/Jeffail/gabs"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/application/notification"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo     gateways.OrderRepo
	logistic gateways.LogisticService
	noti     notification.NotificationService
}

func NewOrderService(
	repo gateways.OrderRepo, logistic gateways.LogisticService, noti notification.NotificationService,
) gateways.OrderService {
	return &service{
		repo:     repo,
		logistic: logistic.OrderRepo(repo),
		noti:     noti,
	}
}

func (s *service) Create(order *models.OrderPayload, params ...int) (*ent.Order, error) {
	result, err := s.repo.Insert(order, params...)
	if err != nil {
		return nil, err
	}
	if params != nil {
		if c, err := result.Edges.CustomerOrErr(); err == nil {
			URL := fmt.Sprintf("%s", config.App().AsinyoURL)
			if application.UsernameType(
				c.Username, "phone",
			) && application.UsernameType(c.Username, "phone") {
				s.noti.Broadcast(
					&notification.Message{
						Data: services.SMSPayload{
							Recipients: []string{c.Username},
							Message: fmt.Sprintf(
								"New Purchase Request!\nYour Asinyo account manager placed a new purchase order on your behalf. Please visit your portal for reviewal and approval. Order #: %s\n%s",
								result.OrderNumber,
								URL,
							),
						},
					},
				)
			}
			if application.UsernameType(c.Username, "email") {
				s.noti.Broadcast(
					&notification.Message{
						Data: services.MailerMessage{
							To:       c.Username,
							Subject:  "ASINYO NEW PURCHASE REQUEST",
							Template: "newpurchaseorder",
							Data: struct {
								URL     string
								Message string
							}{
								URL,
								fmt.Sprintf(
									"Your Asinyo account manager placed a new purchase order on your behalf. Please visit your portal for reviewal and approval. Order #: %s",
									result.OrderNumber,
								),
							},
						},
					},
				)
			}

		}
		return result, nil
	}

	storeMerchant, err := s.repo.ReadOrderStoreMerchants(result.ID)
	if err != nil {
		return nil, err
	}
	s.sendOrderPlacedNotifications(storeMerchant)
	return result, nil
}
func (s *service) SaveOrderUpdate(id int, order *models.OrderPayload) (*ent.Order, error) {
	return s.repo.SaveOrderUpdate(id, order)
}
func (s *service) TestCreate(orderId int) (*ent.Order, error) {
	// TODO Send SMS OR Email to store
	storeMerchant, err := s.repo.ReadOrderStoreMerchants(orderId)
	if err != nil {
		return nil, err
	}
	// msg := fmt.Sprintf(
	// 	"New Order Received! Please visit your portal to start the order process. %s",
	// 	config.App().AppURL,
	// )
	// for _, store := range storeMerchant.Edges.Stores {
	// 	if application.UsernameType(store.Edges.Merchant.Username, "phone") {
	// 		_, _ = s.sms.Send(
	// 			&services.SMSPayload{
	// 				Recipients: []string{store.Edges.Merchant.Username},
	// 				MailerMessage:    msg,
	// 			},
	// 		)
	// 	}
	// 	if application.UsernameType(store.Edges.Merchant.Username, "email") {
	// 		s.mail.Send(
	// 			&services.MailerMessage{
	// 				To:      store.Edges.Merchant.Username,
	// 				Subject: "ASINYO ORDER NOTIFICATION",
	// 				Data:    msg,
	// 			},
	// 		)
	// 	}
	// }
	return storeMerchant, nil
}

func (s *service) FetchAllByUser(userType string, id int) ([]*ent.Order, error) {
	return s.repo.ReadAllByUser(userType, id)
}

func (s *service) FetchAllByStore(merchantId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByStore(merchantId)
}

func (s *service) FetchAllByAgentStore(agentId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByAgentStore(agentId)
}

func (s *service) FetchByStore(id, userId int, userType string) (*ent.Order, error) {
	if userType == "agent" {
		return s.repo.ReadByAgentStore(id, userId)
	}
	return s.repo.ReadByStore(id, userId)
}

func (s *service) FetchByUser(userType string, id int) (*ent.Order, error) {
	return s.repo.ReadByUser(userType, id)
}

func (s *service) FetchAll() ([]*ent.Order, error) {
	return s.repo.ReadAll()
}

func (s *service) Fetch(id int) (*ent.Order, error) {
	return s.repo.Read(id)
}

func (s *service) Update(order *services.PaystackResponse) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s *service) Remove(id string) error {
	// TODO implement me
	panic("implement me")
}

func (s *service) UpdateOrderDetailStatus(request []byte, logisticType string) (*ent.Order, error) {
	resBody, err := gabs.ParseJSON(request)
	if err != nil {
		return nil, err
	}
	requestData, err := resBody.ChildrenMap()
	if err != nil {
		return nil, err
	}
	var status map[string]*gabs.Container
	dispatch, err := requestData["dispatch"].ChildrenMap()
	if err != nil {
		status = requestData
	}
	if logisticType == "Asinyo" && dispatch != nil {
		status, err = requestData["status"].ChildrenMap()
		if err != nil {
			return nil, err
		}
	}

	if result, err := s.repo.UpdateOrderDetailStatus(status); err != nil {
		return nil, err
	} else {
		if result.Status == "in_progress" {
			if logisticType == "Asinyo" {
				dispatch, err := requestData["dispatch"].ChildrenMap()
				if err == nil {
					s.logistic.ExecuteTask(result, dispatch["action"].Data().(string), dispatch)
				}
			} else if logisticType == "Tookan" {
				s.logistic.ExecuteTask(result, "pickup_delivery_tasks")
			}
		}
		return result, nil
	}
}
func (s *service) UpdateOrderApprovalStatus(orderId int, status string) (*ent.Order, error) {
	result, err := s.repo.UpdateOrderApprovalStatus(orderId, status)
	storeMerchant, err := s.repo.ReadOrderStoreMerchants(result.ID)
	if err != nil {
		return nil, err
	}
	s.sendOrderPlacedNotifications(storeMerchant)
	return result, nil
}

func (s *service) sendOrderPlacedNotifications(result *ent.Order) {
	URL := fmt.Sprintf("%s", config.App().AsinyoURL)
	for _, store := range result.Edges.Stores {
		if application.UsernameType(
			store.Edges.Merchant.Username, "phone",
		) && application.UsernameType(store.Edges.Agent.Username, "phone") {
			s.noti.Broadcast(
				&notification.Message{
					Data: services.SMSPayload{
						Recipients: []string{store.Edges.Merchant.Username, store.Edges.Agent.Username},
						Message: fmt.Sprintf(
							"New Order Request! Visit your portal to start processing the order. Order #: %s\n%s",
							result.OrderNumber,
							URL,
						),
					},
				},
			)
		}
		if application.UsernameType(store.Edges.Merchant.Username, "email") {
			s.noti.Broadcast(
				&notification.Message{
					Data: services.MailerMessage{
						To:       store.Edges.Merchant.Username,
						Subject:  "ASINYO NEW ORDER REQUEST",
						Template: "neworder",
						Data: struct {
							URL     string
							Message string
						}{
							URL,
							fmt.Sprintf(
								"Visit your portal to start processing the order. Order #: %s",
								result.OrderNumber,
							),
						},
					},
				},
				&notification.Message{
					Data: services.MailerMessage{
						To:       store.Edges.Agent.Username,
						Subject:  "ASINYO NEW ORDER REQUEST",
						Template: "neworder",
						Data: struct {
							URL     string
							Message string
						}{
							URL,
							fmt.Sprintf(
								"Visit your portal to start processing the order. Order #: %s",
								result.OrderNumber,
							),
						},
					},
				},
			)
		}
	}
}
