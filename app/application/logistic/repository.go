package logistic

import (
	"context"
	"fmt"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/order"
)

type repository struct {
	db *ent.Client
}

func NewLogisticRepo(db *database.Adapter) gateways.LogisticRepo {
	return &repository{
		db: db.DB,
	}
}

func (r repository) InsertResponse(response *models.TookanMultiTaskResponse) (*ent.Logistic, error) {
	ctx := context.Background()
	result, err := r.db.Logistic.Create().
		SetTrackingLink(response.Deliveries[0].ResultTrackingLink).
		SetTasks(response).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.db.Order.Update().Where(order.OrderNumber(response.Deliveries[0].OrderID)).AddLogistic(result).Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (r repository) UpdateResponse(response *models.TookanMultiTaskResponse) (*ent.Logistic, error) {
	ctx := context.Background()
	_, err := r.db.Logistic.UpdateOneID(r.db.Order.Query().Where(order.OrderNumber(response.Deliveries[0].OrderID)).QueryLogistic().OnlyIDX(ctx)).SetTasks(response).Save(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (r repository) UpdateOrderDeliveryTask(orderId string, ids []int) error {
	ctx := context.Background()
	o := r.db.Order.Query().Where(order.OrderNumber(orderId)).OnlyX(ctx)
	newIds := ids
	newIds = append(newIds, o.StoreTasksCreated...)
	fmt.Println(ids)
	fmt.Println(newIds)
	if _, err := r.db.Order.Update().Where(order.OrderNumber(orderId)).SetStoreTasksCreated(newIds).Save(context.Background()); err != nil {
		return err
	}
	return nil
}
func (r repository) Delete(id string) error {
	// TODO implement me
	panic("implement me")
}
