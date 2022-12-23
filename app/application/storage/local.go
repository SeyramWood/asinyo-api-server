package storage

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/uploadcare/uploadcare-go/ucare"

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
}

func newLocal() storageDriverService {
	return &local{
		secretKey: config.Uploadcare().SecKey,
		publicKey: config.Uploadcare().PubKey,
		url:       config.Uploadcare().URL,
		client:    nil,
		ctx:       context.Background(),
		TaskType:  "delete_file",
	}
}

func (l *local) uploadFile(dir string, f *multipart.FileHeader) (string, error) {
	// TODO implement me

	panic("implement me UploadFiles")
}

func (l *local) uploadFiles(dir string, files []*multipart.FileHeader) ([]string, error) {
	// TODO implement me

	panic("implement me UploadFiles")
}

func (l *local) deleteFile(path any) error {
	fPath := path.(string)
	fmt.Println("local", fPath)
	return nil
}

func (l *local) deleteFiles(paths []any) error {
	// TODO implement me
	panic("implement me")
}

func (l *local) getFileInfo(f *multipart.FileHeader, prefix string) (map[string]string, error) {
	buffer, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer buffer.Close()
	head := make([]byte, 512)
	_, err = buffer.Read(head)
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
