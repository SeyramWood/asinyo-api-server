package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
)

type mailer struct {
	Mailer      string
	FromAddress string
	FromName    string
	WG          *sync.WaitGroup
	MailerChan  chan *services.Message
	DoneChan    chan bool
	ErrorChan   chan error
}

func NewEmail(app *config.MailServer) gateways.EmailService {
	wg := app.WG
	return &mailer{
		Mailer:      config.Mailer().Mailer,
		FromAddress: config.Mailer().FromAddress,
		FromName:    config.Mailer().FromName,
		WG:          wg,
		MailerChan:  app.MailerChan,
		DoneChan:    app.DoneChan,
		ErrorChan:   app.ErrorChan,
	}
}

func (m *mailer) Listen() {
	for {
		select {
		case msg := <-m.MailerChan:
			go m.sendMail(msg, m.ErrorChan)
		case err := <-m.ErrorChan:
			fmt.Println(err)
		case <-m.DoneChan:
			return
		}
	}
}
func (m *mailer) Send(msg *services.Message) {
	m.WG.Add(1)
	m.MailerChan <- msg
}
func (m *mailer) Done() {
	m.DoneChan <- true
}
func (m *mailer) CloseChannels() {
	close(m.MailerChan)
	close(m.ErrorChan)
	close(m.DoneChan)
}

func (m *mailer) sendMail(msg *services.Message, errorChan chan error) {
	defer m.WG.Done()
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}
	msg.DataMap = data
	// build html mail
	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		fmt.Println(err)
		errorChan <- err
	}

	// build plain text mail
	plainMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		fmt.Println(err)
		errorChan <- err
	}

	smtpClient, err := config.SMTPServer().Connect()
	if err != nil {
		fmt.Println(err)
		errorChan <- err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)

	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		fmt.Println(err)
		errorChan <- err
	}
}

func (m *mailer) buildHTMLMessage(msg *services.Message) (string, error) {
	templateToRender := fmt.Sprintf("./app/framework/web/template/emails/%s.html.gohtml", msg.Template)

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *mailer) buildPlainTextMessage(msg *services.Message) (string, error) {
	templateToRender := fmt.Sprintf("./app/framework/web/template/emails/%s.plain.gohtml", msg.Template)

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *mailer) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}
