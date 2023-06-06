package services

type (
	MailerMessage struct {
		From        string
		FromName    string
		To          string
		Subject     string
		Attachments []string
		Data        any
		DataMap     map[string]any
		Template    string
	}
)
