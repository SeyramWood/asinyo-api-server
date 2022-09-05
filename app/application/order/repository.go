package order

import (
	"context"
	"strconv"
	"sync"

	"github.com/Jeffail/gabs"
	"github.com/samber/lo"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/address"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
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
	store := r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().OnlyIDX(ctx)
	results, err := r.db.Order.Query().
		Where(order.HasStoresWith(merchantstore.ID(store))).
		Order(ent.Asc(order.FieldCreatedAt)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Where(orderdetail.HasStoreWith(merchantstore.ID(store)))
				odq.Select(orderdetail.FieldAmount, orderdetail.FieldStatus)
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
	store := r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().OnlyIDX(ctx)
	result, err := r.db.Order.Query().
		Where(order.ID(id)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Where(orderdetail.HasStoreWith(merchantstore.ID(store)))
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
					address.FieldCity, address.FieldRegion,
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
		Order(ent.Asc(order.FieldCreatedAt)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r repository) Insert(res *models.OrderResponse) (*ent.Order, error) {
	switch res.MetaData.UserType {
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
	results, err := r.db.Order.Query().
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
					address.FieldCity, address.FieldRegion,
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
		Order(ent.Asc(order.FieldCreatedAt)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil

}

func (r repository) ReadAll() ([]*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r repository) Update(order *models.OrderResponse) (*ent.Order, error) {
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
	return r.ReadByStore(oId, mId)
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
func (r repository) insertMerchantOrder(res *models.OrderResponse) (*ent.Order, error) {
	ctx := context.Background()
	deliveryFee, _ := strconv.ParseFloat(res.MetaData.DeliveryFee, 32)
	userId, _ := strconv.Atoi(res.MetaData.User)
	addressId, _ := strconv.Atoi(res.MetaData.Address)
	pickupId, _ := strconv.Atoi(res.MetaData.Pickup)
	c := r.db.Merchant.Query().Where(merchant.ID(userId)).OnlyX(ctx)

	stores := r.getNewStores(res.MetaData.Products)

	if addressId != 0 {
		addr := r.db.Address.Query().Where(address.ID(addressId)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetMerchant(c).
			SetAddress(addr).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(deliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.MetaData.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.MetaData.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.MetaData.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.MetaData, o); err != nil {
			return nil, err
		}
		return o, nil
	}
	if pickupId != 0 {
		psd := r.db.PickupStation.Query().Where(pickupstation.ID(pickupId)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetMerchant(c).
			SetPickup(psd).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(deliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.MetaData.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.MetaData.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.MetaData.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.MetaData, o); err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil

}
func (r repository) insertCustomerOrder(res *models.OrderResponse) (*ent.Order, error) {
	ctx := context.Background()
	deliveryFee, _ := strconv.ParseFloat(res.MetaData.DeliveryFee, 32)
	userId, _ := strconv.Atoi(res.MetaData.User)
	addressId, _ := strconv.Atoi(res.MetaData.Address)
	pickupId, _ := strconv.Atoi(res.MetaData.Pickup)
	c := r.db.Customer.Query().Where(customer.ID(userId)).OnlyX(ctx)
	stores := r.getNewStores(res.MetaData.Products)
	if addressId != 0 {
		addr := r.db.Address.Query().Where(address.ID(addressId)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetCustomer(c).
			SetAddress(addr).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(deliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.MetaData.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.MetaData.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.MetaData.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.MetaData, o); err != nil {
			return nil, err
		}
		return o, nil
	}
	if pickupId != 0 {
		psd := r.db.PickupStation.Query().Where(pickupstation.ID(pickupId)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetCustomer(c).
			SetPickup(psd).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(deliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.MetaData.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.MetaData.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.MetaData.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.MetaData, o); err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil
}
func (r repository) insertAgentOrder(res *models.OrderResponse) (*ent.Order, error) {
	ctx := context.Background()
	deliveryFee, _ := strconv.ParseFloat(res.MetaData.DeliveryFee, 32)
	userId, _ := strconv.Atoi(res.MetaData.User)
	addressId, _ := strconv.Atoi(res.MetaData.Address)
	pickupId, _ := strconv.Atoi(res.MetaData.Pickup)
	c := r.db.Agent.Query().Where(agent.ID(userId)).OnlyX(ctx)

	stores := r.getNewStores(res.MetaData.Products)

	if addressId != 0 {
		addr := r.db.Address.Query().Where(address.ID(addressId)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetAgent(c).
			SetAddress(addr).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(deliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.MetaData.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.MetaData.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.MetaData.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.MetaData, o); err != nil {
			return nil, err
		}
		return o, nil
	}
	if pickupId != 0 {
		psd := r.db.PickupStation.Query().Where(pickupstation.ID(pickupId)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetAgent(c).
			SetPickup(psd).
			AddStores(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(deliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.MetaData.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.MetaData.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.MetaData.PaymentMethod)).
			SetStatus("pending").
			Save(ctx)
		if oErr != nil {
			return nil, oErr
		}
		if err := r.insertOrderDetails(res.MetaData, o); err != nil {
			return nil, err
		}
		return o, nil
	}

	return nil, nil

}
func (r repository) calculateAmount(product *models.ProductDetailsResponse) float64 {
	var amount float64
	price, _ := strconv.ParseFloat(product.Price, 32)
	promoPrice, _ := strconv.ParseFloat(product.PromoPrice, 32)
	qty, _ := strconv.ParseFloat(product.Quantity, 32)
	if promoPrice > 0 {
		amount = promoPrice * qty
	} else {
		amount = price * qty
	}

	return amount
}

func (r repository) readCustomerOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasCustomerWith(customer.ID(id))).
		Order(ent.Asc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r repository) readAgentOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasAgentWith(agent.ID(id))).
		Order(ent.Asc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r repository) readMerchantOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasMerchantWith(merchant.ID(id))).
		Order(ent.Asc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r repository) insertOrderDetails(metadata *models.OrderResponseMetadata, o *ent.Order) error {
	ctx := context.Background()
	bulk := make([]*ent.OrderDetailCreate, len(metadata.Products))
	wg := sync.WaitGroup{}
	for i, item := range metadata.Products {
		wg.Add(1)
		go func(item *models.ProductDetailsResponse, i int) {
			defer wg.Done()
			price, _ := strconv.ParseFloat(item.Price, 32)
			promoPrice, _ := strconv.ParseFloat(item.PromoPrice, 32)
			qty, _ := strconv.Atoi(item.Quantity)
			amount := r.calculateAmount(item)
			pId, _ := strconv.Atoi(item.ID)
			sId, _ := strconv.Atoi(item.Store)
			prod := r.db.Product.Query().Where(product.ID(pId)).OnlyX(ctx)
			store := r.db.MerchantStore.Query().Where(merchantstore.ID(sId)).OnlyX(ctx)
			bulk[i] = r.db.OrderDetail.Create().
				SetOrder(o).
				SetProduct(prod).
				SetStore(store).
				SetPrice(price).
				SetPromoPrice(promoPrice).
				SetAmount(amount).
				SetQuantity(qty)
		}(item, i)
	}
	wg.Wait()
	_, err := r.db.OrderDetail.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) getNewStores(products []*models.ProductDetailsResponse) []*ent.MerchantStore {
	stores := map[string]interface{}{
		"ids":       []int{},
		"instances": []*ent.MerchantStore{},
	}
	for _, p := range products {
		id, _ := strconv.Atoi(p.Store)
		if !lo.Contains[int](stores["ids"].([]int), id) {
			stores["ids"] = append(stores["ids"].([]int), id)
			stores["instances"] = append(
				stores["instances"].([]*ent.MerchantStore),
				r.db.MerchantStore.Query().Where(merchantstore.ID(id)).OnlyX(context.Background()),
			)
		}
	}
	return stores["instances"].([]*ent.MerchantStore)
}
