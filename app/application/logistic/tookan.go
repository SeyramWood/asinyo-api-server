package logistic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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
		DoneChan:     conf.DoneChan,
		ErrorChan:    conf.ErrorChan,
	}
}

func (t *tookan) New(repo gateways.OrderRepo) gateways.LogisticService {
	t.orderRepo = repo
	return t
}

func (t *tookan) DoTask(order *ent.Order, deliveryType string) {
	t.WG.Add(1)
	t.DataChan <- order
	if deliveryType != "" {
		t.TaskType = deliveryType
	}
}

func (t *tookan) Listen() {
	for {
		select {
		case ord := <-t.DataChan:
			go t.createTask(ord, t.ErrorChan)
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
	close(t.ErrorChan)
	close(t.DoneChan)
}

func (t *tookan) ListenOnWebhook() {
	// TODO implement me
	panic("implement me")
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

func (t *tookan) createTask(order *ent.Order, errorChan chan error) {
	defer t.WG.Done()
	switch t.TaskType {
	case "delivery_task":
		if err := t.createDeliveryTask(order); err != nil {
			errorChan <- err
		}
	case "create_multiple_tasks":
		if err := t.createPickupDeliveryTasks(order); err != nil {
			errorChan <- err
		}
	case "edit_multiple_tasks":
		if err := t.updatePickupDeliveryTasks(order); err != nil {
			errorChan <- err
		}
	}
}

func (t *tookan) createDeliveryTask(order *ent.Order) error {
	currentTime := time.Date(2022, time.December, 25, 23, 0, 0, 0, time.UTC)
	task := &services.TookanDeliveryTask{
		APIKey:           t.APIKey,
		OrderID:          order.OrderNumber,
		JobDescription:   "Order Delivery",
		CustomerEmail:    "",
		CustomerUsername: fmt.Sprintf("%s %s", order.Edges.Address.OtherName, order.Edges.Address.LastName),
		CustomerPhone:    fmt.Sprintf("%s", order.Edges.Address.Phone),
		CustomerAddress: fmt.Sprintf(
			"%s %s\n%s\n%s, %s.\n%s", order.Edges.Address.OtherName, order.Edges.Address.LastName,
			order.Edges.Address.Address,
			order.Edges.Address.City, order.Edges.Address.Region, order.Edges.Address.Phone,
		),
		Latitude:            "",
		Longitude:           "",
		JobDeliveryDatetime: currentTime.Format("01-02-2006 15:04"),
		CustomFieldTemplate: "order_delivery",
		MetaData:            t.formatMetadata(order),
		TeamID:              "",
		AutoAssignment:      "1",
		HasPickup:           "0",
		HasDelivery:         "1",
		LayoutType:          "0",
		TrackingLink:        1,
		Timezone:            "-330",
		FleetID:             "",
		RefImages:           nil,
		Notify:              1,
		Tags:                "delivery, order",
		Geofence:            1,
	}

	// log.Fatalln(task)

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
	defer res.Body.Close()
	// TODO work with response
	var response services.TookanTaskResponse
	derr := json.NewDecoder(res.Body).Decode(&response)
	if derr != nil {
		return derr
	}
	fmt.Println(response)

	if response.Status == 200 {
		// TODO save data to db

	}

	return nil
}

func (t *tookan) createPickupDeliveryTasks(order *ent.Order) error {
	task := &services.TookanPickupDeliveryTask{
		APIKey:         t.APIKey,
		FleetID:        0,
		Timezone:       -330,
		HasPickup:      1,
		HasDelivery:    1,
		LayoutType:     0,
		Geofence:       0,
		TeamID:         "",
		AutoAssignment: 0,
		Tags:           "order, pickup, delivery",
		Pickups:        t.formatPickups(order),
		Deliveries:     t.formatDeliveries(order),
	}
	payloadBytes, err := json.Marshal(&task)

	if err != nil {
		return err
	}

	body := bytes.NewReader(payloadBytes)

	req, reqerr := http.NewRequest("POST", fmt.Sprintf("%s/create_multiple_tasks", t.URL), body)

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
	defer res.Body.Close()

	if response.Status == 200 {
		_ = t.repo.UpdateOrderDeliveryTask(response.Data.Deliveries[0].OrderID, t.StoreIds)
		resData := &models.TookanMultiTaskResponse{
			Pickups:    response.Data.Pickups,
			Deliveries: response.Data.Deliveries,
			Geofence:   response.Data.GeofenceDetails,
		}
		_, err := t.repo.InsertResponse(resData)
		if err != nil {
			t.ErrorChan <- err
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
		resData := &models.TookanMultiTaskResponse{
			Pickups:    response.Data.Pickups,
			Deliveries: response.Data.Deliveries,
			Geofence:   response.Data.GeofenceDetails,
		}
		_, err := t.repo.UpdateResponse(resData)
		if err != nil {
			t.ErrorChan <- err
		}
	}

	return nil
}

func (t *tookan) formatMetadata(data *ent.Order) []services.TookanMetadata {
	type statusCheck struct {
		details []*ent.OrderDetail
	}
	var orderStatus statusCheck

	for _, detail := range data.Edges.Details {
		if lo.Contains[int](data.StoreTasksCreated, detail.Edges.Store.ID) {
			continue
		}
		orderStatus = statusCheck{
			details: append(orderStatus.details, detail),
		}
	}

	return []services.TookanMetadata{
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
			Data:  data.Amount,
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
				"%s %s\n%s\n%s, %s\n%s", data.Edges.Address.OtherName, data.Edges.Address.LastName,
				data.Edges.Address.Address,
				data.Edges.Address.City, data.Edges.Address.Region, data.Edges.Address.Phone,
			),
		},
		{
			Label: "Products",
			Data:  t.formatOrderDetails(orderStatus.details, 0),
		},
	}
}

func (t *tookan) formatPickupMetadata(
	data []*ent.OrderDetail, orderNum string, storeId int, pickupStatus string,
) []services.TookanMetadata {

	return []services.TookanMetadata{
		{
			Label: "Reference",
			Data:  orderNum,
		},
		{
			Label: "Status",
			Data:  pickupStatus,
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

func (t *tookan) formatPickups(order *ent.Order) []services.TookanPickupDelivery {

	currentTime := time.Date(2022, time.December, 25, 23, 0, 0, 0, time.UTC)
	var response []services.TookanPickupDelivery
	type statusCheck struct {
		detail *ent.OrderDetail
		status []string
	}

	orderStatus := make(map[int]statusCheck)

	for _, detail := range order.Edges.Details {
		if lo.Contains[int](order.StoreTasksCreated, detail.Edges.Store.ID) {
			continue
		}
		if !lo.Contains[int](t.StoreIds, detail.Edges.Store.ID) {
			t.StoreIds = append(t.StoreIds, detail.Edges.Store.ID)
		}
		orderStatus[detail.Edges.Store.ID] = statusCheck{
			detail: detail,
			status: append(orderStatus[detail.Edges.Store.ID].status, string(detail.Status)),
		}
	}

	for _, o := range orderStatus {
		formattedPickup := func() *services.TookanPickupDelivery {
			pickupStatus := func() string {
				canceled := lo.CountBy[string](
					o.status, func(s string) bool {
						return s == "canceled"
					},
				)
				if canceled == len(o.status) {
					return "Declined"
				}
				return "Accepted"
			}()
			if s, err := o.detail.Edges.Store.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
				return &services.TookanPickupDelivery{
					Address: fmt.Sprintf(
						"%s %s\n%s, %s.", o.detail.Edges.Store.Address.StreetName,
						o.detail.Edges.Store.Address.StreetName,
						o.detail.Edges.Store.Address.City, o.detail.Edges.Store.Address.Region,
					),
					Latitude:       0,
					Longitude:      0,
					Time:           currentTime.Format("01-02-2006 15:04"),
					Phone:          s.Phone,
					JobDescription: "Order Pickup",
					TemplateName:   "order_pickup",
					TemplateData: t.formatPickupMetadata(
						order.Edges.Details, order.OrderNumber, o.detail.Edges.Store.ID, pickupStatus,
					),
					RefImages: nil,
					Name:      fmt.Sprintf("%s %s", s.OtherName, s.LastName),
					Email:     "",
					OrderID:   order.OrderNumber,
				}
			}
			if r, err := o.detail.Edges.Store.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
				return &services.TookanPickupDelivery{
					Address: fmt.Sprintf(
						"%s %s\n%s, %s.", o.detail.Edges.Store.Address.StreetName,
						o.detail.Edges.Store.Address.StreetName,
						o.detail.Edges.Store.Address.City, o.detail.Edges.Store.Address.Region,
					),
					Latitude:       0,
					Longitude:      0,
					Time:           currentTime.Format("01-02-2006 15:04"),
					Phone:          r.Phone,
					JobDescription: "Order Pickup",
					TemplateName:   "order_pickup",
					TemplateData: t.formatPickupMetadata(
						order.Edges.Details, order.OrderNumber, o.detail.Edges.Store.ID, pickupStatus,
					),
					RefImages: nil,
					Name:      fmt.Sprintf("%s %s", r.OtherName, r.LastName),
					Email:     "",
					OrderID:   order.OrderNumber,
				}
			}
			return nil
		}()
		response = append(response, *formattedPickup)
	}

	return response
}

func (t *tookan) formatDeliveries(order *ent.Order) []services.TookanPickupDelivery {
	currentTime := time.Date(2022, time.December, 25, 23, 0, 0, 0, time.UTC)

	return []services.TookanPickupDelivery{
		{
			Address: fmt.Sprintf(
				"%s %s\n%s\n%s, %s.\n%s", order.Edges.Address.OtherName, order.Edges.Address.LastName,
				order.Edges.Address.Address,
				order.Edges.Address.City, order.Edges.Address.Region, order.Edges.Address.Phone,
			),
			Latitude:       0,
			Longitude:      0,
			Time:           currentTime.Format("01-02-2006 15:04"),
			Phone:          fmt.Sprintf("%s", order.Edges.Address.Phone),
			JobDescription: "Order Delivery",
			TemplateName:   "order_delivery",
			TemplateData:   t.formatMetadata(order),
			RefImages:      nil,
			Name:           fmt.Sprintf("%s %s", order.Edges.Address.OtherName, order.Edges.Address.LastName),
			Email:          "",
			OrderID:        order.OrderNumber,
		},
	}
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
