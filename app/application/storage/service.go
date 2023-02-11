package storage

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

var drivers = map[string]string{
	"local":      "local",
	"uploadcare": "uploadcare",
}

type storageDriverService interface {
	uploadFile(dir string, f *multipart.FileHeader) (string, error)
	uploadFiles(dir string, files []*multipart.FileHeader) ([]string, error)
	deleteFile(path any) error
	deleteFiles(paths any) error
}
type storage struct {
	WG        *sync.WaitGroup
	DataChan  chan any
	DoneChan  chan bool
	ErrorChan chan error
	TaskType  string
	diskType  storageDriverService
}

func NewStorageService(wg *sync.WaitGroup) gateways.StorageService {
	dataChan := make(chan any, 1024)
	doneChan := make(chan bool)
	errorChan := make(chan error)
	var diskType storageDriverService

	switch config.App().FilesystemDriver {
	case "uploadcare":
		diskType = newUploadcare()
	default:
		diskType = newLocal()
	}
	return &storage{
		WG:        wg,
		DataChan:  dataChan,
		DoneChan:  doneChan,
		ErrorChan: errorChan,
		TaskType:  "delete_file",
		diskType:  diskType,
	}

}

func (s *storage) Listen() {
	for {
		select {
		case data := <-s.DataChan:
			go s.runTask(data)
		case err := <-s.ErrorChan:
			fmt.Println("error:: ", err)
		case <-s.DoneChan:
			return
		}
	}
}

func (s *storage) ExecuteTask(data any, taskType string) {
	s.WG.Add(1)
	s.DataChan <- data
	if taskType != "" {
		s.TaskType = taskType
	}
}

func (s *storage) Done() {
	s.DoneChan <- true
}

func (s *storage) CloseChannels() {
	close(s.DataChan)
	close(s.ErrorChan)
	close(s.DoneChan)
}

func (s *storage) Disk(disk string) gateways.StorageService {
	if v, ok := drivers[disk]; ok && v == "uploadcare" {
		s.diskType = newUploadcare()
		return s
	}
	s.diskType = newLocal()
	return s
}

func (s *storage) UploadFile(dir string, f *multipart.FileHeader) (string, error) {
	return s.diskType.uploadFile(dir, f)
}

func (s *storage) UploadFiles(dir string, files []*multipart.FileHeader) ([]string, error) {
	return s.diskType.uploadFiles(dir, files)
}

func (s *storage) runTask(data any) {
	defer s.WG.Done()
	switch s.TaskType {
	case "delete_file":
		if err := s.deleteFile(data); err != nil {
			s.ErrorChan <- err
		}
	case "delete_files":
		if err := s.deleteFiles(data); err != nil {
			s.ErrorChan <- err
		}
	}
}

func (s *storage) deleteFile(path any) error {
	return s.diskType.deleteFile(path)
}

func (s *storage) deleteFiles(paths any) error {
	return s.diskType.deleteFiles(paths)
}
