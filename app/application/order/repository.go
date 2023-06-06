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
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/SeyramWood/ent/pickupstation"
	"github.com/SeyramWood/ent/product"
	"github.com/SeyramWood/ent/productcategoryminor"
)

type repository struct {
	db *ent.Client
}

func NewOrderRepo(db *database.Adapter) gateways.OrderRepo {
	return &repository{db.DB}
}

func (r *repository) ReadByUser(userType string, id int) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r *repository) ReadAllByUser(userType string, id int) ([]*ent.Order, error) {
	switch userType {
	case "retailer", "supplier":
		return r.readMerchantOrders(id)
	case "business", "individual":
		return r.readCustomerOrders(id)
	case "agent":
		return r.readAgentOrders(id)
	default:
		return nil, nil
	}
}

func (r *repository) ReadAllByStore(merchantId int) ([]*ent.Order, error) {
	ctx := context.Background()
	results, err := r.db.Merchant.Query().Where(merchant.ID(merchantId)).QueryStore().
		QueryOrders().
		Where(
			order.Or(
				order.CustomerApprovalIsNil(),
				order.CustomerApprovalEQ(order.CustomerApprovalApproved),
			),
		).
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

func (r *repository) ReadAllByAgentStore(agentId int) ([]*ent.Order, error) {
	ctx := context.Background()
	results, err := r.db.Agent.Query().Where(agent.ID(agentId)).
		QueryStore().
		QueryOrders().
		Where(
			order.Or(
				order.CustomerApprovalIsNil(),
				order.CustomerApprovalEQ(order.CustomerApprovalApproved),
			),
		).
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

func (r *repository) ReadByStore(id, merchantId int) (*ent.Order, error) {
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
							product.FieldWeight, product.FieldQuantity,
						)
						pq.WithMinor(
							func(mq *ent.ProductCategoryMinorQuery) {
								mq.Select(productcategoryminor.FieldCategory)
							},
						)
					},
				)
			},
		).
		WithAddress().
		WithPickup().
		WithStores(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(merchantstore.FieldID, merchantstore.FieldName)
			},
		).
		WithLogistic().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) ReadByAgentStore(id, agentId int) (*ent.Order, error) {
	ctx := context.Background()
	result, err := r.db.Agent.Query().Where(agent.ID(agentId)).QueryStore().
		Select(merchantstore.FieldID, merchantstore.FieldName).
		QueryOrders().
		Where(
			order.ID(id),
		).
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
							product.FieldWeight, product.FieldQuantity,
						)
						pq.WithMinor(
							func(mq *ent.ProductCategoryMinorQuery) {
								mq.Select(productcategoryminor.FieldCategory)
							},
						)
					},
				)
			},
		).
		WithAddress().
		WithPickup().
		WithStores(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(merchantstore.FieldID, merchantstore.FieldName)
			},
		).
		WithLogistic().
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) ReadByStoreOrderDetail(orderId int) ([]*ent.OrderDetail, error) {
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

func (r *repository) Insert(res *models.OrderPayload, params ...int) (*ent.Order, error) {
	switch res.Metadata.UserType {
	case "retailer", "supplier":
		return r.insertMerchantOrder(res)
	case "business", "individual":
		return r.insertCustomerOrder(res, params...)
	case "agent":
		return r.insertAgentOrder(res)
	default:
		return nil, nil
	}
}
func (r *repository) SaveOrderUpdate(id int, res *models.OrderPayload) (*ent.Order, error) {
	ctx := context.Background()
	detailIds, err := r.db.Order.Query().Where(order.ID(id)).QueryDetails().IDs(ctx)
	if err != nil {
		return nil, err
	}
	_, deleteErr := r.db.OrderDetail.Delete().Where(orderdetail.IDIn(detailIds...)).Exec(ctx)
	if deleteErr != nil {
		return nil, deleteErr
	}
	stores := r.getNewStores(res.Metadata.Products)
	if res.Metadata.Address != 0 {
		o, oErr := r.db.Order.UpdateOneID(id).
			SetCustomerID(res.Metadata.User).
			SetAddressID(res.Metadata.Address).
			AddStoreIDs(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
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
		o, oErr := r.db.Order.UpdateOneID(id).
			SetCustomerID(res.Metadata.User).
			SetPickupID(res.Metadata.Pickup).
			AddStoreIDs(stores...).
			SetAmount(res.Amount / 100).
			SetDeliveryFee(res.Metadata.DeliveryFee).
			SetReference(res.Reference).
			SetChannel(res.Channel).
			SetCurrency(res.Currency).
			SetPaidAt(res.PaidAt).
			SetOrderNumber(res.Metadata.OrderNumber).
			SetDeliveryMethod(order.DeliveryMethod(res.Metadata.DeliveryMethod)).
			SetPaymentMethod(order.PaymentMethod(res.Metadata.PaymentMethod)).
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
func (r *repository) Read(id int) (*ent.Order, error) {
	result, err := r.db.Order.Query().
		Where(order.ID(id)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.WithStore(
					func(msq *ent.MerchantStoreQuery) {
						msq.Select(merchantstore.FieldID, merchantstore.FieldName, merchantstore.FieldCoordinate)
					},
				)
				odq.WithProduct(
					func(pq *ent.ProductQuery) {
						pq.Select(
							product.FieldID, product.FieldName, product.FieldUnit, product.FieldImage,
							product.FieldWeight, product.FieldQuantity,
						)
						pq.WithMinor(
							func(mq *ent.ProductCategoryMinorQuery) {
								mq.Select(productcategoryminor.FieldCategory)
							},
						)
					},
				)
			},
		).
		WithCustomer(
			func(cq *ent.CustomerQuery) {
				cq.Select(customer.FieldID)
			},
		).
		WithAgent(
			func(aq *ent.AgentQuery) {
				aq.Select(agent.FieldID)
			},
		).
		WithMerchant(
			func(mq *ent.MerchantQuery) {
				mq.Select(merchant.FieldID)
			},
		).
		WithAddress().
		WithPickup().
		WithLogistic().
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) ReadAll() ([]*ent.Order, error) {
	ctx := context.Background()
	results, err := r.db.Order.Query().
		Order(ent.Desc(order.FieldCreatedAt)).
		WithDetails(
			func(odq *ent.OrderDetailQuery) {
				odq.Select(orderdetail.FieldAmount, orderdetail.FieldStatus)
			},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) Update(order *services.PaystackResponse) (*ent.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (r *repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}

func (r *repository) UpdateOrderDetailStatus(requests map[string]*gabs.Container) (*ent.Order, error) {
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
	if strings.Compare(requests["userType"].Data().(string), "asinyo") == 0 {
		return r.Read(oId)
	}
	if strings.Compare(requests["userType"].Data().(string), "agent") == 0 {
		return r.ReadByAgentStore(oId, mId)
	}
	return r.ReadByStore(oId, mId)
}

func (r *repository) UpdateOrderApprovalStatus(orderId int, status string) (*ent.Order, error) {
	_, err := r.db.Order.UpdateOneID(orderId).SetCustomerApproval(order.CustomerApproval(status)).Save(context.Background())
	if err != nil {
		return nil, err
	}
	return r.Read(orderId)
}
func (r *repository) ReadOrderStoreMerchants(orderId int) (*ent.Order, error) {
	result, err := r.db.Order.Query().Where(order.ID(orderId)).WithStores(
		func(msq *ent.MerchantStoreQuery) {
			msq.WithMerchant()
			msq.WithAgent()
		},
	).
		Only(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) checkOrderStatus(data []*ent.OrderDetail) order.Status {
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

func (r *repository) insertMerchantOrder(res *models.OrderPayload) (*ent.Order, error) {
	ctx := context.Background()
	c := r.db.Merchant.Query().Where(merchant.ID(res.Metadata.User)).OnlyX(ctx)
	stores := r.getNewStores(res.Metadata.Products)
	if res.Metadata.Address != 0 {
		addr := r.db.Address.Query().Where(address.ID(res.Metadata.Address)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetMerchant(c).
			SetAddress(addr).
			AddStoreIDs(stores...).
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
			AddStoreIDs(stores...).
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

func (r *repository) insertCustomerOrder(res *models.OrderPayload, params ...int) (*ent.Order, error) {
	ctx := context.Background()
	stores := r.getNewStores(res.Metadata.Products)
	if params != nil {
		if res.Metadata.Address != 0 {
			o, oErr := r.db.Order.Create().
				SetCustomerID(res.Metadata.User).
				SetAddressID(res.Metadata.Address).
				AddStoreIDs(stores...).
				SetPurchaseRequestID(params[0]).
				SetCustomerApproval(order.CustomerApprovalPending).
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
			result, err := r.db.Order.Query().Where(order.ID(o.ID)).WithCustomer().Only(ctx)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		if res.Metadata.Pickup != 0 {
			o, oErr := r.db.Order.Create().
				SetCustomerID(res.Metadata.User).
				SetPickupID(res.Metadata.Pickup).
				AddStoreIDs(stores...).
				SetPurchaseRequestID(params[0]).
				SetCustomerApproval(order.CustomerApprovalPending).
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
			result, err := r.db.Order.Query().Where(order.ID(o.ID)).WithCustomer().Only(ctx)
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, nil
	}
	if res.Metadata.Address != 0 {
		o, oErr := r.db.Order.Create().
			SetCustomerID(res.Metadata.User).
			SetAddressID(res.Metadata.Address).
			AddStoreIDs(stores...).
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
		result, err := r.db.Order.Query().Where(order.ID(o.ID)).WithCustomer().Only(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	if res.Metadata.Pickup != 0 {
		o, oErr := r.db.Order.Create().
			SetCustomerID(res.Metadata.User).
			SetPickupID(res.Metadata.Pickup).
			AddStoreIDs(stores...).
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
		result, err := r.db.Order.Query().Where(order.ID(o.ID)).WithCustomer().Only(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	return nil, nil
}

func (r *repository) insertAgentOrder(res *models.OrderPayload) (*ent.Order, error) {
	ctx := context.Background()
	c := r.db.Agent.Query().Where(agent.ID(res.Metadata.User)).OnlyX(ctx)
	stores := r.getNewStores(res.Metadata.Products)
	if res.Metadata.Address != 0 {
		addr := r.db.Address.Query().Where(address.ID(res.Metadata.Address)).OnlyX(ctx)
		o, oErr := r.db.Order.Create().
			SetAgent(c).
			SetAddress(addr).
			AddStoreIDs(stores...).
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
			AddStoreIDs(stores...).
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

func (r *repository) calculateAmount(product *services.ProductDetails) float64 {
	var amount float64
	if product.PromoPrice > 0 {
		amount = product.PromoPrice * float64(product.Quantity)
	} else {
		amount = product.Price * float64(product.Quantity)
	}
	return amount
}

func (r *repository) readCustomerOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasCustomerWith(customer.ID(id))).
		Order(ent.Desc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) readAgentOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasAgentWith(agent.ID(id))).
		Order(ent.Desc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) readMerchantOrders(id int) ([]*ent.Order, error) {
	results, err := r.db.Order.Query().
		Where(order.HasMerchantWith(merchant.ID(id))).
		Order(ent.Desc(order.FieldCreatedAt)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) insertOrderDetails(metadata *models.OrderPayloadMetadata, o *ent.Order) error {
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

func (r *repository) getNewStores(products []*services.ProductDetails) []int {
	stores := make([]int, 0)
	for _, p := range products {
		if lo.Contains(stores, p.Store) {
			continue
		}
		stores = append(stores, p.Store)
	}
	return stores
}
