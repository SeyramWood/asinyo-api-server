package storage

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
)

type storage struct {
}

func NewStorage() *storage {
	return &storage{}
}

func (s *storage) SaveFile(c *fiber.Ctx, field, directory string) (interface{}, error) {
	file, _ := c.FormFile(field)
	buffer, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer buffer.Close()
	head := make([]byte, 512)
	buffer.Read(head)

	mtype := mimetype.Detect(head)

	filename := fmt.Sprintf("asinyo_%s%s", uuid.New(), mtype.Extension())

	dir := fmt.Sprintf("public/%s", directory)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Something went wrong",
		})
	}
	if err := c.SaveFile(file, fmt.Sprintf("./%s/%s", dir, filename)); err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Something went wrong",
		})
	}

	return fmt.Sprintf("%s/%s/%s", c.BaseURL(), directory, filename), nil

}

func (s *storage) SaveFiles(c *fiber.Ctx, field, directory string) (interface{}, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}

	files := form.File[field]

	dir := fmt.Sprintf("public/%s", directory)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": "Something went wrong",
		})
	}

	urls := []string{}

	for _, file := range files {
		buffer, err := file.Open()
		if err != nil {
			return nil, err
		}
		head := make([]byte, 512)
		buffer.Read(head)
		buffer.Close()

		mtype := mimetype.Detect(head)

		filename := fmt.Sprintf("asinyo_%s%s", uuid.New(), mtype.Extension())

		if err := c.SaveFile(file, fmt.Sprintf("./%s/%s", dir, filename)); err != nil {
			return nil, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"msg": "Something went wrong",
			})
		}

		urls = append(urls, fmt.Sprintf("%s/%s/%s", c.BaseURL(), directory, filename))

	}

	return urls, nil

}
