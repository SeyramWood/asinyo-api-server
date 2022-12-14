package storage

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/uploadcare/uploadcare-go/ucare"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

type uploadcare struct {
	secretKey string
	publicKey string
	url       string
	client    ucare.Client
	ctx       context.Context
}

func newUploadcare() gateways.StorageService {
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
	}
}

func (u *uploadcare) Disk(disk string) gateways.StorageService {
	if disk == drivers["uploadcare"] {
		return newUploadcare()
	}
	return newLocal()
}

func (u *uploadcare) UploadFile(dir string, f *multipart.FileHeader) (string, error) {
	// TODO implement me
	panic("implement me uploadcare")
}

func (u *uploadcare) UploadFiles(dir string, f []*multipart.FileHeader) ([]*string, error) {
	// TODO implement me
	panic("implement me")
}

func (u *uploadcare) DeleteFile(path *string) error {
	// TODO implement me
	panic("implement me")
}

func (u *uploadcare) DeleteFiles(path []*string) error {
	// TODO implement me
	panic("implement me")
}
