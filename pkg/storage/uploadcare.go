package storage

import (
	"context"
	"fmt"
	"github.com/SeyramWood/config"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/uploadcare/uploadcare-go/ucare"
	"github.com/uploadcare/uploadcare-go/upload"
	"log"
	"mime/multipart"
	"sync"
)

type uploadcare struct {
	secretKey string
	publicKey string
	url       string
	client    ucare.Client
}

func NewUploadCare() *uploadcare {
	return &uploadcare{
		secretKey: config.Uploadcare().SecKey,
		publicKey: config.Uploadcare().PubKey,
		url:       config.Uploadcare().URL,
		client:    nil,
	}
}

func (c *uploadcare) Client() *uploadcare {
	client, err := ucare.NewClient(ucare.APICreds{
		SecretKey: c.secretKey,
		PublicKey: c.publicKey,
	}, &ucare.Config{
		SignBasedAuthentication: true,
		APIVersion:              ucare.APIv06,
	})
	if err != nil {
		log.Fatalf("creating uploadcare API client: %s", err)
	}
	c.client = client
	return c
}

func (c *uploadcare) Upload(f *multipart.FileHeader, prefix string) (string, error) {
	info, err := c.getFileInfo(f, prefix)
	if err != nil {
		return "", err
	}
	file, err := f.Open()
	defer file.Close()
	if err != nil {
		return "", err
	}

	uploadSvc := upload.NewService(c.client)
	params := upload.FileParams{
		Data:        file,
		Name:        info["name"],
		ContentType: info["mime"],
	}
	fID, err := uploadSvc.File(context.Background(), params)

	if err != nil {
		return "", err
	}
	fPath := fmt.Sprintf("%s/%s/", c.url, fID)
	return fPath, nil
}

func (c *uploadcare) Uploads(files []*multipart.FileHeader, prefix string) ([]string, error) {
	var urls []string
	wg := sync.WaitGroup{}
	for _, f := range files {
		wg.Add(1)
		go func(f *multipart.FileHeader) {
			defer wg.Done()
			path, err := c.Upload(f, prefix)
			if err != nil {
				log.Fatalln(fmt.Sprintf("error uploading file [error]: %s", err))
			}
			urls = append(urls, path)
		}(f)
	}
	wg.Wait()
	return urls, nil
}

func (c *uploadcare) getFileInfo(f *multipart.FileHeader, prefix string) (map[string]string, error) {
	buffer, err := f.Open()
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
