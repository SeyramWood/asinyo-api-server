package storage

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"regexp"
	"sync"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/uploadcare/uploadcare-go/file"
	"github.com/uploadcare/uploadcare-go/ucare"
	"github.com/uploadcare/uploadcare-go/upload"

	"github.com/SeyramWood/config"
)

type uploadcare struct {
	secretKey string
	publicKey string
	url       string
	client    ucare.Client
	ctx       context.Context
	TaskType  string
}

func newUploadcare() storageDriverService {
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
	}
}

func (u *uploadcare) uploadFile(dir string, f *multipart.FileHeader) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()
	info, err := u.getFileInfo(f, dir)
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

func (u *uploadcare) uploadFiles(dir string, files []*multipart.FileHeader) ([]string, error) {
	var urls []string
	wg := sync.WaitGroup{}
	for _, f := range files {
		wg.Add(1)
		go func(f *multipart.FileHeader) {
			defer wg.Done()
			path, err := u.uploadFile(dir, f)
			if err != nil {
				panic(fmt.Sprintf("error uploading file [error]: %s", err))
			}
			urls = append(urls, path)
		}(f)
	}
	wg.Wait()
	return urls, nil
}

func (u *uploadcare) deleteFile(path any) error {
	fPath := path.(string)
	r, _ := regexp.Compile(`^(?:https://ucarecdn.com/)([0-9a-zA-Z]{8}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{12})(?:/)$`)
	matches := r.FindAllStringSubmatch(fPath, -1)[0]
	fileSvc := file.NewService(u.client)
	_, err := fileSvc.Delete(u.ctx, matches[1])
	if err != nil {
		return err
	}
	return nil
}

func (u *uploadcare) deleteFiles(paths []any) error {
	var ids []string
	for _, path := range paths {
		fPath := path.(string)
		r, _ := regexp.Compile(`^(?:https://ucarecdn.com/)([0-9a-zA-Z]{8}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{12})(?:/)$`)
		matches := r.FindAllStringSubmatch(fPath, -1)[0]
		ids = append(ids, matches[1])
	}
	fileSvc := file.NewService(u.client)
	_, err := fileSvc.BatchDelete(u.ctx, ids)
	if err != nil {
		return err
	}
	return nil
}

func (u *uploadcare) getFileInfo(f *multipart.FileHeader, prefix string) (map[string]string, error) {
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
