package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/uploadcare/uploadcare-go/ucare"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

var disks = map[string]string{
	"public": "public/local",
	"local":  "local",
}

type local struct {
	secretKey string
	publicKey string
	url       string
	client    ucare.Client
	ctx       context.Context
	TaskType  string
	WG        *sync.WaitGroup
	DataChan  chan any
	DoneChan  chan bool
	ErrorChan chan error
	conf      *storage
}

func newLocal(conf *storage) gateways.StorageService {
	return &local{
		secretKey: config.Uploadcare().SecKey,
		publicKey: config.Uploadcare().PubKey,
		url:       config.Uploadcare().URL,
		client:    nil,
		ctx:       context.Background(),
		TaskType:  "delete_file",
		WG:        conf.WG,
		DataChan:  conf.DataChan,
		DoneChan:  conf.DoneChan,
		ErrorChan: conf.ErrorChan,
		conf:      conf,
	}
}

func (l *local) Listen() {
	for {
		select {
		case data := <-l.DataChan:
			go l.runTask(data)
		case err := <-l.ErrorChan:
			fmt.Println(err)
		case <-l.DoneChan:
			return
		}
	}
}

func (l *local) ExecuteTask(data any, taskType string) {
	l.WG.Add(1)
	l.DataChan <- data
	if taskType != "" {
		l.TaskType = taskType
	}
}

func (l *local) Done() {
	l.DoneChan <- true
}

func (l *local) CloseChannels() {
	close(l.DataChan)
	close(l.ErrorChan)
	close(l.DoneChan)
}

func (l *local) Disk(disk string) gateways.StorageService {
	if disk == drivers["uploadcare"] {
		return newUploadcare(l.conf)
	}
	return newLocal(l.conf)
}

func (l *local) UploadFile(dir string, f *multipart.FileHeader) (string, error) {
	// TODO implement me
	panic("implement me local")
}

func (l *local) UploadFiles(dir string, f []*multipart.FileHeader) ([]*string, error) {
	// TODO implement me
	panic("implement me")
}

func (l *local) runTask(data any) {
	defer l.WG.Done()
	switch l.TaskType {
	case "delete_file":
		if err := l.deleteFile(data); err != nil {
			l.ErrorChan <- err
		}
	}
}

func (l *local) deleteFile(path any) error {
	// TODO implement me
	panic("implement me")
}

func (l *local) deleteFiles(path []*string) error {
	// TODO implement me
	panic("implement me")
}
