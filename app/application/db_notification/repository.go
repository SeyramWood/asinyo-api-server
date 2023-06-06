package db_notification

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/samber/lo"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/application"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/app/domain/services"
	"github.com/SeyramWood/app/framework/database"
	"github.com/SeyramWood/ent"
	"github.com/SeyramWood/ent/admin"
	"github.com/SeyramWood/ent/agent"
	"github.com/SeyramWood/ent/customer"
	"github.com/SeyramWood/ent/merchant"
	"github.com/SeyramWood/ent/notification"
	"github.com/SeyramWood/ent/order"
)

type repository struct {
	db *ent.Client
}

func NewDBNotificationRepo(db *database.Adapter) gateways.DBNotificationRepo {
	return &repository{db.DB}
}

func (r *repository) Insert(msg *services.DBNotificationMessage) (*ent.Notification, error) {
	result, err := r.db.Notification.
		Create().
		AddAdminIDs(msg.AdminIDs...).
		AddMerchantIDs(msg.MerchantIDs...).
		AddAgentIDs(msg.AgentIDs...).
		AddCustomerIDs(msg.CustomerIDs...).
		SetSubjectType(msg.SubjectType).
		SetSubjectID(msg.SubjectId).
		SetCreatorType(msg.CreatorType).
		SetCreatorID(msg.CreatorId).
		SetEvent(msg.Event).
		SetActivity(msg.Activity).
		SetDescription(msg.Description).
		SetData(
			&struct {
				Data any `json:"data"`
			}{
				Data: msg.Data,
			},
		).
		Save(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed creating notification: %w", err)
	}
	return result, nil
}

func (r *repository) ReadNotification() (*ent.Notification, error) {
	result, err := r.db.Notification.Query().
		Order(ent.Desc(order.FieldCreatedAt)).
		First(context.Background())
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r *repository) ReadNotifications(limit, offset int) ([]*ent.Notification, error) {
	ctx := context.Background()
	results, err := r.db.Notification.Query().
		Order(ent.Desc(order.FieldCreatedAt)).
		Limit(limit).
		Offset(offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *repository) ReadNewNotifications() ([]*ent.Notification, error) {
	ctx := context.Background()
	results, err := r.db.Notification.Query().
		Order(ent.Desc(order.FieldCreatedAt)).Limit(15).
		Where(
			func(s *sql.Selector) {
				s.Where(sql.ExprP("created_at > now() - interval 24 hour"))
			},
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *repository) ReadUserNotifications(params *services.NotificationFilters) (
	[]*ent.Notification, error,
) {
	switch params.UserType {
	case "retailer", "supplier":
		return r.readMerchant(params)
	case "business", "individual":
		return r.readCustomer(params)
	case "agent":
		return r.readAgent(params)
	default:
		return nil, nil
	}
}
func (r *repository) ReadAdminNotifications(params *services.NotificationFilters) ([]*ent.Notification, error) {
	var results []*ent.Notification
	var err error
	permissions, err := application.ReadAdminPermissions(r.db, params.UserId)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if isAllActions := lo.Contains[string](permissions, "all_actions"); isAllActions {
		query := r.db.Admin.Query().QueryNotifications()
		if params.MainFilter != "" {
			if params.Filter == "unread" {
				if params.OrderBy == "Desc" {
					if params.SortBy == "date" {
						results, err = query.
							Where(notification.Event(params.MainFilter)).
							Where(
								notification.Or(
									notification.AdminReadAtIsNil(),
									notification.Not(
										func(s *sql.Selector) {
											s.Where(
												sqljson.StringContains(
													notification.FieldAdminReadAt,
													fmt.Sprintf("%d", params.UserId),
												),
											)
										},
									),
								),
							).
							Order(ent.Desc(order.FieldCreatedAt)).
							Limit(params.Limit).
							Offset(params.Offset).
							All(ctx)
						if err != nil {
							return nil, err
						}
						return results, nil
					}

				}
				if params.OrderBy == "Asc" {
					if params.SortBy == "date" {
						results, err = query.
							Where(notification.Event(params.MainFilter)).
							Where(
								notification.Or(
									notification.AdminReadAtIsNil(),
									notification.Not(
										func(s *sql.Selector) {
											s.Where(
												sqljson.StringContains(
													notification.FieldAdminReadAt,
													fmt.Sprintf("%d", params.UserId),
												),
											)
										},
									),
								),
							).
							Order(ent.Asc(order.FieldCreatedAt)).
							Limit(params.Limit).
							Offset(params.Offset).
							All(ctx)
						if err != nil {
							return nil, err
						}
						return results, nil
					}

				}

			}

			if params.OrderBy == "Desc" {
				if params.SortBy == "date" {
					results, err = query.
						Where(notification.Event("NewOrder")).
						Order(ent.Desc(order.FieldCreatedAt)).
						Limit(params.Limit).
						Offset(params.Offset).
						All(ctx)
					if err != nil {
						return nil, err
					}
					return results, nil
				}

			}
			if params.OrderBy == "Asc" {
				if params.SortBy == "date" {
					results, err = query.
						Where(notification.Event(params.MainFilter)).
						Order(ent.Asc(order.FieldCreatedAt)).
						Limit(params.Limit).
						Offset(params.Offset).
						All(ctx)
					if err != nil {
						return nil, err
					}
					return results, nil
				}

			}

		}

		if params.Filter == "unread" {
			if params.OrderBy == "Desc" {
				if params.SortBy == "date" {
					results, err = query.
						Where(
							notification.Or(
								notification.AdminReadAtIsNil(),
								notification.Not(
									func(s *sql.Selector) {
										s.Where(
											sqljson.StringContains(
												notification.FieldAdminReadAt,
												fmt.Sprintf("%d", params.UserId),
											),
										)
									},
								),
							),
						).
						Order(ent.Desc(order.FieldCreatedAt)).
						Limit(params.Limit).
						Offset(params.Offset).
						All(ctx)
					if err != nil {
						return nil, err
					}
					return results, nil
				}

			}
			if params.OrderBy == "Asc" {
				if params.SortBy == "date" {
					results, err = query.
						Where(
							notification.Or(
								notification.AdminReadAtIsNil(),
								notification.Not(
									func(s *sql.Selector) {
										s.Where(
											sqljson.StringContains(
												notification.FieldAdminReadAt,
												fmt.Sprintf("%d", params.UserId),
											),
										)
									},
								),
							),
						).
						Order(ent.Asc(order.FieldCreatedAt)).
						Limit(params.Limit).
						Offset(params.Offset).
						All(ctx)
					if err != nil {
						return nil, err
					}
					return results, nil
				}

			}
		}

		if params.OrderBy == "Desc" {
			if params.SortBy == "date" {
				results, err = query.
					Order(ent.Desc(order.FieldCreatedAt)).
					Limit(params.Limit).
					Offset(params.Offset).
					All(ctx)
				if err != nil {
					return nil, err
				}
				return results, nil
			}

		}
		if params.OrderBy == "Asc" {
			if params.SortBy == "date" {
				results, err = query.
					Order(ent.Asc(order.FieldCreatedAt)).
					Limit(params.Limit).
					Offset(params.Offset).
					All(ctx)
				if err != nil {
					return nil, err
				}
				return results, nil
			}

		}

		results, err = query.
			Order(ent.Desc(order.FieldCreatedAt)).
			Limit(params.Limit).
			Offset(params.Offset).
			All(ctx)
		if err != nil {
			return nil, err
		}
		return results, nil
	}

	// Default Query
	query := r.db.Admin.Query().Where(admin.ID(params.UserId)).QueryNotifications()
	if params.MainFilter != "" {
		if params.Filter == "unread" {
			if params.OrderBy == "Desc" {
				if params.SortBy == "date" {
					results, err = query.
						Where(notification.Event(params.MainFilter)).
						Where(
							notification.Or(
								notification.AdminReadAtIsNil(),
								notification.Not(
									func(s *sql.Selector) {
										s.Where(
											sqljson.StringContains(
												notification.FieldAdminReadAt,
												fmt.Sprintf("%d", params.UserId),
											),
										)
									},
								),
							),
						).
						Order(ent.Desc(order.FieldCreatedAt)).
						Limit(params.Limit).
						Offset(params.Offset).
						All(ctx)
					if err != nil {
						return nil, err
					}
					log.Println(results)
					return results, nil
				}

			}
			if params.OrderBy == "Asc" {
				if params.SortBy == "date" {
					results, err = query.
						Where(notification.Event(params.MainFilter)).
						Where(
							notification.Or(
								notification.AdminReadAtIsNil(),
								notification.Not(
									func(s *sql.Selector) {
										s.Where(
											sqljson.StringContains(
												notification.FieldAdminReadAt,
												fmt.Sprintf("%d", params.UserId),
											),
										)
									},
								),
							),
						).
						Order(ent.Asc(order.FieldCreatedAt)).
						Limit(params.Limit).
						Offset(params.Offset).
						All(ctx)
					if err != nil {
						return nil, err
					}
					return results, nil
				}

			}

		}

		if params.OrderBy == "Desc" {
			if params.SortBy == "date" {
				results, err = query.
					Where(notification.Event(params.MainFilter)).
					Order(ent.Desc(order.FieldCreatedAt)).
					Limit(params.Limit).
					Offset(params.Offset).
					All(ctx)
				if err != nil {
					return nil, err
				}
				return results, nil
			}

		}
		if params.OrderBy == "Asc" {
			if params.SortBy == "date" {
				results, err = query.
					Where(notification.Event(params.MainFilter)).
					Order(ent.Asc(order.FieldCreatedAt)).
					Limit(params.Limit).
					Offset(params.Offset).
					All(ctx)
				if err != nil {
					return nil, err
				}
				return results, nil
			}

		}

	}

	if params.Filter == "unread" {
		if params.OrderBy == "Desc" {
			if params.SortBy == "date" {
				results, err = query.
					Where(
						notification.Or(
							notification.AdminReadAtIsNil(),
							notification.Not(
								func(s *sql.Selector) {
									s.Where(
										sqljson.StringContains(
											notification.FieldAdminReadAt,
											fmt.Sprintf("%d", params.UserId),
										),
									)
								},
							),
						),
					).
					Order(ent.Desc(order.FieldCreatedAt)).
					Limit(params.Limit).
					Offset(params.Offset).
					All(ctx)
				if err != nil {
					return nil, err
				}
				return results, nil
			}

		}
		if params.OrderBy == "Asc" {
			if params.SortBy == "date" {
				results, err = query.
					Where(
						notification.Or(
							notification.AdminReadAtIsNil(),
							notification.Not(
								func(s *sql.Selector) {
									s.Where(
										sqljson.StringContains(
											notification.FieldAdminReadAt,
											fmt.Sprintf("%d", params.UserId),
										),
									)
								},
							),
						),
					).
					Order(ent.Asc(order.FieldCreatedAt)).
					Limit(params.Limit).
					Offset(params.Offset).
					All(ctx)
				if err != nil {
					return nil, err
				}
				return results, nil
			}

		}

	}
	if params.OrderBy == "Desc" {
		if params.SortBy == "date" {
			results, err = query.
				Order(ent.Desc(order.FieldCreatedAt)).
				Limit(params.Limit).
				Offset(params.Offset).
				All(ctx)
			if err != nil {
				return nil, err
			}
			return results, nil
		}

	}
	if params.OrderBy == "Asc" {
		if params.SortBy == "date" {
			results, err = query.
				Order(ent.Asc(order.FieldCreatedAt)).
				Limit(params.Limit).
				Offset(params.Offset).
				All(ctx)
			if err != nil {
				return nil, err
			}
			return results, nil
		}
	}
	results, err = query.
		Order(ent.Desc(order.FieldCreatedAt)).
		Limit(params.Limit).
		Offset(params.Offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *repository) MarkAsRead(userId, notificationId int, userType, timestamp string) (*ent.Notification, error) {

	if userType == "asinyo" {
		ctx := context.Background()
		prevReads := r.db.Notification.Query().Where(notification.ID(notificationId)).OnlyX(ctx)
		newReads := append(
			[]*models.AdminRead{
				{
					ID:     userId,
					ReadAt: timestamp,
				},
			}, prevReads.AdminReadAt...,
		)
		result, err := r.db.Notification.UpdateOneID(notificationId).
			SetAdminReadAt(newReads).Save(ctx)
		if err != nil {
			return nil, err
		}
		return result, nil
	}

	if userType == "supplier" || userType == "retailer" {
		result, err := r.db.Notification.UpdateOneID(notificationId).
			SetMerchantReadAt(timestamp).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	if userType == "business" || userType == "individual" {
		result, err := r.db.Notification.UpdateOneID(notificationId).
			SetCustomerReadAt(timestamp).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	if userType == "agent" {
		result, err := r.db.Notification.UpdateOneID(notificationId).
			SetAgentReadAt(timestamp).
			Save(context.Background())
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, nil
}
func (r *repository) MarkSelectedAsRead(userId int, notificationIds []int, userType, timestamp string) error {
	if userType == "asinyo" {
		ctx := context.Background()
		newReads := []*models.AdminRead{
			{
				ID:     userId,
				ReadAt: timestamp,
			},
		}
		// wg := &sync.WaitGroup{}
		for _, id := range notificationIds {
			// wg.Add(1)
			log.Println(id)
			var myErr error
			prevReads := r.db.Notification.Query().Where(notification.ID(id)).OnlyX(ctx)
			reads := append(newReads, prevReads.AdminReadAt...)
			_, err := r.db.Notification.UpdateOneID(id).
				SetAdminReadAt(reads).Save(ctx)
			if err != nil {
				myErr = err
			}
			// go func(newReads []*models.AdminRead, w *sync.WaitGroup, ctx context.Context, myErr error) {
			// 	defer w.Done()
			// 	prevReads := r.db.Notification.Query().Where(notification.ID(id)).OnlyX(ctx)
			// 	reads := append(newReads, prevReads.AdminReadAt...)
			// 	_, err := r.db.Notification.UpdateOneID(id).
			// 		SetAdminReadAt(reads).Save(ctx)
			// 	if err != nil {
			// 		myErr = err
			// 	}
			// }(newReads, wg, ctx, myErr)
			if myErr != nil {
				log.Println(myErr)
				return myErr
			}
		}
		// wg.Wait()
		return nil
	}

	if userType == "supplier" || userType == "retailer" {
		_, err := r.db.Notification.Update().Where(notification.IDIn(notificationIds...)).
			SetMerchantReadAt(timestamp).
			Save(context.Background())
		if err != nil {
			return err
		}
		return nil
	}
	if userType == "business" || userType == "individual" {
		_, err := r.db.Notification.Update().Where(notification.IDIn(notificationIds...)).
			SetCustomerReadAt(timestamp).
			Save(context.Background())
		if err != nil {
			return err
		}
		return nil
	}
	if userType == "agent" {
		_, err := r.db.Notification.Update().Where(notification.IDIn(notificationIds...)).
			SetAgentReadAt(timestamp).
			Save(context.Background())
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (r *repository) RemoveSelected(ids []int) error {
	_, err := r.db.Notification.Delete().Where(notification.IDIn(ids...)).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) Delete(id int) error {
	err := r.db.Notification.DeleteOneID(id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) readMerchant(params *services.NotificationFilters) (
	[]*ent.Notification, error,
) {
	ctx := context.Background()
	if params.MainFilter != "" {
		results, err := r.db.Merchant.Query().Where(merchant.ID(params.UserId)).QueryNotifications().
			Where(notification.Event(params.MainFilter)).
			Order(ent.Desc(order.FieldCreatedAt)).
			Limit(params.Limit).
			Offset(params.Offset).
			All(ctx)
		if err != nil {
			return nil, err
		}
		return results, nil
	}
	results, err := r.db.Merchant.Query().Where(merchant.ID(params.UserId)).QueryNotifications().
		Order(ent.Desc(order.FieldCreatedAt)).
		Limit(params.Limit).
		Offset(params.Offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *repository) readAgent(params *services.NotificationFilters) (
	[]*ent.Notification, error,
) {
	ctx := context.Background()
	if params.MainFilter != "" {
		results, err := r.db.Agent.Query().Where(agent.ID(params.UserId)).QueryNotifications().
			Where(notification.Event(params.MainFilter)).
			Order(ent.Desc(order.FieldCreatedAt)).
			Limit(params.Limit).
			Offset(params.Offset).
			All(ctx)
		if err != nil {
			return nil, err
		}
		return results, nil
	}
	results, err := r.db.Agent.Query().Where(agent.ID(params.UserId)).QueryNotifications().
		Order(ent.Desc(order.FieldCreatedAt)).
		Limit(params.Limit).
		Offset(params.Offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (r *repository) readCustomer(params *services.NotificationFilters) (
	[]*ent.Notification, error,
) {
	ctx := context.Background()
	if params.MainFilter != "" {
		results, err := r.db.Customer.Query().Where(customer.ID(params.UserId)).QueryNotifications().
			Where(notification.Event(params.MainFilter)).
			Order(ent.Desc(order.FieldCreatedAt)).
			Limit(params.Limit).
			Offset(params.Offset).
			All(ctx)
		if err != nil {
			return nil, err
		}
		return results, nil
	}
	results, err := r.db.Customer.Query().Where(customer.ID(params.UserId)).QueryNotifications().
		Order(ent.Desc(order.FieldCreatedAt)).
		Limit(params.Limit).
		Offset(params.Offset).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}
