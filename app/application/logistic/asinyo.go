package logistic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
	"github.com/goccy/go-json"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type asinyo struct {
	repo         gateways.LogisticRepo
	orderRepo    gateways.OrderRepo
	GoogleAPIKey string
	APIKey       string
	URL          string
	TaskType     string
	WG           *sync.WaitGroup
	DataChan     chan *ent.Order
	WebhookChan  chan any
	DoneChan     chan bool
	ErrorChan    chan error
	Dispatches   map[string]*gabs.Container
}

func newAsinyoService(conf *logistic, repo gateways.LogisticRepo) gateways.LogisticService {
	return &asinyo{
		repo:         repo,
		orderRepo:    nil,
		GoogleAPIKey: config.Google().APIKey,
		APIKey:       config.Tookan().APIKey,
		URL:          config.Tookan().URL,
		TaskType:     "create",
		WG:           conf.WG,
		DataChan:     conf.DataChan,
		WebhookChan:  conf.WebhookChan,
		DoneChan:     conf.DoneChan,
		ErrorChan:    conf.ErrorChan,
		Dispatches:   nil,
	}
}

func (a *asinyo) OrderRepo(repo gateways.OrderRepo) gateways.LogisticService {
	a.orderRepo = repo
	return a
}

func (a *asinyo) ExecuteTask(order *ent.Order, task ...any) {
	a.WG.Add(1)
	a.DataChan <- order
	if task != nil {
		a.TaskType = task[0].(string)
		a.Dispatches = task[1].(map[string]*gabs.Container)
	}
}

func (a *asinyo) ExecuteWebhook(response any) {
	a.WG.Add(1)
	a.WebhookChan <- response
}

func (a *asinyo) Listen() {
	for {
		select {
		case data := <-a.DataChan:
			go a.createTask(data)
		case response := <-a.WebhookChan:
			go a.processWebhookResponse(response)
		case err := <-a.ErrorChan:
			log.Println(err)
		case <-a.DoneChan:
			return
		}
	}
}

func (a *asinyo) Done() {
	a.DoneChan <- true
}

func (a *asinyo) CloseChannels() {
	close(a.DataChan)
	close(a.WebhookChan)
	close(a.ErrorChan)
	close(a.DoneChan)
}

func (a *asinyo) FareEstimate(coordinates *models.OrderFareEstimateRequest) (
	[]*services.FareEstimateResponseData, error,
) {
	var response []*services.FareEstimateResponseData
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for _, coordinate := range coordinates.Pickups {
		wg.Add(1)
		go func(pickup *services.Coordinate) {
			defer wg.Done()
			result, err := a.getFareEstimate(coordinates.Delivery, pickup)
			if err != nil {
				panic(fmt.Errorf("%s", err))
			}
			mu.Lock()
			response = append(response, result)
			mu.Unlock()
		}(coordinate)

	}
	wg.Wait()
	return response, nil
}

func (a *asinyo) createTask(order *ent.Order) {
	defer a.WG.Done()
	switch a.TaskType {
	case "create":
		if err := a.createDispatch(order); err != nil {
			a.ErrorChan <- err
		}
	case "update":
		if err := a.updateDispatch(order); err != nil {
			a.ErrorChan <- err
		}
	}
}

func (a *asinyo) createDispatch(order *ent.Order) error {
	var dispatchIds []int
	var dispatches []*models.AsinyoDispatchInfo
	err := json.Unmarshal(a.Dispatches["dispatches"].Bytes(), &dispatches)
	if err != nil {
		return err
	}
	err = json.Unmarshal(a.Dispatches["dispatchIds"].Bytes(), &dispatchIds)
	if err != nil {
		return err
	}
	_, err = a.repo.InsertResponse("Asinyo", order.OrderNumber, models.AsinyoDispatch{
		DispatchIds: dispatchIds,
		Dispatches:  dispatches,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *asinyo) updateDispatch(order *ent.Order) error {
	var dispatchIds []int
	var dispatches []*models.AsinyoDispatchInfo
	err := json.Unmarshal(a.Dispatches["dispatches"].Bytes(), &dispatches)
	if err != nil {
		return err
	}
	err = json.Unmarshal(a.Dispatches["dispatchIds"].Bytes(), &dispatchIds)
	if err != nil {
		return err
	}
	_, err = a.repo.UpdateResponse(int(a.Dispatches["logistic"].Data().(float64)), models.AsinyoDispatch{
		DispatchIds: dispatchIds,
		Dispatches:  dispatches,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *asinyo) processWebhookResponse(response any) {
	// TODO implement me
	panic("implement me")
}

func (a *asinyo) getFareEstimate(delivery, pickup *services.Coordinate) (*services.FareEstimateResponseData, error) {
	resData := services.FareEstimateRequest{
		TemplateName:      "order_delivery",
		PickupLongitude:   fmt.Sprintf("%f", pickup.Longitude),
		PickupLatitude:    fmt.Sprintf("%f", pickup.Latitude),
		APIKey:            a.APIKey,
		DeliveryLatitude:  fmt.Sprintf("%f", delivery.Latitude),
		DeliveryLongitude: fmt.Sprintf("%f", delivery.Longitude),
		FormulaType:       2,
		MapKeys: struct {
			MapPlanType  int    `json:"map_plan_type,omitempty"`
			GoogleAPIKey string `json:"google_api_key,omitempty"`
		}{
			MapPlanType:  1,
			GoogleAPIKey: a.GoogleAPIKey,
		},
	}
	payloadBytes, err := json.Marshal(resData)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(payloadBytes)

	req, reqerr := http.NewRequest("POST", fmt.Sprintf("%s/get_fare_estimate", a.URL), body)

	if reqerr != nil {
		return nil, reqerr
	}
	req.Header.Set("Content-Type", "application/json")

	res, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		return nil, reserr
	}
	defer res.Body.Close()
	if res.Status != "200 OK" {
		return nil, reserr
	}
	resp_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resBody, errr := gabs.ParseJSON(resp_body)
	if errr != nil {
		return nil, errr
	}
	formulas, err := resBody.Path("data.formula_fields").Children()
	if err != nil {
		return nil, err
	}
	var respFormulas []*services.FareEstimateResponseFormula
	for _, formula := range formulas {
		respFormulas = append(
			respFormulas, &services.FareEstimateResponseFormula{
				DisplayName: formula.Path("display_name").Data().(string),
				Key:         formula.Path("key").Data().(string),
				Type:        formula.Path("type").Data().(float64),
				Surge: func() string {
					if formula.Path("surge").Data() != nil {
						return formula.Path("surge").Data().(string)
					}
					return ""
				}(),
				MultiplyingValue: formula.Path("multiplying_value").Data().(string),
				Expression:       formula.Path("expression").Data().(string),
				Sum:              formula.Path("sum").Data().(float64),
			},
		)
	}
	response := &services.FareEstimateResponseData{
		Distance:      resBody.Path("data.distance").Data().(float64),
		Time:          resBody.Path("data.time").Data().(float64),
		Formula:       respFormulas,
		EstimatedFare: resBody.Path("data.estimated_fare").Data().(float64),
	}
	return response, nil
}
