package application

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unsafe"

	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/admin"
)

func Rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func GenerateOTP(max int) (string, error) {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}
func GeneratePassword(max int) (string, error) {
	var table = [...]byte{
		'1', '2', '3', '4', '5', '6', '7', '8', '9', '0', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l',
		'm', 'n', 'o', 'p', 'g', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		return "", err
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b), nil
}

func RandomString(size int) string {
	chars := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, size)
	_, _ = rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

func UsernameType(username, delimiter string) bool {
	if strings.Contains(username, "@") && delimiter == "email" {
		return true
	}
	phone, _ := regexp.Compile(`^0\d{9}$`)
	if phone.MatchString(username) && delimiter == "phone" {
		return true
	}
	return false
}

func ReadAdminPermissions(db *ent.Client, id int) ([]string, error) {
	results, err := db.Admin.Query().Where(admin.ID(id)).QueryRoles().QueryPermissions().All(context.Background())
	if err != nil {
		return nil, err
	}
	var permissions []string
	for _, result := range results {
		permissions = append(permissions, result.Slug)
	}
	return permissions, nil

}
