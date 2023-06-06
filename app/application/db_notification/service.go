package db_notification

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
)

type dbnotification struct {
	repo              gateways.DBNotificationRepo
	FailedCount       float64
	FailedDelay       time.Duration
	WG                *sync.WaitGroup
	DataChan          chan *services.DBNotificationMessage
	FailedDataChan    chan *services.DBNotificationMessage
	NotificationsChan chan []*ent.Notification
	NewNotifications  []*ent.Notification
	NewClients        chan *services.NotificationClient
	ClosingClients    chan string
	Clients           map[string]*services.NotificationClient
	MarkAsReadChan    chan map[string]any
	DoneChan          chan bool
	ErrorChan         chan error
}

func NewDBNotificationService(wg *sync.WaitGroup, adapter *database.Adapter) gateways.DBNotificationService {
	dataChan := make(chan *services.DBNotificationMessage, 1024)
	failedDataChan := make(chan *services.DBNotificationMessage, 1024)
	notificationsChan := make(chan []*ent.Notification, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)

	newNotifications := make([]*ent.Notification, 0)
	newClients := make(chan *services.NotificationClient)
	closingClients := make(chan string)
	clients := make(map[string]*services.NotificationClient)
	markAsRead := make(chan map[string]any)

	return &dbnotification{
		repo:              NewDBNotificationRepo(adapter),
		FailedCount:       0,
		FailedDelay:       1 * time.Second,
		WG:                wg,
		DataChan:          dataChan,
		FailedDataChan:    failedDataChan,
		NotificationsChan: notificationsChan,
		NewNotifications:  newNotifications,
		NewClients:        newClients,
		ClosingClients:    closingClients,
		Clients:           clients,
		MarkAsReadChan:    markAsRead,
		DoneChan:          doneChan,
		ErrorChan:         errorChan,
	}
}

func (n *dbnotification) Listen() {
	for {
		select {
		case msg := <-n.DataChan:
			go n.createNotification(msg)
		case msg := <-n.FailedDataChan:
			if n.FailedCount <= 3 {
				n.FailedDelay = time.Duration(n.FailedCount) * n.FailedDelay
				time.Sleep(n.FailedDelay)
				go n.createNotification(msg)
			}
		case c := <-n.NewClients:
			n.Clients[c.ID] = c
		case c := <-n.ClosingClients:
			delete(n.Clients, c)
		case request := <-n.MarkAsReadChan:
			n.markAsRead(request)
		case <-time.After(5 * time.Second):
			for _, client := range n.Clients {
				client.Message <- n.getNewNotifications(client.Params)
			}
		case err := <-n.ErrorChan:
			log.Println(err)
		case <-n.DoneChan:
			return
		}
	}
}

func (n *dbnotification) Send(msg *services.DBNotificationMessage) {
	n.WG.Add(1)
	n.DataChan <- msg
}
func (n *dbnotification) Done() {
	n.DoneChan <- true
}
func (n *dbnotification) CloseChannels() {
	close(n.DataChan)
	close(n.FailedDataChan)
	close(n.NotificationsChan)
	close(n.NewClients)
	close(n.ClosingClients)
	close(n.MarkAsReadChan)
	close(n.ErrorChan)
	close(n.DoneChan)
}
func (n *dbnotification) Broker(
	action string, clientId string, params *services.NotificationFilters,
) *services.NotificationClient {
	if action == "get_client" {
		if client, ok := n.Clients[clientId]; ok {
			return client
		}
	}
	if action == "add_client" {
		client := &services.NotificationClient{
			ID:      clientId,
			Params:  params,
			Message: make(chan []byte),
		}
		n.NewClients <- client
		return client
	}
	if action == "close_client" {
		n.ClosingClients <- clientId
		return nil
	}
	return nil
}
func (n *dbnotification) FetchNotifications(limit, offset int) ([]*ent.Notification, error) {
	// results, err := n.repo.ReadNotifications()
	// if err != nil {
	// 	n.ErrorChan <- err
	// } else {
	// 	response <- results
	// }
	return nil, nil
}

func (n *dbnotification) FetchUserNotifications(params *services.NotificationFilters) (
	[]*ent.Notification, error,
) {
	if params.UserType == "asinyo" {
		return n.repo.ReadAdminNotifications(params)
	}
	return n.repo.ReadUserNotifications(params)
}

func (n *dbnotification) MarkAsRead(userId, notificationId int, userType, timestamp string) (*ent.Notification, error) {
	return n.repo.MarkAsRead(userId, notificationId, userType, timestamp)
}
func (n *dbnotification) MarkSelectedAsRead(userId int, notificationIds []int, userType, timestamp string) error {
	request := make(map[string]any)
	request["userId"] = userId
	request["notificationIds"] = notificationIds
	request["userType"] = userType
	request["timestamp"] = timestamp

	n.MarkAsReadChan <- request

	return nil
	// return n.repo.MarkSelectedAsRead(userId, notificationIds, userType, timestamp)
}

func (n *dbnotification) RemoveSelected(ids []int) error {
	return n.repo.RemoveSelected(ids)
}

func (n *dbnotification) Remove(id int) error {
	return n.repo.Delete(id)
}

func (n *dbnotification) markAsRead(request map[string]any) {
	notificationIds := request["notificationIds"].([]int)
	if len(notificationIds) > 0 {
		userId := request["userId"].(int)
		userType := request["userType"].(string)
		timestamp := request["timestamp"].(string)
		for _, id := range notificationIds {
			if _, err := n.repo.MarkAsRead(userId, id, userType, timestamp); err != nil {
				n.ErrorChan <- err
			}
		}
	}
}
func (n *dbnotification) createNotification(msg *services.DBNotificationMessage) {
	if _, err := n.repo.Insert(msg); err != nil {
		n.ErrorChan <- err
	}
}
func (n *dbnotification) getNewNotifications(params *services.NotificationFilters) []byte {
	if params.UserType == "asinyo" {
		noti, err := n.repo.ReadAdminNotifications(params)
		if err != nil {
			n.ErrorChan <- err
			return []byte{}
		}
		if len(noti) == 0 {
			return []byte{}
		}
		byteData, err := json.Marshal(presenters.FormatNotificationsResponse(noti))
		if err != nil {
			n.ErrorChan <- err
			return []byte{}
		}
		return byteData
	}
	noti, err := n.repo.ReadUserNotifications(params)
	if err != nil {
		n.ErrorChan <- err
		return []byte{}
	}
	if len(noti) == 0 {
		return []byte{}
	}
	byteData, err := json.Marshal(presenters.FormatNotificationsResponse(noti))
	if err != nil {
		n.ErrorChan <- err
		return []byte{}
	}

	return byteData
}
