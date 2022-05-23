package storage

import (
	"fmt"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
)

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) fiberStorage(c *fiber.Ctx, directory string) error {
	file, _ := c.FormFile("image")
	buffer, err := file.Open()
	if err != nil {
		return err
	}
	defer buffer.Close()
	head := make([]byte, 512)
	buffer.Read(head)
	mtype := mimetype.Detect(head)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), mtype.Extension())
	return c.SaveFile(file, fmt.Sprintf("./storage/app/%s/%s", directory, fileName))
}
