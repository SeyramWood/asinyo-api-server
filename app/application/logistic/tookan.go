package logistic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Jeffail/gabs"
	"github.com/goccy/go-json"
	"github.com/samber/lo"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/config"
	"github.com/SeyramWood/ent"
)

type tookan struct {
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
	StoreIds     []int
}

func newTookanService(conf *logistic, repo gateways.LogisticRepo) gateways.LogisticService {
	return &tookan{
		repo:         repo,
		orderRepo:    nil,
		GoogleAPIKey: config.Google().APIKey,
		APIKey:       config.Tookan().APIKey,
		URL:          config.Tookan().URL,
		TaskType:     "create_multiple_tasks",
		WG:           conf.WG,
		DataChan:     conf.DataChan,
		WebhookChan:  conf.WebhookChan,
		DoneChan:     conf.DoneChan,
		ErrorChan:    conf.ErrorChan,
	}
}

func (t *tookan) New(repo gateways.OrderRepo) gateways.LogisticService {
	t.orderRepo = repo
	return t
}

func (t *tookan) ExecuteTask(order *ent.Order, deliveryType string) {
	t.WG.Add(1)
	t.DataChan <- order
	if deliveryType != "" {
		t.TaskType = deliveryType
	}
}

func (t *tookan) ExecuteWebhook(response any) {
	t.WG.Add(1)
	t.WebhookChan <- response
}

func (t *tookan) Listen() {
	for {
		select {
		case data := <-t.DataChan:
			go t.createTask(data)
		case response := <-t.WebhookChan:
			go t.processWebhookResponse(response)
		case err := <-t.ErrorChan:
			fmt.Println(err)
		case <-t.DoneChan:
			return
		}
	}
}

func (t *tookan) Done() {
	t.DoneChan <- true
}

func (t *tookan) CloseChannels() {
	close(t.DataChan)
	close(t.WebhookChan)
	close(t.ErrorChan)
	close(t.DoneChan)
}

func (t *tookan) FareEstimate(coordinates *models.OrderFareEstimateRequest) (
	[]*services.FareEstimateResponseData, error,
) {
	var response []*services.FareEstimateResponseData
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	for _, coordinate := range coordinates.Pickups {
		wg.Add(1)
		go func(pickup *services.Coordinate) {
			defer wg.Done()
			result, err := t.getFareEstimate(coordinates.Delivery, pickup)
			if err != nil {
				fmt.Println(err)
			}
			mu.Lock()
			response = append(response, result)
			mu.Unlock()

		}(coordinate)

	}
	wg.Wait()
	return response, nil
}

func (t *tookan) createTask(order *ent.Order) {
	defer t.WG.Done()
	switch t.TaskType {
	case "pickup_delivery_tasks":
		if err := t.createPickupAndDeliveryTask(order); err != nil {
			t.ErrorChan <- err
		}
	case "edit_multiple_tasks":
		if err := t.updatePickupDeliveryTasks(order); err != nil {
			t.ErrorChan <- err
		}
	}
}
func (t *tookan) processWebhookResponse(response any) {
	defer t.WG.Done()
	res, err := t.formatWebhookPayload(response)
	if err != nil {
		t.ErrorChan <- err
	}
	fmt.Println(res)
	switch res.JobStatus {
	case 1:
		// Job Started
		if err := t.repo.UpdateOrderStatus(res.JobToken, "dispatched"); err != nil {
			t.ErrorChan <- err
		}
	case 2:
		// Job Successful
		if err := t.repo.UpdateOrderStatus(res.JobToken, "delivered"); err != nil {
			t.ErrorChan <- err
		}
	}
}

func (t *tookan) createPickupAndDeliveryTask(order *ent.Order) error {
	tasks, stores := t.formatPickupAndDelivery(order)
	if len(tasks) == 0 {
		return nil
	}

	for index, task := range tasks {
		payloadBytes, err := json.Marshal(&task)
		if err != nil {
			return err
		}

		body := bytes.NewReader(payloadBytes)
		req, reqerr := http.NewRequest("POST", fmt.Sprintf("%s/create_task", t.URL), body)
		if reqerr != nil {
			return reqerr
		}
		// req.Header.Set("Cache-Control", fmt.Sprintf("no-cache"))
		req.Header.Set("Content-Type", "application/json")

		res, reserr := http.DefaultClient.Do(req)
		if reserr != nil {
			return reserr
		}
		var response services.TookanPickupAndDeliveryResponse
		derr := json.NewDecoder(res.Body).Decode(&response)
		if derr != nil {
			return derr
		}
		defer res.Body.Close()
		if lo.Contains([]int{100, 101, 201, 404}, response.Status) {
			return fmt.Errorf("(%d) %s", response.Status, response.Message)
		}
		err = t.repo.UpdateOrderDeliveryTask(order.OrderNumber, stores[index])
		if err != nil {
			return err
		}
		resData := &models.TookanPickupAndDeliveryTaskResponse{
			PickupTrackingLink:  response.Data.PickupTrackingLink,
			DeliveryTracingLink: response.Data.DeliveryTracingLink,
			JobID:               response.Data.JobID,
			JobToken:            response.Data.JobToken,
			PickupHash:          response.Data.PickupHash,
			DeliveryHash:        response.Data.DeliveryHash,
			CustomerName:        response.Data.CustomerName,
			CustomerAddress:     response.Data.CustomerAddress,
			GeofenceDetails:     response.Data.GeofenceDetails,
		}
		_, err = t.repo.InsertResponse(order.OrderNumber, stores[index], resData)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t *tookan) updatePickupDeliveryTasks(order *ent.Order) error {
	task := &services.TookanPickupDeliveryUpdateTask{
		APIKey:     t.APIKey,
		Pickups:    t.formatPickups(order),
		Deliveries: t.formatDeliveries(order),
	}
	// log.Fatalln("UPDATE")
	payloadBytes, err := json.Marshal(&task)

	if err != nil {
		return err
	}

	body := bytes.NewReader(payloadBytes)

	req, reqerr := http.NewRequest("POST", fmt.Sprintf("%s/edit_multiple_tasks", t.URL), body)

	if reqerr != nil {
		return reqerr
	}

	req.Header.Set("Content-Type", "application/json")

	res, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		return reserr

	}

	var response services.TookanPickupDeliveryResponse
	derr := json.NewDecoder(res.Body).Decode(&response)

	if derr != nil {
		return derr
	}

	// fmt.Println(response)
	if response.Status == 200 {
		// resData := &models.TookanMultiTaskResponse{
		// 	Pickups:    response.Data.Pickups,
		// 	Deliveries: response.Data.Deliveries,
		// 	Geofence:   response.Data.GeofenceDetails,
		// }
		// _, err := t.repo.UpdateResponse(resData)
		// if err != nil {
		// 	t.ErrorChan <- err
		// }
	}

	return nil
}

func (t *tookan) formatMetadata(data *ent.Order, storeId int) []*services.TookanMetadata {
	type statusCheck struct {
		details []*ent.OrderDetail
	}
	var orderStatus statusCheck
	var amount float64

	for _, detail := range data.Edges.Details {
		if lo.Contains(data.StoreTasksCreated, detail.Edges.Store.ID) {
			continue
		}
		if detail.Edges.Store.ID == storeId {
			amount = amount + detail.Amount
		}
		orderStatus = statusCheck{
			details: append(orderStatus.details, detail),
		}
	}
	return []*services.TookanMetadata{
		{
			Label: "Reference",
			Data:  data.OrderNumber,
		},
		{
			Label: "Currency",
			Data:  data.Currency,
		},
		{
			Label: "DeliveryFee",
			Data:  data.DeliveryFee,
		},
		{
			Label: "Amount",
			Data:  amount,
		},
		{
			Label: "DeliveryMethod",
			Data:  data.DeliveryMethod,
		},
		{
			Label: "PaymentMethod",
			Data:  data.PaymentMethod,
		},
		{
			Label: "Status",
			Data:  data.Status,
		},
		{
			Label: "Address",
			Data: fmt.Sprintf(
				"%s %s\n%s\n%s,\n%s %s\n%s-%s\n%s", data.Edges.Address.OtherName, data.Edges.Address.LastName,
				data.Edges.Address.Address, data.Edges.Address.City, data.Edges.Address.StreetName,
				data.Edges.Address.District, data.Edges.Address.Region, data.Edges.Address.Country,
				data.Edges.Address.Phone,
			),
		},
		{
			Label: "Products",
			Data:  t.formatOrderDetails(orderStatus.details, storeId),
		},
	}
}

func (t *tookan) formatPickupMetadata(
	data []*ent.OrderDetail, orderNum string, storeId int,
) []*services.TookanMetadata {

	return []*services.TookanMetadata{
		{
			Label: "Reference",
			Data:  orderNum,
		},
		{
			Label: "Status",
			Data:  t.generatePickupStatus(data),
		},
		{
			Label: "Address",
			Data:  t.getMerchantAddress(data, storeId),
		},
		{
			Label: "Products",
			Data:  t.formatOrderDetails(data, storeId),
		},
	}
}

func (t *tookan) formatPickupAndDelivery(order *ent.Order) ([]*services.TookanPickupAndDeliveryTask, []int) {

	var response []*services.TookanPickupAndDeliveryTask
	var storeIds []int

	pickupTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 0, 0, 0, time.UTC)
	deliveryTime := time.Date(time.Now().Year(), time.Now().Month(), (time.Now().Day() + 3), 23, 0, 0, 0, time.UTC)

	for _, detail := range order.Edges.Details {
		if lo.Contains(order.StoreTasksCreated, detail.Edges.Store.ID) {
			continue
		}
		if !lo.Contains(storeIds, detail.Edges.Store.ID) && detail.Status == "processing" {
			tasks := func() *services.TookanPickupAndDeliveryTask {
				if s, err := detail.Edges.Store.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
					return &services.TookanPickupAndDeliveryTask{
						APIKey:         t.APIKey,
						OrderID:        order.OrderNumber,
						TeamID:         "",
						AutoAssignment: "0",
						JobDescription: "Order Pickup and Delivery",
						JobPickupPhone: s.Phone,
						JobPickupName:  fmt.Sprintf("%s %s", s.OtherName, s.LastName),
						JobPickupEmail: "",
						JobPickupAddress: fmt.Sprintf(
							"%s\n%s, %s\n%s-%s", detail.Edges.Store.Address.City,
							detail.Edges.Store.Address.StreetName,
							detail.Edges.Store.Address.District, detail.Edges.Store.Address.Region,
							detail.Edges.Store.Address.Country,
						),
						JobPickupLatitude:  "",
						JobPickupLongitude: "",
						JobPickupDatetime:  pickupTime.Format("01-02-2006 15:04"),
						CustomerEmail:      "",
						CustomerUsername: fmt.Sprintf(
							"%s %s", order.Edges.Address.OtherName, order.Edges.Address.LastName,
						),
						CustomerPhone: fmt.Sprintf("%s", order.Edges.Address.Phone),
						CustomerAddress: fmt.Sprintf(
							"%s\n%s, %s\n%s-%s", order.Edges.Address.City, order.Edges.Address.StreetName,
							order.Edges.Address.District,
							order.Edges.Address.Region, order.Edges.Address.Country,
						),
						Latitude:                  "",
						Longitude:                 "",
						JobDeliveryDatetime:       deliveryTime.Format("01-02-2006 15:04"),
						HasPickup:                 "1",
						HasDelivery:               "1",
						LayoutType:                "0",
						TrackingLink:              1,
						Timezone:                  "-330",
						CustomFieldTemplate:       "order_delivery",
						MetaData:                  t.formatMetadata(order, detail.Edges.Store.ID),
						PickupCustomFieldTemplate: "order_pickup",
						PickupMetaData: t.formatPickupMetadata(
							order.Edges.Details, order.OrderNumber, detail.Edges.Store.ID,
						),
						FleetID:    "",
						PRefImages: nil,
						RefImages:  nil,
						Notify:     0,
						Tags:       "pickup, delivery, order",
						Geofence:   0,
						RideType:   0,
					}
				}
				if r, err := detail.Edges.Store.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
					return &services.TookanPickupAndDeliveryTask{
						APIKey:         t.APIKey,
						OrderID:        order.OrderNumber,
						TeamID:         "",
						AutoAssignment: "0",
						JobDescription: "Order Pickup and Delivery",
						JobPickupPhone: r.Phone,
						JobPickupName:  fmt.Sprintf("%s %s", r.OtherName, r.LastName),
						JobPickupEmail: "",
						JobPickupAddress: fmt.Sprintf(
							"%s\n%s, %s\n%s-%s", detail.Edges.Store.Address.City,
							detail.Edges.Store.Address.StreetName,
							detail.Edges.Store.Address.District, detail.Edges.Store.Address.Region,
							detail.Edges.Store.Address.Country,
						),
						JobPickupLatitude:  "",
						JobPickupLongitude: "",
						JobPickupDatetime:  pickupTime.Format("01-02-2006 15:04"),
						CustomerEmail:      "",
						CustomerUsername: fmt.Sprintf(
							"%s %s", order.Edges.Address.OtherName, order.Edges.Address.LastName,
						),
						CustomerPhone: fmt.Sprintf("%s", order.Edges.Address.Phone),
						CustomerAddress: fmt.Sprintf(
							"%s\n%s, %s\n%s-%s", order.Edges.Address.City, order.Edges.Address.StreetName,
							order.Edges.Address.District,
							order.Edges.Address.Region, order.Edges.Address.Country,
						),
						Latitude:                  "",
						Longitude:                 "",
						JobDeliveryDatetime:       deliveryTime.Format("01-02-2006 15:04"),
						HasPickup:                 "1",
						HasDelivery:               "1",
						LayoutType:                "0",
						TrackingLink:              1,
						Timezone:                  "-330",
						CustomFieldTemplate:       "order_delivery",
						MetaData:                  t.formatMetadata(order, detail.Edges.Store.ID),
						PickupCustomFieldTemplate: "order_pickup",
						PickupMetaData: t.formatPickupMetadata(
							order.Edges.Details, order.OrderNumber, detail.Edges.Store.ID,
						),
						FleetID:    "",
						PRefImages: nil,
						RefImages:  nil,
						Notify:     0,
						Tags:       "pickup, delivery, order",
						Geofence:   0,
						RideType:   0,
					}
				}
				return nil
			}()

			response = append(response, tasks)

			storeIds = append(storeIds, detail.Edges.Store.ID)
		}
	}
	return response, storeIds
}

func (t *tookan) formatPickups(order *ent.Order) []*services.TookanPickupDelivery {

	currentTime := time.Date(2022, time.December, 25, 23, 0, 0, 0, time.UTC)
	var response []*services.TookanPickupDelivery

	for _, detail := range order.Edges.Details {
		if lo.Contains(order.StoreTasksCreated, detail.Edges.Store.ID) {
			continue
		}
		if !lo.Contains(t.StoreIds, detail.Edges.Store.ID) && detail.Status == "processing" {
			formattedPickup := func() *services.TookanPickupDelivery {
				if s, err := detail.Edges.Store.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
					return &services.TookanPickupDelivery{
						Address: fmt.Sprintf(
							"%s %s\n%s, %s.", detail.Edges.Store.Address.StreetName,
							detail.Edges.Store.Address.StreetName,
							detail.Edges.Store.Address.City, detail.Edges.Store.Address.Region,
						),
						Latitude:       0,
						Longitude:      0,
						Time:           currentTime.Format("01-02-2006 15:04"),
						Phone:          s.Phone,
						JobDescription: "Order Pickup",
						TemplateName:   "order_pickup",
						TemplateData: t.formatPickupMetadata(
							order.Edges.Details, order.OrderNumber, detail.Edges.Store.ID,
						),
						RefImages: nil,
						Name:      fmt.Sprintf("%s %s", s.OtherName, s.LastName),
						Email:     "",
						OrderID:   order.OrderNumber,
					}
				}
				if r, err := detail.Edges.Store.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
					return &services.TookanPickupDelivery{
						Address: fmt.Sprintf(
							"%s %s\n%s, %s.", detail.Edges.Store.Address.StreetName,
							detail.Edges.Store.Address.StreetName,
							detail.Edges.Store.Address.City, detail.Edges.Store.Address.Region,
						),
						Latitude:       0,
						Longitude:      0,
						Time:           currentTime.Format("01-02-2006 15:04"),
						Phone:          r.Phone,
						JobDescription: "Order Pickup",
						TemplateName:   "order_pickup",
						TemplateData: t.formatPickupMetadata(
							order.Edges.Details, order.OrderNumber, detail.Edges.Store.ID,
						),
						RefImages: nil,
						Name:      fmt.Sprintf("%s %s", r.OtherName, r.LastName),
						Email:     "",
						OrderID:   order.OrderNumber,
					}
				}
				return nil
			}()
			response = append(response, formattedPickup)
			t.StoreIds = append(t.StoreIds, detail.Edges.Store.ID)
		}
	}

	return response
}

func (t *tookan) formatDeliveries(order *ent.Order) []*services.TookanPickupDelivery {
	currentTime := time.Date(2022, time.December, 25, 23, 0, 0, 0, time.UTC)
	var response []*services.TookanPickupDelivery
	var storeIds []int

	for _, detail := range order.Edges.Details {
		if lo.Contains(order.StoreTasksCreated, detail.Edges.Store.ID) {
			continue
		}
		if !lo.Contains(storeIds, detail.Edges.Store.ID) && detail.Status == "processing" {
			response = append(
				response, &services.TookanPickupDelivery{

					Address: fmt.Sprintf(
						"%s\n%s, %s\n%s-%s",
						order.Edges.Address.City, order.Edges.Address.StreetName, order.Edges.Address.District,
						order.Edges.Address.Region, order.Edges.Address.Country,
					),
					Latitude:       0,
					Longitude:      0,
					Time:           currentTime.Format("01-02-2006 15:04"),
					Phone:          order.Edges.Address.Phone,
					JobDescription: "Order Delivery",
					TemplateName:   "order_delivery",
					TemplateData:   t.formatMetadata(order, detail.Edges.Store.ID),
					RefImages:      nil,
					Name:           fmt.Sprintf("%s %s", order.Edges.Address.OtherName, order.Edges.Address.LastName),
					Email:          "",
					OrderID:        order.OrderNumber,
				},
			)

			storeIds = append(storeIds, detail.Edges.Store.ID)
		}
		break
	}

	return response
}

func (t *tookan) formatOrderDetails(data []*ent.OrderDetail, storeId int) [][]any {
	var response [][]any
	for _, v := range data {

		if storeId != 0 && v.Edges.Store.ID != storeId {
			continue
		}
		formattedSlice := append(
			[]any{
				v.Edges.Product.Name,
				v.Edges.Product.Unit,
				v.Edges.Product.Weight,
				v.Edges.Product.Image,
				v.Quantity,
				v.Price,
				v.PromoPrice,
				v.Amount,
				v.Edges.Store.Name,
			}, t.formatStoreDetails(v.Edges.Store)...,
		)
		response = append(response, formattedSlice)
	}

	return response
}

func (t *tookan) formatStoreDetails(data *ent.MerchantStore) []any {
	if r, err := data.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
		return []any{
			r.Phone,
			data.Address.StreetName,
			fmt.Sprintf("%s\n%s, %s", data.Address.City, data.Address.District, data.Address.Region),
		}
	}
	if r, err := data.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
		return []any{
			r.Phone,
			data.Address.StreetName,
			fmt.Sprintf("%s\n%s, %s", data.Address.City, data.Address.District, data.Address.Region),
		}
	}
	return []any{}
}

func (t *tookan) getMerchantAddress(data []*ent.OrderDetail, storeId int) string {
	var response string
	for _, detail := range data {
		if storeId != 0 && detail.Edges.Store.ID == storeId {
			response = func() string {
				if s, err := detail.Edges.Store.Edges.Merchant.Edges.SupplierOrErr(); err == nil {

					return fmt.Sprintf(
						"%s %s\n%s\n%s, %s.\n%s", s.OtherName, s.LastName,
						detail.Edges.Store.Address.StreetName,
						detail.Edges.Store.Address.City, detail.Edges.Store.Address.Region, s.Phone,
					)
				}
				if r, err := detail.Edges.Store.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
					return fmt.Sprintf(
						"%s %s\n%s\n%s, %s.\n%s", r.OtherName, r.LastName,
						detail.Edges.Store.Address.StreetName,
						detail.Edges.Store.Address.City, detail.Edges.Store.Address.Region, r.Phone,
					)
				}
				return ""
			}()

			break
		}

	}
	return response
}

func (t *tookan) getFareEstimate(delivery, pickup *services.Coordinate) (*services.FareEstimateResponseData, error) {
	resData := services.FareEstimateRequest{
		TemplateName:      "order_delivery",
		PickupLongitude:   fmt.Sprintf("%f", pickup.Longitude),
		PickupLatitude:    fmt.Sprintf("%f", pickup.Latitude),
		APIKey:            t.APIKey,
		DeliveryLatitude:  fmt.Sprintf("%f", delivery.Latitude),
		DeliveryLongitude: fmt.Sprintf("%f", delivery.Longitude),
		FormulaType:       2,
		MapKeys: struct {
			MapPlanType  int    `json:"map_plan_type,omitempty"`
			GoogleAPIKey string `json:"google_api_key,omitempty"`
		}{
			MapPlanType:  1,
			GoogleAPIKey: t.GoogleAPIKey,
		},
	}

	payloadBytes, err := json.Marshal(resData)
	if err != nil {
		return nil, err
	}

	body := bytes.NewReader(payloadBytes)

	req, reqerr := http.NewRequest("POST", fmt.Sprintf("%s/get_fare_estimate", t.URL), body)

	if reqerr != nil {
		return nil, reqerr
	}
	req.Header.Set("Content-Type", "application/json")

	res, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		return nil, reserr
	}
	defer res.Body.Close()

	resp_body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resBody, errr := gabs.ParseJSON(resp_body)
	if errr != nil {
		return nil, err
	}

	formulas, err := resBody.Path("data.formula_fields").Children()
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

func (t *tookan) generatePickupStatus(data []*ent.OrderDetail) string {
	var status []string
	for _, d := range data {
		status = append(status, string(d.Status))
	}
	canceled := lo.CountBy[string](
		status, func(s string) bool {
			return s == "canceled"
		},
	)
	if canceled == len(status) {
		return "Declined"
	}
	return "Accepted"
}

func (t *tookan) formatWebhookPayload(request any) (*services.TookanWebhookResponse, error) {
	var response *services.TookanWebhookResponse
	resBody, err := gabs.ParseJSON(request.([]byte))
	if err != nil {
		return nil, err
	}
	var status int
	val, ok := resBody.Path("job_status").Data().(float64)
	if ok {
		status, _ = strconv.Atoi(fmt.Sprintf("%g", val))
	} else {
		val, ok := resBody.Path("job_status").Data().(string)
		if ok {
			status, _ = strconv.Atoi(val)
		} else {
			return nil, fmt.Errorf("could not cast job status to the appropriate type")
		}
	}
	response = &services.TookanWebhookResponse{
		JobStatus: status,
		JobState:  resBody.Path("job_state").Data().(string),
		JobToken:  resBody.Path("job_token").Data().(string),
	}
	return response, nil
}
