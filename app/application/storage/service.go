package storage

import (
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/config"
)

var drivers = map[string]string{
	"local":      "local",
	"uploadcare": "uploadcare",
}

func NewStorageService() gateways.StorageService {
	switch config.App().FilesystemDriver {
	case "uploadcare":
		return newUploadcare()
	default:
		return newLocal()
	}
}
