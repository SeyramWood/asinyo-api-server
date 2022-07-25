package storage

import (
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type storage struct {
	disk string
}

var disks = map[string]string{
	"public": "public/storage",
	"local":  "storage",
}

func NewStorage() *storage {
	return &storage{}
}
func (s *storage) Disk(disk string) *storage {
	if _, ok := disks[strings.ToLower(disk)]; !ok {
		log.Fatalln(fmt.Sprintf("the [%s] disk is not available.", disk))
	}
	return &storage{disk}
}
func (s *storage) Exist(path string) bool {
	if _, err := os.Stat(filepath.Join(disks[s.disk], path)); err != nil {
		return false
	}
	return true

}
func (s *storage) Delete(path string) error {
	pth := filepath.Join(disks[s.disk], path)
	if _, err := os.Stat(pth); err != nil {
		return err
	}
	if err := os.Remove(pth); err != nil {
		return err
	}
	return nil
}
func (s *storage) DeleteAll(path string) error {
	pth := filepath.Join(disks[s.disk], path)
	if _, err := os.Stat(pth); err != nil {
		return err
	}
	if err := os.RemoveAll(pth); err != nil {
		return err
	}
	return nil
}
func (s *storage) MakeDirectory(path string) error {
	pth := filepath.Join(disks[s.disk], path)
	if err := os.MkdirAll(pth, 0755); err != nil {
		return err
	}
	return nil
}
func (s *storage) GetPath(path string) (string, error) {
	pth := filepath.Join(disks[s.disk], path)
	if !s.Exist(path) {
		return "", fmt.Errorf("the path [%s] dosn't exist", pth)
	}
	return pth, nil
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
