package payment

import (
	"bytes"
	"fmt"
	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/config"
	"github.com/goccy/go-json"
	"net/http"
)

type paystackService struct {
	repo   gateways.PaymentRepo
	URL    string
	secKey string
	pubKey string
	email  string
}

func NewPaystackService(repo gateways.PaymentRepo) gateways.PaymentService {
	return &paystackService{
		repo:   repo,
		URL:    config.Paystack().URL,
		secKey: config.Paystack().SecKey,
		pubKey: config.Paystack().PubKey,
		email:  config.Paystack().Email,
	}
}

func (p paystackService) Pay(model interface{}) (interface{}, error) {
	request := model.(models.OrderRequest)
	data := models.OrderRequest{
		Amount:   request.Amount,
		Email:    request.Email,
		Currency: request.Currency,
		MetaData: request.MetaData,
	}

	return p.initiateTransaction(data)
}

func (p paystackService) Verify(reference string) (interface{}, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/transaction/verify/%s", p.URL, reference), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.secKey))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (p paystackService) initiateTransaction(request models.OrderRequest) (*http.Response, error) {

	payloadBytes, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/transaction/initialize", p.URL), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.secKey))
	//req.Header.Set("Cache-Control", fmt.Sprintf("no-cache"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p paystackService) verifyTransaction(reference string) (*http.Response, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/transaction/verify/:%s", p.URL, reference), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.secKey))
	req.Header.Set("Cache-Control", fmt.Sprintf("no-cache"))

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp, nil
}
