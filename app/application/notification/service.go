package notification

import (
	"container/list"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
)

type (
	NotiType int
	Message  struct {
		Data any
	}
)

const (
	NewUser NotiType = iota
	NewOrder
	PasswordReset
	VERIFICATION
)

type noti struct {
	subs   *list.List
	dbNoti gateways.DBNotificationService
	sms    gateways.SMSService
	mail   gateways.EmailService
}

func NewNotification(
	dbNoti gateways.DBNotificationService, sms gateways.SMSService, mail gateways.EmailService,
) NotificationService {

	return &noti{
		subs:   new(list.List),
		dbNoti: dbNoti,
		sms:    sms,
		mail:   mail,
	}
}

func (n *noti) Subscribe(listener Listener) {
	n.subs.PushBack(listener)

}

func (n *noti) Unsubscribe(listener Listener) {
	for i := n.subs.Front(); i != nil; i = i.Next() {
		if i.Value.(Listener) == listener {
			n.subs.Remove(i)
		}
	}
}

func (n *noti) Broadcast(message ...*Message) {
	for i := n.subs.Front(); i != nil; i = i.Next() {
		i.Value.(Listener).listen(message...)
	}
}

func (n *noti) listen(message ...*Message) {

	if message == nil {
		fmt.Println("No notification message found")
	} else {
		for _, m := range message {
			if data, ok := m.Data.(services.DBNotificationMessage); ok {
				n.dbNoti.Send(&data)
			}
			if data, ok := m.Data.(services.MailerMessage); ok {
				n.mail.Send(&data)
			}
			if data, ok := m.Data.(services.SMSPayload); ok {
				n.sms.Send(&data)
			}
		}
	}

}
