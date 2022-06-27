package presenters

import (
	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
	"strings"
	"sync"
	"time"
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
		Reference      string           `json:"reference"`
		Channel        string           `json:"channel"`
		PaidAt         string           `json:"paidAt"`
		Status         string           `json:"status"`
		DeliveredAt    *time.Time       `json:"deliveredAt"`
		CreatedAt      time.Time        `json:"createdAt"`
		UpdatedAt      time.Time        `json:"updatedAt"`
		Products       []*OrderProducts `json:"products"`
		Address        *OrderAddress    `json:"address"`
		Pickup         *OrderPickup     `json:"pickup"`
	}
	Order struct {
		ID          int       `json:"id"`
		OrderNumber string    `json:"orderNumber"`
		Amount      float64   `json:"amount"`
		Currency    string    `json:"currency"`
		Channel     string    `json:"channel"`
		PaidAt      string    `json:"paidAt"`
		Status      string    `json:"status"`
		CreatedAt   time.Time `json:"createdAt"`
		UpdatedAt   time.Time `json:"updatedAt"`
	}
)

func OrderSuccessResponse(data *ent.Order) *fiber.Map {
	return successResponse(DetailOrder{
		ID:             data.ID,
		OrderNumber:    data.OrderNumber,
		Amount:         data.Amount,
		Currency:       data.Currency,
		Channel:        strings.Replace(data.Channel, "_", " ", 1),
		PaidAt:         data.PaidAt,
		Status:         string(data.Status),
		DeliveryFee:    data.DeliveryFee,
		DeliveryMethod: string(data.DeliveryMethod),
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
	})
}
func OrdersSuccessResponse(data []*ent.Order) *fiber.Map {
	var response []Order
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Order) {
			defer wg.Done()
			response = append(response, Order{
				ID:          v.ID,
				OrderNumber: v.OrderNumber,
				Amount:      v.Amount,
				Currency:    v.Currency,
				Channel:     strings.Replace(v.Channel, "_", " ", 1),
				PaidAt:      v.PaidAt,
				Status:      string(v.Status),
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
			})
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
			response = append(response, &OrderProducts{
				ID:         v.ID,
				Amount:     v.Amount,
				Quantity:   v.Quantity,
				Price:      v.Price,
				PromoPrice: v.PromoPrice,
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
			})
		}(v)
	}

	wg.Wait()

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
