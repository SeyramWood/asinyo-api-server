package sms

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/goccy/go-json"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
)

type arkesel struct {
	APIKey         string
	URL            string
	Sender         string
	FailedCount    int64
	FailedDelay    time.Duration
	WG             *sync.WaitGroup
	DataChan       chan any
	FailedDataChan chan any
	DoneChan       chan bool
	ErrorChan      chan error
}

func newArkeselService(sms *sms) gateways.SMSService {
	return &arkesel{
		APIKey:         config.Arkesel().APIKey,
		URL:            config.Arkesel().URL,
		Sender:         config.SMS().Sender,
		FailedCount:    0,
		FailedDelay:    60 * time.Second,
		WG:             sms.WG,
		DataChan:       sms.DataChan,
		FailedDataChan: sms.FailedDataChan,
		DoneChan:       nil,
		ErrorChan:      nil,
	}
}

func (a *arkesel) Listen() {
	for {
		select {
		case msg := <-a.DataChan:
			go a.sendSMS(msg)
		case msg := <-a.FailedDataChan:
			if a.FailedCount <= 3 {
				a.FailedDelay = time.Duration(a.FailedCount) * a.FailedDelay
				time.Sleep(a.FailedDelay)
				go a.sendSMS(msg)
			}
		case err := <-a.ErrorChan:
			fmt.Println(err)
		case <-a.DoneChan:
			return
		}
	}
}
func (a *arkesel) Send(msg *services.SMSPayload) {
	a.WG.Add(1)
	a.DataChan <- msg
}
func (a *arkesel) Done() {
	a.DoneChan <- true
}
func (a *arkesel) CloseChannels() {
	close(a.DataChan)
	close(a.ErrorChan)
	close(a.DoneChan)
}

func (a *arkesel) sendSMS(request any) {
	if req, ok := request.(*services.SMSPayload); ok {
		data := services.ArkeselPayload{
			Sender:     a.Sender,
			Message:    req.Message,
			Recipients: a.formatRecipients(req.Recipients),
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			a.ErrorChan <- err
		}
		body := bytes.NewReader(payloadBytes)
		req, err := http.NewRequest("POST", a.URL, body)
		if err != nil {
			a.ErrorChan <- err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("api-key", a.APIKey)
		resp, resError := http.DefaultClient.Do(req)
		if resError != nil {
			a.ErrorChan <- resError
		}
		resp.Body.Close()
		if resp.StatusCode != 200 {
			a.FailedCount++
			a.FailedDataChan <- req
		}

	} else {
		fmt.Println("Error casting SMS payload")
	}

}
func (a *arkesel) formatRecipients(recipients []string) []string {
	var reps []string
	for _, recipient := range recipients {
		reps = append(reps, fmt.Sprintf("233%s", strings.Join(strings.Split(recipient, "")[1:], "")))
	}
	return reps
}
