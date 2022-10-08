package application

import (
	"crypto/rand"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/SeyramWood/ent"
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
	// Send SMS if username is phone
	// Send Email if username is email
	return string(b), nil
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
