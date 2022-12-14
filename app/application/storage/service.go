package storage

import (
	"sync"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

var drivers = map[string]string{
	"local":      "local",
	"uploadcare": "uploadcare",
}

type storage struct {
	WG        *sync.WaitGroup
	DataChan  chan any
	DoneChan  chan bool
	ErrorChan chan error
}

func NewStorageService(wg *sync.WaitGroup) gateways.StorageService {
	dataChan := make(chan any, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)

	conf := &storage{
		WG:        wg,
		DataChan:  dataChan,
		DoneChan:  doneChan,
		ErrorChan: errorChan,
	}

	switch config.App().FilesystemDriver {
	case "uploadcare":
		return newUploadcare(conf)
	default:
		return newLocal(conf)
	}
}
