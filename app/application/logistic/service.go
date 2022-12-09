package logistic

import (
	"sync"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type logistic struct {
	WG          *sync.WaitGroup
	DataChan    chan *ent.Order
	WebhookChan chan any
	DoneChan    chan bool
	ErrorChan   chan error
}

func NewLogistic(wg *sync.WaitGroup, adapter *database.Adapter) gateways.LogisticService {
	dataChan := make(chan *ent.Order, 1024)
	WebhookChan := make(chan any, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)

	conf := &logistic{
		WG:          wg,
		DataChan:    dataChan,
		WebhookChan: WebhookChan,
		DoneChan:    doneChan,
		ErrorChan:   errorChan,
	}

	repo := NewLogisticRepo(adapter)

	switch config.Logistic().Gateway {
	case "tookan":
		return newTookanService(conf, repo)
	default:
		return nil
	}
}
