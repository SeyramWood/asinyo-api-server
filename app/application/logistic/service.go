package logistic

import (
	"log"
	"sync"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/framework/database"
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
	gateway, err := getGateway(repo)
	if err != nil {
		log.Println(err)
		log.Panicln("Could not fetch logistic gateway")
		return nil
	}
	switch gateway {
	case "Tookan":
		return newTookanService(conf, repo)
	case "Asinyo":
		return newAsinyoService(conf, repo)
	default:
		log.Panicln("Failed to instantiate a logistic service")
		return nil
	}
}

func getGateway(repo gateways.LogisticRepo) (string, error) {
	gateway, err := repo.ReadLogistic()
	if !ent.IsNotFound(err) {
		return "", err
	}
	if ent.IsNotFound(err) {
		return "Asinyo", nil
	}
	gatewayData := gateway.Data.Data.(map[string]any)
	return gatewayData["current"].(string), nil
}
