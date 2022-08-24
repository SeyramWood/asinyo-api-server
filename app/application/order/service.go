package order

import (
	"strings"
	"sync"

	"github.com/Jeffail/gabs"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type service struct {
	repo gateways.OrderRepo
}

func NewOrderService(repo gateways.OrderRepo) gateways.OrderService {
	return &service{repo: repo}
}

func (s service) Create(order *models.OrderResponse) (*ent.Order, error) {
	return s.repo.Insert(order)
}

func (s service) FetchAllByUser(userType string, id int) ([]*ent.Order, error) {
	return s.repo.ReadAllByUser(userType, id)
}
func (s service) FetchAllByStore(merchantId int) ([]*ent.Order, error) {
	return s.repo.ReadAllByStore(merchantId)
}
func (s service) FetchByStore(id, merchantId int) (*ent.Order, error) {
	return s.repo.ReadByStore(id, merchantId)
}

func (s service) FetchByUser(userType string, id int) (*ent.Order, error) {
	return s.repo.ReadByUser(userType, id)
}
func (s service) FetchAll() ([]*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s service) Fetch(id int) (*ent.Order, error) {
	return s.repo.Read(id)
}

func (s service) Update(order *models.OrderResponse) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (s service) Remove(id string) error {
	// TODO implement me
	panic("implement me")
}

func (s service) FormatOrderRequest(request []byte) (*models.OrderResponse, error) {
	var response *models.OrderResponse
	resBody, err := gabs.ParseJSON(request)

	if err != nil {
		return nil, err
	}
	if strings.Compare(resBody.Path("data.metadata.paymentMethod").Data().(string), "POD") == 0 {
		response = &models.OrderResponse{
			Event:     resBody.Path("event").Data().(string),
			Amount:    resBody.Path("data.amount").Data().(float64),
			Currency:  resBody.Path("data.currency").Data().(string),
			Channel:   resBody.Path("data.channel").Data().(string),
			Reference: resBody.Path("data.reference").Data().(string),
			PaidAt:    "",
			MetaData: &models.OrderResponseMetadata{
				User:           resBody.Path("data.metadata.user").Data().(string),
				UserType:       resBody.Path("data.metadata.userType").Data().(string),
				OrderNumber:    resBody.Path("data.metadata.orderNumber").Data().(string),
				Address:        resBody.Path("data.metadata.address").Data().(string),
				DeliveryMethod: resBody.Path("data.metadata.deliveryMethod").Data().(string),
				PaymentMethod:  resBody.Path("data.metadata.paymentMethod").Data().(string),
				DeliveryFee:    resBody.Path("data.metadata.deliveryFee").Data().(string),
				Pickup:         resBody.Path("data.metadata.pickup").Data().(string),
				Products: func() []*models.ProductDetailsResponse {
					var products []*models.ProductDetailsResponse
					children, _ := resBody.Path("data.metadata.products").Children()
					wg := sync.WaitGroup{}
					for _, child := range children {
						wg.Add(1)
						go func(child *gabs.Container) {
							defer wg.Done()
							pro := child.Data().(map[string]interface{})
							products = append(
								products, &models.ProductDetailsResponse{
									ID:         pro["id"].(string),
									Store:      pro["store"].(string),
									Quantity:   pro["quantity"].(string),
									Price:      pro["price"].(string),
									PromoPrice: pro["promoPrice"].(string),
								},
							)
						}(child)
					}
					wg.Wait()
					return products
				}(),
			},
		}
	} else {
		response = &models.OrderResponse{
			Event:     resBody.Path("event").Data().(string),
			Amount:    resBody.Path("data.amount").Data().(float64),
			Currency:  resBody.Path("data.currency").Data().(string),
			Channel:   resBody.Path("data.channel").Data().(string),
			Reference: resBody.Path("data.reference").Data().(string),
			PaidAt:    resBody.Path("data.paid_at").Data().(string),
			MetaData: &models.OrderResponseMetadata{
				User:           resBody.Path("data.metadata.user").Data().(string),
				UserType:       resBody.Path("data.metadata.userType").Data().(string),
				OrderNumber:    resBody.Path("data.metadata.orderNumber").Data().(string),
				Address:        resBody.Path("data.metadata.address").Data().(string),
				DeliveryMethod: resBody.Path("data.metadata.deliveryMethod").Data().(string),
				PaymentMethod:  resBody.Path("data.metadata.paymentMethod").Data().(string),
				DeliveryFee:    resBody.Path("data.metadata.deliveryFee").Data().(string),
				Pickup:         resBody.Path("data.metadata.pickup").Data().(string),
				Products: func() []*models.ProductDetailsResponse {
					var products []*models.ProductDetailsResponse
					children, _ := resBody.Path("data.metadata.products").Children()
					wg := sync.WaitGroup{}
					for _, child := range children {
						wg.Add(1)
						go func(child *gabs.Container) {
							defer wg.Done()
							pro := child.Data().(map[string]interface{})
							products = append(
								products, &models.ProductDetailsResponse{
									ID:         pro["id"].(string),
									Store:      pro["store"].(string),
									Quantity:   pro["quantity"].(string),
									Price:      pro["price"].(string),
									PromoPrice: pro["promoPrice"].(string),
								},
							)
						}(child)
					}
					wg.Wait()
					return products
				}(),
			},
		}
	}
	return response, nil
}

func (s service) UpdateOrderDetailStatus(request []byte) (*ent.Order, error) {
	resBody, err := gabs.ParseJSON(request)
	if err != nil {
		return nil, err
	}
	statuses, errr := resBody.ChildrenMap()

	if errr != nil {
		return nil, errr
	}

	return s.repo.UpdateOrderDetailStatus(statuses)
}
