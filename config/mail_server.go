package config

import (
	"strconv"
	"sync"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/SeyramWood/app/domain/services"
)

type MailServer struct {
	SMTP       *mail.SMTPServer
	WG         *sync.WaitGroup
	MailerChan chan *services.Message
	DoneChan   chan bool
	ErrorChan  chan error
}

func SMTPServer() *mail.SMTPServer {

	port, _ := strconv.Atoi(Mailer().Port)

	server := mail.NewSMTPClient()

	server.Host = Mailer().Host
	server.Port = port
	server.Username = Mailer().Username
	server.Password = Mailer().Password
	server.Encryption = getEncryption(Mailer().Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	return server
}

func getEncryption(e string) mail.Encryption {
	switch e {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
