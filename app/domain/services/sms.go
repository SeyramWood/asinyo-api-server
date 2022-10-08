package services

type (
	SMSPayload struct {
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}
	ArkeselPayload struct {
		Sender     string   `json:"sender"`
		Message    string   `json:"message"`
		Recipients []string `json:"recipients"`
	}
)
