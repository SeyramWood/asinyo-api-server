package presenters

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"

	"github.com/SeyramWood/ent"
)

type (
	OrderProductStore struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	OrderAddress struct {
		ID        int    `json:"id"`
		LastName  string `json:"lastName"`
		OtherName string `json:"otherName"`
		Address   string `json:"address"`
		City      string `json:"city"`
		Region    string `json:"region"`
	}
	OrderPickup struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		City    string `json:"city"`
		Region  string `json:"region"`
	}
	OrderProductDetail struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Unit  string `json:"unit"`
		Image string `json:"image"`
	}
	OrderProducts struct {
		ID         int                 `json:"id"`
		Price      float64             `json:"price"`
		PromoPrice float64             `json:"promoPrice"`
		Amount     float64             `json:"amount"`
		Quantity   int                 `json:"quantity"`
		Status     string              `json:"status"`
		CreatedAt  time.Time           `json:"createdAt"`
		UpdatedAt  time.Time           `json:"updatedAt"`
		Product    *OrderProductDetail `json:"product"`
		Store      *OrderProductStore  `json:"store"`
	}
	DetailOrder struct {
		ID             int              `json:"id"`
		OrderNumber    string           `json:"orderNumber"`
		Amount         float64          `json:"amount"`
		Currency       string           `json:"currency"`
		DeliveryFee    float64          `json:"deliveryFee"`
		DeliveryMethod string           `json:"deliveryMethod"`
		PaymentMethod  string           `json:"paymentMethod"`
		Reference      *string          `json:"reference"`
		Channel        *string          `json:"channel"`
		PaidAt         *string          `json:"paidAt"`
		Status         string           `json:"status"`
		DeliveredAt    *time.Time       `json:"deliveredAt"`
		CreatedAt      time.Time        `json:"createdAt"`
		UpdatedAt      time.Time        `json:"updatedAt"`
		Products       []*OrderProducts `json:"products"`
		Address        *OrderAddress    `json:"address"`
		Pickup         *OrderPickup     `json:"pickup"`
	}
	Order struct {
		ID            int       `json:"id"`
		OrderNumber   string    `json:"orderNumber"`
		Amount        float64   `json:"amount"`
		Currency      string    `json:"currency"`
		Channel       *string   `json:"channel"`
		PaymentMethod string    `json:"paymentMethod"`
		PaidAt        *string   `json:"paidAt"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	}
	StoreOrder struct {
		ID            int       `json:"id"`
		OrderNumber   string    `json:"orderNumber"`
		Amount        float64   `json:"amount"`
		Currency      string    `json:"currency"`
		Channel       *string   `json:"channel"`
		PaymentMethod string    `json:"paymentMethod"`
		PaidAt        *string   `json:"paidAt"`
		Status        string    `json:"status"`
		CreatedAt     time.Time `json:"createdAt"`
		UpdatedAt     time.Time `json:"updatedAt"`
	}
)

func OrderSuccessResponse(data *ent.Order) *fiber.Map {
	return successResponse(
		DetailOrder{
			ID:             data.ID,
			OrderNumber:    data.OrderNumber,
			Amount:         data.Amount,
			Currency:       data.Currency,
			Channel:        data.Channel,
			PaidAt:         data.PaidAt,
			Status:         string(data.Status),
			DeliveryFee:    data.DeliveryFee,
			DeliveryMethod: string(data.DeliveryMethod),
			PaymentMethod:  string(data.PaymentMethod),
			Reference:      data.Reference,
			DeliveredAt:    data.DeliveredAt,
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
			Products:       formatOrderDetails(data.Edges.Details),
			Address: func(edges *ent.Order) *OrderAddress {
				if add, err := edges.Edges.AddressOrErr(); err == nil {
					return &OrderAddress{
						ID:        add.ID,
						LastName:  add.LastName,
						OtherName: add.OtherName,
						Address:   add.Address,
						City:      add.City,
						Region:    add.Region,
					}
				}
				return nil
			}(data),
			Pickup: func(edges *ent.Order) *OrderPickup {
				if pick, err := edges.Edges.PickupOrErr(); err == nil {
					return &OrderPickup{
						ID:      pick.ID,
						Name:    pick.Name,
						Address: pick.Address,
						City:    pick.City,
						Region:  pick.Region,
					}
				}
				return nil
			}(data),
		},
	)
}
func OrdersSuccessResponse(data []*ent.Order) *fiber.Map {
	var response []Order
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Order) {
			defer wg.Done()
			response = append(
				response, Order{
					ID:            v.ID,
					OrderNumber:   v.OrderNumber,
					Amount:        v.Amount,
					Currency:      v.Currency,
					Channel:       v.Channel,
					PaymentMethod: string(v.PaymentMethod),
					PaidAt:        v.PaidAt,
					Status:        string(v.Status),
					CreatedAt:     v.CreatedAt,
					UpdatedAt:     v.UpdatedAt,
				},
			)
		}(v)
	}
	wg.Wait()
	return successResponse(response)
}
func StoreOrderSuccessResponse(data *ent.Order) *fiber.Map {
	detail := calculateStoreAmountOrderDetails(data.Edges.Details)
	return successResponse(
		DetailOrder{
			ID:             data.ID,
			OrderNumber:    data.OrderNumber,
			Amount:         detail["subtotal"].(float64),
			Currency:       data.Currency,
			Channel:        data.Channel,
			PaidAt:         data.PaidAt,
			Status:         detail["status"].(string),
			DeliveryFee:    data.DeliveryFee,
			DeliveryMethod: string(data.DeliveryMethod),
			PaymentMethod:  string(data.PaymentMethod),
			Reference:      data.Reference,
			DeliveredAt:    data.DeliveredAt,
			CreatedAt:      data.CreatedAt,
			UpdatedAt:      data.UpdatedAt,
			Products:       formatOrderDetails(data.Edges.Details),
			Address: func(edges *ent.Order) *OrderAddress {
				if add, err := edges.Edges.AddressOrErr(); err == nil {
					return &OrderAddress{
						ID:        add.ID,
						LastName:  add.LastName,
						OtherName: add.OtherName,
						Address:   add.Address,
						City:      add.City,
						Region:    add.Region,
					}
				}
				return nil
			}(data),
			Pickup: func(edges *ent.Order) *OrderPickup {
				if pick, err := edges.Edges.PickupOrErr(); err == nil {
					return &OrderPickup{
						ID:      pick.ID,
						Name:    pick.Name,
						Address: pick.Address,
						City:    pick.City,
						Region:  pick.Region,
					}
				}
				return nil
			}(data),
		},
	)
}
func StoreOrdersSuccessResponse(data []*ent.Order) *fiber.Map {
	var response []Order
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Order) {
			defer wg.Done()
			detail := calculateStoreAmountOrderDetails(v.Edges.Details)
			response = append(
				response, Order{
					ID:            v.ID,
					OrderNumber:   v.OrderNumber,
					Amount:        detail["subtotal"].(float64),
					Currency:      v.Currency,
					Channel:       v.Channel,
					PaymentMethod: string(v.PaymentMethod),
					PaidAt:        v.PaidAt,
					Status:        detail["status"].(string),
					CreatedAt:     v.CreatedAt,
					UpdatedAt:     v.UpdatedAt,
				},
			)
		}(v)
	}
	wg.Wait()
	return successResponse(response)
}

func OrderErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}

func formatOrderDetails(data []*ent.OrderDetail) []*OrderProducts {
	var response []*OrderProducts
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.OrderDetail) {
			defer wg.Done()
			response = append(
				response, &OrderProducts{
					ID:         v.ID,
					Amount:     v.Amount,
					Quantity:   v.Quantity,
					Price:      v.Price,
					PromoPrice: v.PromoPrice,
					Status:     string(v.Status),
					CreatedAt:  v.CreatedAt,
					UpdatedAt:  v.UpdatedAt,
					Product: &OrderProductDetail{
						ID:    v.Edges.Product.ID,
						Name:  v.Edges.Product.Name,
						Unit:  v.Edges.Product.Unit,
						Image: v.Edges.Product.Image,
					},
					Store: &OrderProductStore{
						ID:   v.Edges.Store.ID,
						Name: v.Edges.Store.Name,
					},
				},
			)
		}(v)
	}

	wg.Wait()

	return response
}
func calculateAmountOrderDetails(data []*ent.OrderDetail) map[string]interface{} {
	response := map[string]interface{}{}

	delivered := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "delivered"
		},
	)
	pending := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "pending"
		},
	)
	subtotal := lo.Reduce[*ent.OrderDetail, float64](
		data, func(agg float64, item *ent.OrderDetail, _ int) float64 {
			return agg + item.Amount
		}, 0,
	)

	if delivered == len(data) {
		response["status"] = "fulfilled"
		response["subtotal"] = subtotal
	} else if pending == len(data) {
		response["status"] = "pending"
		response["subtotal"] = subtotal
	} else {
		response["status"] = "in_progress"
		response["subtotal"] = subtotal
	}

	return response
}
func calculateStoreAmountOrderDetails(data []*ent.OrderDetail) map[string]interface{} {
	response := map[string]interface{}{}

	delivered := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "delivered"
		},
	)
	pending := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "pending"
		},
	)
	canceled := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "canceled"
		},
	)

	dispatched := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "dispatched"
		},
	)
	processing := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "processing"
		},
	)

	subtotal := lo.Reduce[*ent.OrderDetail, float64](
		data, func(agg float64, item *ent.OrderDetail, _ int) float64 {
			return agg + item.Amount
		}, 0,
	)

	if delivered == len(data) {
		response["status"] = "fulfilled"
		response["subtotal"] = subtotal
	} else if pending == len(data) {
		response["status"] = "pending"
		response["subtotal"] = subtotal
	} else if canceled == len(data) {
		response["status"] = "canceled"
		response["subtotal"] = subtotal
	} else if dispatched == len(data) {
		response["status"] = "dispatched"
		response["subtotal"] = subtotal
	} else if processing == len(data) {
		response["status"] = "processing"
		response["subtotal"] = subtotal
	} else {
		response["status"] = "in_progress"
		response["subtotal"] = subtotal
	}

	return response
}
func formatOrderAddress(edges *ent.Order) *OrderAddress {
	if add, err := edges.Edges.AddressOrErr(); err == nil {
		return &OrderAddress{
			ID:        add.ID,
			LastName:  add.LastName,
			OtherName: add.OtherName,
			Address:   add.Address,
			City:      add.City,
			Region:    add.Region,
		}
	}
	return nil
}
