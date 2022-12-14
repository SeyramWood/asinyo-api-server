package storage

import (
	"context"
	"mime/multipart"

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
}

func (l *local) Disk(disk string) gateways.StorageService {
	if disk == drivers["uploadcare"] {
		return newUploadcare()
	}
	return newLocal()
}

func (l *local) UploadFile(dir string, f *multipart.FileHeader) (string, error) {
	// TODO implement me
	panic("implement me local")
}

func (l *local) UploadFiles(dir string, f []*multipart.FileHeader) ([]*string, error) {
	// TODO implement me
	panic("implement me")
}

func (l *local) DeleteFile(path *string) error {
	// TODO implement me
	panic("implement me")
}

func (l *local) DeleteFiles(path []*string) error {
	// TODO implement me
	panic("implement me")
}

func newLocal() gateways.StorageService {
	return &local{
		secretKey: config.Uploadcare().SecKey,
		publicKey: config.Uploadcare().PubKey,
		url:       config.Uploadcare().URL,
		client:    nil,
		ctx:       context.Background(),
	}
}
