package notification

type (
	Producer interface {
		Produce(message any)
	}
	Listener interface {
		listen(message ...*Message)
	}
	NotificationService interface {
		Subscribe(listener Listener)
		Unsubscribe(listener Listener)
		Broadcast(message ...*Message)
		Listener
	}
)
