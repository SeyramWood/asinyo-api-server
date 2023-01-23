package logistic

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	entLogistic "github.com/SeyramWood/ent/logistic"
	"github.com/SeyramWood/ent/merchantstore"
	"github.com/SeyramWood/ent/order"
	"github.com/SeyramWood/ent/orderdetail"
	"github.com/samber/lo"
)

type repository struct {
	db *ent.Client
}

func NewLogisticRepo(db *database.Adapter) gateways.LogisticRepo {
	return &repository{
		db: db.DB,
	}
}

func (r repository) InsertResponse(
	orderNum string, storeId int, response *models.TookanPickupAndDeliveryTaskResponse,
) (*ent.Logistic, error) {
	ctx := context.Background()
	result, err := r.db.Logistic.Create().
		SetStoreID(storeId).
		SetTask(response).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Order.Update().Where(order.OrderNumber(orderNum)).AddLogistic(result).Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r repository) UpdateOrderStatus(token, status string) error {
	ctx := context.Background()
	result, err := r.db.Logistic.Query().Where(
		func(s *sql.Selector) {
			s.Where(sqljson.ValueEQ(entLogistic.FieldTask, token, sqljson.Path("job_token")))
		},
	).
		WithStore(
			func(msq *ent.MerchantStoreQuery) {
				msq.Select(merchantstore.FieldID)
			},
		).
		WithOrder(
			func(orq *ent.OrderQuery) {
				orq.Select(order.FieldID, order.FieldOrderNumber)
			},
		).
		Only(ctx)

	if err != nil {
		return err
	}
	fmt.Println(result.Edges.Order[0].ID, result.Edges.Order[0].OrderNumber, result.Edges.Store.ID)
	_, err = r.db.OrderDetail.Update().Where(
		orderdetail.HasOrderWith(
			func(ord *sql.Selector) {
				ord.Where(sql.InInts(orderdetail.OrderColumn, result.Edges.Order[0].ID))
			},
		),
		orderdetail.HasStoreWith(
			func(stq *sql.Selector) {
				stq.Where(sql.InInts(orderdetail.StoreColumn, result.Edges.Store.ID))
			},
		),
	).SetStatus(orderdetail.Status(status)).Save(ctx)
	if err != nil {
		return err
	}

	o, err := r.db.Order.Query().Where(order.ID(result.Edges.Order[0].ID)).WithDetails(func(odq *ent.OrderDetailQuery) {
		odq.Select(orderdetail.FieldStatus)
	}).Only(ctx)
	if err != nil {
		return err
	}

	// fmt.Println(o.Edges.Details)

	if (r.isOrderStatusFulfilled(o.Edges.Details)){
		_, err := o.Update().SetStatus(order.StatusFulfilled).Save(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}


func (r repository) UpdateResponse(response *models.TookanPickupAndDeliveryTaskResponse) (*ent.Logistic, error) {
	// ctx := context.Background()
	// _, err := r.db.Logistic.UpdateOneID(r.db.Order.Query().Where(order.OrderNumber(response.Deliveries[0].OrderID)).QueryLogistic().OnlyIDX(ctx)).SetTasks(response).Save(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	return nil, nil
}
func (r repository) UpdateOrderDeliveryTask(orderId string, storeId int) error {
	ctx := context.Background()
	o := r.db.Order.Query().Where(order.OrderNumber(orderId)).OnlyX(ctx)
	newIds := []int{storeId}
	newIds = append(newIds, o.StoreTasksCreated...)
	if _, err := o.Update().SetStoreTasksCreated(newIds).Save(context.Background()); err != nil {
		return err
	}
	return nil
}
func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}

func (r repository) isOrderStatusFulfilled(data []*ent.OrderDetail) bool {
	delivered := lo.CountBy[*ent.OrderDetail](
		data, func(d *ent.OrderDetail) bool {
			return d.Status == "delivered"
		},
	)
	if delivered == len(data) {
		return true
	}
	return false
}
