package logistic

import (
	"context"

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
				orq.Select(order.FieldID)
			},
		).
		Only(ctx)

	if err != nil {
		return err
	}
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
	if _, err := r.db.Order.Update().Where(order.OrderNumber(orderId)).SetStoreTasksCreated(newIds).Save(context.Background()); err != nil {
		return err
	}
	return nil
}
func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}
