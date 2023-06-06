package sms

import (
	"sync"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

type sms struct {
	WG             *sync.WaitGroup
	DataChan       chan any
	FailedDataChan chan any
	DoneChan       chan bool
	ErrorChan      chan error
}

func NewSMSService(wg *sync.WaitGroup) gateways.SMSService {
	dataChan := make(chan any, 1024)
	failedDataChan := make(chan any, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)

	conf := &sms{
		WG:             wg,
		DataChan:       dataChan,
		FailedDataChan: failedDataChan,
		DoneChan:       doneChan,
		ErrorChan:      errorChan,
	}
	switch config.SMS().Gateway {
	case "arkesel":
		return newArkeselService(conf)
	default:
		return nil
	}
}
