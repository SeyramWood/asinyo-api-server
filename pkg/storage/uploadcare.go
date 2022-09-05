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

	"github.com/SeyramWood/config"
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
	client, err := ucare.NewClient(
		ucare.APICreds{
			SecretKey: c.secretKey,
			PublicKey: c.publicKey,
		}, &ucare.Config{
			SignBasedAuthentication: true,
			APIVersion:              ucare.APIv06,
		},
	)
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

func (c *uploadcare) UploadMerchantStore(file *multipart.FileHeader, form *multipart.Form) (string, []string, error) {
	type (
		logo struct {
			FP  string
			Err error
		}
		images struct {
			FP  []string
			Err error
		}
	)
	res := map[string]interface{}{
		"logo":   logo{},
		"images": images{},
	}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	for key := range res {
		wg.Add(1)
		go func(key string, file *multipart.FileHeader, form *multipart.Form) {
			defer wg.Done()
			if key == "logo" {
				fPath, err := c.Upload(file, "merchant_logo")
				if err != nil {
					mut.Lock()
					res[key] = logo{
						FP:  "",
						Err: err,
					}
					mut.Unlock()
				} else {
					mut.Lock()
					res[key] = logo{
						FP:  fPath,
						Err: nil,
					}
					mut.Unlock()
				}

			}
			if key == "images" {
				fPaths, err := c.Uploads(form.File["otherImages"], "merchant_store")
				if err != nil {
					mut.Lock()
					res[key] = images{
						FP:  []string{},
						Err: err,
					}
					mut.Unlock()
				} else {
					mut.Lock()
					res[key] = images{
						FP:  fPaths,
						Err: nil,
					}
					mut.Unlock()
				}

			}
		}(key, file, form)
	}
	wg.Wait()

	resLogo, _ := res["logo"].(logo)
	resImages, _ := res["images"].(images)

	if resLogo.Err != nil || resImages.Err != nil {
		return "", []string{}, fmt.Errorf("upload error")
	}

	return resLogo.FP, resImages.FP, nil
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
