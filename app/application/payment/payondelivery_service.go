package payment

import (
	"strconv"
	"sync"

	"github.com/Jeffail/gabs"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
)

type payOnDeliveryService struct {
	repo gateways.PaymentRepo
}

func newPayOnDeliveryService(repo gateways.PaymentRepo) gateways.PaymentService {
	return &payOnDeliveryService{
		repo: repo,
	}
}

func (p payOnDeliveryService) Pay(request any) (any, error) {
	// TODO implement me
	panic("implement me")
}

func (p payOnDeliveryService) Verify(reference string) (any, error) {
	// TODO implement me
	panic("implement me")
}

func (p payOnDeliveryService) FormatPayload(request any) (*models.OrderPayload, error) {
	var response *services.PaystackResponse
	resBody, err := gabs.ParseJSON(request.([]byte))
	if err != nil {
		return nil, err
	}
	response = &services.PaystackResponse{
		Event:     "",
		Amount:    resBody.Path("data.amount").Data().(float64),
		Currency:  resBody.Path("data.currency").Data().(string),
		Channel:   resBody.Path("data.channel").Data().(string),
		Reference: resBody.Path("data.reference").Data().(string),
		PaidAt:    "",
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
