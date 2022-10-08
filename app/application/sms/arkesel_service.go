package sms

import (
	"bytes"
	"net/http"
	"sync"

	"github.com/goccy/go-json"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
)

type arkesel struct {
	APIKey string
	URL    string
	Sender string
}

func newArkeselService() gateways.SMSService {
	return &arkesel{
		APIKey: config.Arkesel().APIKey,
		URL:    config.Arkesel().URL,
		Sender: config.SMS().Sender,
	}
}

func (a *arkesel) Send(request *services.SMSPayload) (any, error) {

	var resp *http.Response
	var resError error

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(a *arkesel, request *services.SMSPayload) {
		defer wg.Done()
		data := services.ArkeselPayload{
			Sender:     a.Sender,
			Message:    request.Message,
			Recipients: request.Recipients,
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			resp, resError = nil, err
		}
		body := bytes.NewReader(payloadBytes)
		req, err := http.NewRequest("POST", a.URL, body)
		if err != nil {
			resp, resError = nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("api-key", a.APIKey)
		resp, resError = http.DefaultClient.Do(req)

	}(a, request)

	wg.Wait()

	if resError != nil {
		return nil, resError
	}
	resp.Body.Close()

	return resp, resError
}
