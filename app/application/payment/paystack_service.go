package payment

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/Jeffail/gabs"
	"github.com/goccy/go-json"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
)

type paystackService struct {
	repo   gateways.PaymentRepo
	URL    string
	secKey string
	pubKey string
	email  string
}

func newPaystackService(repo gateways.PaymentRepo) gateways.PaymentService {
	return &paystackService{
		repo:   repo,
		URL:    config.Paystack().URL,
		secKey: config.Paystack().SecKey,
		pubKey: config.Paystack().PubKey,
		email:  config.Paystack().Email,
	}
}

func (p *paystackService) Pay(request any) (any, error) {
	req := request.(services.PaystackPayload)

	return p.initiateTransaction(req)
}

func (p *paystackService) Verify(reference string) (interface{}, error) {

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

func (p *paystackService) FormatPayload(request any) (*models.OrderPayload, error) {
	var response *services.PaystackResponse
	resBody, err := gabs.ParseJSON(request.([]byte))
	if err != nil {
		return nil, err
	}
	response = &services.PaystackResponse{
		Event:     resBody.Path("event").Data().(string),
		Amount:    resBody.Path("data.amount").Data().(float64),
		Currency:  resBody.Path("data.currency").Data().(string),
		Channel:   resBody.Path("data.channel").Data().(string),
		Reference: resBody.Path("data.reference").Data().(string),
		PaidAt:    resBody.Path("data.paid_at").Data().(string),
		MetaData: &services.OrderResponseMetadata{
			User:           resBody.Path("data.metadata.user").Data().(string),
			UserType:       resBody.Path("data.metadata.userType").Data().(string),
			OrderNumber:    resBody.Path("data.metadata.orderNumber").Data().(string),
			Address:        resBody.Path("data.metadata.address").Data().(string),
			DeliveryMethod: resBody.Path("data.metadata.deliveryMethod").Data().(string),
			PaymentMethod:  resBody.Path("data.metadata.paymentMethod").Data().(string),
			DeliveryFee:    resBody.Path("data.metadata.deliveryFee").Data().(string),
			Pickup:         resBody.Path("data.metadata.pickup").Data().(string),
			Products: func() []*services.ProductDetails {
				var products []*services.ProductDetails
				children, _ := resBody.Path("data.metadata.products").Children()
				wg := sync.WaitGroup{}
				for _, child := range children {
					wg.Add(1)
					go func(child *gabs.Container) {
						defer wg.Done()
						pro := child.Data().(map[string]interface{})
						id, _ := strconv.Atoi(pro["id"].(string))
						store, _ := strconv.Atoi(pro["store"].(string))
						quantity, _ := strconv.Atoi(pro["quantity"].(string))
						price, _ := strconv.ParseFloat(pro["price"].(string), 64)
						promoPrice, _ := strconv.ParseFloat(pro["promoPrice"].(string), 64)
						products = append(
							products, &services.ProductDetails{
								ID:         id,
								Store:      store,
								Quantity:   quantity,
								Price:      price,
								PromoPrice: promoPrice,
							},
						)
					}(child)
				}
				wg.Wait()
				return products
			}(),
		},
	}
	if strings.Compare(response.Event, "charge.success") == 0 {
		userId, _ := strconv.Atoi(response.MetaData.User)
		addressId, _ := strconv.Atoi(response.MetaData.Address)
		pickupId, _ := strconv.Atoi(response.MetaData.Pickup)
		deliveryFee, _ := strconv.ParseFloat(response.MetaData.DeliveryFee, 32)
		data := &models.OrderPayload{
			Amount:    response.Amount,
			Reference: response.Reference,
			Currency:  response.Currency,
			Channel:   response.Channel,
			PaidAt:    response.PaidAt,
			Metadata: &models.OrderPayloadMetadata{
				User:           userId,
				Pickup:         pickupId,
				Address:        addressId,
				OrderNumber:    response.MetaData.OrderNumber,
				DeliveryFee:    deliveryFee,
				UserType:       response.MetaData.UserType,
				DeliveryMethod: response.MetaData.DeliveryMethod,
				PaymentMethod:  response.MetaData.PaymentMethod,
				Products:       response.MetaData.Products,
			},
		}
		return data, nil
	}
	return nil, fmt.Errorf("%s", "Unsuccessful payment")
}

func (p *paystackService) initiateTransaction(request services.PaystackPayload) (*http.Response, error) {

	payloadBytes, err := json.Marshal(request)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/transaction/initialize", p.URL), body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.secKey))
	// req.Header.Set("Cache-Control", fmt.Sprintf("no-cache"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *paystackService) verifyTransaction(reference string) (*http.Response, error) {

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
