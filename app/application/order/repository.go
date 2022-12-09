package order

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/Jeffail/gabs"
	"github.com/samber/lo"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/logistic"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/pickupstation"
	"github.com/SeyramWood/ent/product"
)

type repository struct {
	db *ent.Client
}

func NewOrderRepo(db *database.Adapter) gateways.OrderRepo {
	return &repository{db.DB}
}

func (r repository) ReadByUser(userType string, id int) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) ReadAllByUser(userType string, id int) ([]*ent.Order, error) {
	switch userType {
	case "retailer", "supplier":
		return r.readMerchantOrders(id)
	case "customer":
		return r.readCustomerOrders(id)
	case "agent":
		return r.readAgentOrders(id)
	default:
		return nil, nil
	}
}
func (r repository) ReadAllByStore(merchantId int) ([]*ent.Order, error) {
	ctx := context.Background()
	results, err := r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().
		QueryOrders().
		Order(ent.Desc(order.FieldCreatedAt)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Where(
					orderdetail.HasStoreWith(
						merchantstore.ID(r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().OnlyIDX(ctx)),
					),
				)
				odq.Select(orderdetail.FieldAmount, orderdetail.FieldStatus)
			},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r repository) ReadAllByAgentStore(agentId int) ([]*ent.Order, error) {
	ctx := context.Background()
	results, err := r.db.Agent.Query().Where(agent.ID(agentId)).
		QueryStore().
		QueryOrders().
		Order(ent.Desc(order.FieldCreatedAt)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Where(
					orderdetail.HasStore(),
				)
				odq.Select(orderdetail.FieldAmount, orderdetail.FieldStatus)
			},
		).
		WithStores(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(merchantstore.FieldID, merchantstore.FieldName)
			},
		).
		All(ctx)

	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r repository) ReadByStore(id, merchantId int) (*ent.Order, error) {
	ctx := context.Background()
	result, err := r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().
		QueryOrders().
		Where(order.ID(id)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Where(
					orderdetail.HasStoreWith(
						merchantstore.ID(r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().OnlyIDX(ctx)),
					),
				)
				odq.WithStore(
					func(msq *ent.MerchantStoreQuery) {
						msq.Select(
							merchantstore.FieldID, merchantstore.FieldName, merchantstore.FieldAddress,
						)
						msq.WithMerchant(
							func(mq *ent.MerchantQuery) {
								mq.WithRetailer()
								mq.WithSupplier()
							},
						)
					},
				)
				odq.WithProduct(
					func(pq *ent.ProductQuery) {
						pq.Select(
							product.FieldID, product.FieldName, product.FieldUnit, product.FieldImage,
							product.FieldWeight,
						)
					},
				)
			},
		).
		WithAddress(
			func(aq *ent.AddressQuery) {
				aq.Select(
					address.FieldID, address.FieldLastName, address.FieldOtherName, address.FieldAddress,
					address.FieldCity, address.FieldStreetName, address.FieldStreetNumber, address.FieldDistrict,
					address.FieldRegion, address.FieldCountry,
					address.FieldPhone, address.FieldOtherPhone,
				)
			},
		).
		WithPickup(
			func(pq *ent.PickupStationQuery) {
				pq.Select(
					pickupstation.FieldID, pickupstation.FieldName, pickupstation.FieldAddress,
					pickupstation.FieldCity,
					pickupstation.FieldRegion,
				)
			},
		).
		WithStores(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(merchantstore.FieldID, merchantstore.FieldName)
			},
		).
		WithLogistic(
			func(lg *ent.LogisticQuery) {
				lg.Select(
					logistic.FieldID, logistic.FieldTask,
				).WithStore(
					func(msq *ent.MerchantStoreQuery) {
						msq.Select(merchantstore.FieldID, merchantstore.FieldName)
					},
				)
			},
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) ReadByAgentStore(id, agentId int) (*ent.Order, error) {
	ctx := context.Background()
	result, err := r.db.Agent.Query().Where(agent.ID(agentId)).QueryStore().
		Select(merchantstore.FieldID, merchantstore.FieldName).
		QueryOrders().
		Where(order.ID(id)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Where(
					orderdetail.HasStore(),
				)
				odq.WithStore(
					func(msq *ent.MerchantStoreQuery) {
						msq.Select(
							merchantstore.FieldID, merchantstore.FieldName, merchantstore.FieldAddress,
						)
						msq.WithMerchant(
							func(mq *ent.MerchantQuery) {
								mq.WithRetailer()
								mq.WithSupplier()
							},
						)
					},
				)
				odq.WithProduct(
					func(pq *ent.ProductQuery) {
						pq.Select(
							product.FieldID, product.FieldName, product.FieldUnit, product.FieldImage,
							product.FieldWeight,
						)
					},
				)

			},
		).
		WithAddress(
			func(aq *ent.AddressQuery) {
				aq.Select(
					address.FieldID, address.FieldLastName, address.FieldOtherName, address.FieldAddress,
					address.FieldCity, address.FieldStreetName, address.FieldStreetNumber, address.FieldDistrict,
					address.FieldRegion, address.FieldCountry,
					address.FieldPhone, address.FieldOtherPhone,
				)
			},
		).
		WithPickup(
			func(pq *ent.PickupStationQuery) {
				pq.Select(
					pickupstation.FieldID, pickupstation.FieldName, pickupstation.FieldAddress,
					pickupstation.FieldCity,
					pickupstation.FieldRegion,
				)
			},
		).
		WithStores(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(merchantstore.FieldID, merchantstore.FieldName)
			},
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) ReadByStoreOrderDetail(orderId int) ([]*ent.OrderDetail, error) {
	result, err := r.db.Order.Query().
		Where(order.ID(orderId)).QueryDetails().Where(
		orderdetail.HasStore(),
	).
		WithStore(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(
					merchantstore.FieldID, merchantstore.FieldName, merchantstore.FieldAddress,
				)
				msq.WithMerchant(
					func(mq *ent.MerchantQuery) {
						mq.WithRetailer()
						mq.WithSupplier()
					},
				)
			},
		).
		WithProduct(
			func(pq *ent.ProductQuery) {
				pq.Select(
					product.FieldID, product.FieldName, product.FieldUnit, product.FieldImage,
					product.FieldWeight,
				)
			},
		).
		All(context.Background())

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r repository) Insert(res *models.OrderPayload) (*ent.Order, error) {
	switch res.Metadata.UserType {
	case "retailer", "supplier":
		return r.insertMerchantOrder(res)
	case "customer":
		return r.insertCustomerOrder(res)
	case "agent":
		return r.insertAgentOrder(res)
	default:
		return nil, nil
	}
}

func (r repository) Read(id int) (*ent.Order, error) {

	result, err := r.db.Order.Query().
		Where(order.ID(id)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.WithStore(
					func(msq *ent.MerchantStoreQuery) {
						msq.Select(merchantstore.FieldID, merchantstore.FieldName)
					},
				)
				odq.WithProduct(
					func(pq *ent.ProductQuery) {
						pq.Select(
							product.FieldID, product.FieldName, product.FieldUnit, product.FieldImage,
							product.FieldWeight,
						)
					},
				)
			},
		).
		WithAddress(
			func(aq *ent.AddressQuery) {
				aq.Select(
					address.FieldID, address.FieldLastName, address.FieldOtherName, address.FieldAddress,
					address.FieldCity, address.FieldStreetName, address.FieldStreetNumber, address.FieldDistrict,
					address.FieldRegion, address.FieldCountry,
					address.FieldPhone, address.FieldOtherPhone,
				)
			},
		).
		WithPickup(
			func(pq *ent.PickupStationQuery) {
				pq.Select(
					pickupstation.FieldID, pickupstation.FieldName, pickupstation.FieldAddress,
					pickupstation.FieldCity,
					pickupstation.FieldRegion,
				)
			},
		).
		WithLogistic(
			func(lg *ent.LogisticQuery) {
				lg.Select(
					logistic.FieldID, logistic.FieldTask,
				).WithStore(
					func(msq *ent.MerchantStoreQuery) {
						msq.Select(merchantstore.FieldID, merchantstore.FieldName)
					},
				)
			},
		).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (r repository) ReadAll() ([]*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) Update(order *services.PaystackResponse) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}

func (r repository) UpdateOrderDetailStatus(requests map[string]*gabs.Container) (*ent.Order, error) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	statuses, _ := requests["status"].ChildrenMap()

	for key, val := range statuses {
		wg.Add(1)
		id, _ := strconv.Atoi(key)
		status := val.Data().(string)
		go func(id int, status string) {
			defer wg.Done()
			switch status {
			case "delivered":
				_ = r.db.OrderDetail.UpdateOneID(id).SetStatus(orderdetail.StatusDelivered).SaveX(ctx)
			case "processing":
				_ = r.db.OrderDetail.UpdateOneID(id).SetStatus(orderdetail.StatusProcessing).SaveX(ctx)
			case "dispatched":
				_ = r.db.OrderDetail.UpdateOneID(id).SetStatus(orderdetail.StatusDispatched).SaveX(ctx)
			case "canceled":
				_ = r.db.OrderDetail.UpdateOneID(id).SetStatus(orderdetail.StatusCanceled).SaveX(ctx)
			default:
				_ = r.db.OrderDetail.UpdateOneID(id).SetStatus(orderdetail.StatusPending).SaveX(ctx)
			}
		}(id, status)
	}
	wg.Wait()

	oId, _ := strconv.Atoi(requests["order"].String())
	mId, _ := strconv.Atoi(requests["merchant"].String())

	results, err := r.db.Order.Query().
		Where(order.ID(oId)).QueryDetails().
		Select(orderdetail.FieldStatus).
		All(ctx)

	if err != nil {
		return nil, err
	}
	_ = r.db.Order.UpdateOneID(oId).SetStatus(r.checkOrderStatus(results)).SaveX(ctx)
	if strings.Compare(requests["userType"].Data().(string), "agent") == 0 {
		return r.ReadByAgentStore(oId, mId)
	}
	return r.ReadByStore(oId, mId)
}

func (r repository) ReadOrderStoreMerchants(orderId int) (*ent.Order, error) {
	result, err := r.db.Order.Query().Where(order.ID(orderId)).WithStores(
		func(msq *ent.MerchantStoreQuery) {
			msq.WithMerchant()
		},
	).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) checkOrderStatus(data []*ent.OrderDetail) order.Status {
	var status order.Status

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
	if delivered == len(data) {
		status = order.StatusFulfilled
	} else if pending == len(data) {
		status = order.StatusPending
	} else if canceled == len(data) {
		status = order.StatusCanceled
	} else {
		status = order.StatusInProgress
	}

	return status
}

func (r repository) insertMerchantOrder(res *models.OrderPayload) (*ent.Order, error) {
	ctx := context.Background()
	c := r.db.Merchant.Query().Where(merchant.ID(res.Metadata.User)).OnlyX(ctx)
	stores := r.getNewStores(res.Metadata.Products)
	if res.Metadata.Address != 0 {
		addr := r.db.Address.Query().Where(address.ID(res.Metadata.Address)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetMerchant(c).
			SetAddress(addr).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.Metadata, o); err != nil {
			return nil, err
		}
		return o, nil
	}
	if res.Metadata.Pickup != 0 {
		psd := r.db.PickupStation.Query().Where(pickupstation.ID(res.Metadata.Pickup)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetMerchant(c).
			SetPickup(psd).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.Metadata, o); err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil

}
func (r repository) insertCustomerOrder(res *models.OrderPayload) (*ent.Order, error) {
	ctx := context.Background()
	c := r.db.Customer.Query().Where(customer.ID(res.Metadata.User)).OnlyX(ctx)
	stores := r.getNewStores(res.Metadata.Products)
	if res.Metadata.Address != 0 {
		addr := r.db.Address.Query().Where(address.ID(res.Metadata.Address)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetCustomer(c).
			SetAddress(addr).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.Metadata, o); err != nil {
			return nil, err
		}
		return o, nil
	}
	if res.Metadata.Pickup != 0 {
		psd := r.db.PickupStation.Query().Where(pickupstation.ID(res.Metadata.Pickup)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetCustomer(c).
			SetPickup(psd).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.Metadata, o); err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil
}
func (r repository) insertAgentOrder(res *models.OrderPayload) (*ent.Order, error) {
	ctx := context.Background()
	c := r.db.Agent.Query().Where(agent.ID(res.Metadata.User)).OnlyX(ctx)
	stores := r.getNewStores(res.Metadata.Products)
	if res.Metadata.Address != 0 {
		addr := r.db.Address.Query().Where(address.ID(res.Metadata.Address)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetAgent(c).
			SetAddress(addr).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.Metadata, o); err != nil {
			return nil, err
		}
		return o, nil
	}
	if res.Metadata.Pickup != 0 {
		psd := r.db.PickupStation.Query().Where(pickupstation.ID(res.Metadata.Pickup)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetAgent(c).
			SetPickup(psd).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.Metadata, o); err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil

}

func (r repository) calculateAmount(product *services.ProductDetails) float64 {
	var amount float64
	if product.PromoPrice > 0 {
		amount = product.PromoPrice * float64(product.Quantity)
	} else {
		amount = product.Price * float64(product.Quantity)
	}
	return amount
}

func (r repository) readCustomerOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasCustomerWith(customer.ID(id))).
		Order(ent.Desc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r repository) readAgentOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasAgentWith(agent.ID(id))).
		Order(ent.Desc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r repository) readMerchantOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasMerchantWith(merchant.ID(id))).
		Order(ent.Desc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r repository) insertOrderDetails(metadata *models.OrderPayloadMetadata, o *ent.Order) error {
	ctx := context.Background()
	bulk := make([]*ent.OrderDetailCreate, len(metadata.Products))
	wg := sync.WaitGroup{}
	for i, item := range metadata.Products {
		wg.Add(1)
		go func(item *services.ProductDetails, i int) {
			defer wg.Done()

			amount := r.calculateAmount(item)
			// prod := r.db.Product.Query().Where(product.ID(item.ID)).OnlyX(ctx)
			store := r.db.MerchantStore.Query().Where(merchantstore.ID(item.Store)).OnlyX(ctx)
			prod, _ := r.db.Product.UpdateOneID(item.ID).AddBestDeal(1).Save(ctx)

			bulk[i] = r.db.OrderDetail.Create().
				SetOrder(o).
				SetProduct(prod).
				SetStore(store).
				SetPrice(item.Price).
				SetPromoPrice(item.PromoPrice).
				SetAmount(amount).
				SetQuantity(item.Quantity)
		}(item, i)
	}
	wg.Wait()
	_, err := r.db.OrderDetail.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) getNewStores(products []*services.ProductDetails) []*ent.MerchantStore {
	stores := map[string]interface{}{
		"ids":       []int{},
		"instances": []*ent.MerchantStore{},
	}
	for _, p := range products {
		if !lo.Contains[int](stores["ids"].([]int), p.Store) {
			stores["ids"] = append(stores["ids"].([]int), p.Store)
			stores["instances"] = append(
				stores["instances"].([]*ent.MerchantStore),
				r.db.MerchantStore.Query().Where(merchantstore.ID(p.Store)).OnlyX(context.Background()),
			)
		}
	}
	return stores["instances"].([]*ent.MerchantStore)
}
