package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"sync"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/uploadcare/uploadcare-go/ucare"
	"github.com/uploadcare/uploadcare-go/upload"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

type uploadcare struct {
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

func newUploadcare(conf *storage) gateways.StorageService {
	client, err := ucare.NewClient(
		ucare.APICreds{
			SecretKey: config.Uploadcare().SecKey,
			PublicKey: config.Uploadcare().PubKey,
		}, &ucare.Config{
			SignBasedAuthentication: true,
			APIVersion:              ucare.APIv06,
		},
	)
	if err != nil {
		log.Fatalf("failed creating uploadcare API client: %s", err)
	}

	return &uploadcare{
		secretKey: config.Uploadcare().SecKey,
		publicKey: config.Uploadcare().PubKey,
		url:       config.Uploadcare().URL,
		client:    client,
		ctx:       context.Background(),
		TaskType:  "delete_file",
		WG:        conf.WG,
		DataChan:  conf.DataChan,
		DoneChan:  conf.DoneChan,
		ErrorChan: conf.ErrorChan,
		conf:      conf,
	}
}

func (u *uploadcare) Listen() {
	for {
		select {
		case data := <-u.DataChan:
			go u.runTask(data)
		case err := <-u.ErrorChan:
			fmt.Println(err)
		case <-u.DoneChan:
			return
		}
	}
}
func (u *uploadcare) ExecuteTask(data any, taskType string) {
	u.WG.Add(1)
	u.DataChan <- data
	if taskType != "" {
		u.TaskType = taskType
	}
}
func (u *uploadcare) Done() {
	u.DoneChan <- true
}
func (u *uploadcare) CloseChannels() {
	close(u.DataChan)
	close(u.ErrorChan)
	close(u.DoneChan)
}

func (u *uploadcare) Disk(disk string) gateways.StorageService {
	if disk == drivers["uploadcare"] {
		return newUploadcare(u.conf)
	}
	return newLocal(u.conf)
}

func (u *uploadcare) UploadFile(dir string, f *multipart.FileHeader) (string, error) {
	file, err := f.Open()
	defer file.Close()
	if err != nil {
		return "", err
	}
	info, err := u.getFileInfo(file, dir)
	if err != nil {
		return "", err
	}
	uploadSvc := upload.NewService(u.client)
	params := upload.FileParams{
		Data:        file,
		Name:        info["name"],
		ContentType: info["mime"],
	}
	fID, err := uploadSvc.File(context.Background(), params)
	if err != nil {
		return "", err
	}

	fPath := fmt.Sprintf("%s/%s/", u.url, fID)
	return fPath, nil
}

func (u *uploadcare) UploadFiles(dir string, f []*multipart.FileHeader) ([]*string, error) {
	// TODO implement me
	panic("implement me")
}

func (u *uploadcare) runTask(data any) {
	defer u.WG.Done()
	switch u.TaskType {
	case "delete_file":
		if err := u.deleteFile(data); err != nil {
			u.ErrorChan <- err
		}
	}
}
func (u *uploadcare) deleteFile(path any) error {
	fPath := path.(string)
	fmt.Println(fPath)
	return nil
}

func (u *uploadcare) deleteFiles(path []*string) error {
	// TODO implement me
	panic("implement me")
}

func (u *uploadcare) getFileInfo(buffer multipart.File, prefix string) (map[string]string, error) {
	head := make([]byte, 512)
	_, err := buffer.Read(head)
	if err != nil {
		return nil, err
	}
	mtype := mimetype.Detect(head)
	filename := fmt.Sprintf("%s_%s%s", prefix, uuid.New(), mtype.Extension())

	return map[string]string{
		"name": filename,
		"mime": mtype.String(),
	}, nil
}
