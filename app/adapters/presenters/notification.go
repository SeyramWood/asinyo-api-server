package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	AdminRead struct {
		ID     int    `json:"id"`
		ReadAt string `json:"readAt"`
	}
	Notification struct {
		ID           int          `json:"id"`
		SubjectId    int          `json:"subjectId"`
		CreatorId    int          `json:"creatorId"`
		Event        string       `json:"event"`
		Activity     string       `json:"activity"`
		Description  string       `json:"description"`
		Data         any          `json:"data"`
		CustomerRead string       `json:"customerRead"`
		AgentRead    string       `json:"agentRead"`
		MerchantRead string       `json:"merchantRead"`
		AdminRead    []*AdminRead `json:"adminRead"`
		CreatedAt    time.Time    `json:"created_at"`
		UpdatedAt    time.Time    `json:"updated_at"`
	}
)

func NotificationsSuccessResponse(data []*ent.Notification) *fiber.Map {
	var response []*Notification
	for _, noti := range data {
		response = append(
			response, &Notification{
				ID:           noti.ID,
				SubjectId:    noti.SubjectID,
				CreatorId:    noti.CreatorID,
				Event:        noti.Event,
				Activity:     noti.Activity,
				Description:  noti.Description,
				Data:         noti.Data.Data,
				CustomerRead: noti.CustomerReadAt,
				AgentRead:    noti.AgentReadAt,
				MerchantRead: noti.MerchantReadAt,
				AdminRead: func() []*AdminRead {
					if noti.AdminReadAt == nil {
						return nil
					}
					var reads []*AdminRead
					for _, r := range noti.AdminReadAt {
						reads = append(
							reads, &AdminRead{
								ID:     r.ID,
								ReadAt: r.ReadAt,
							},
						)
					}
					return reads
				}(),
				CreatedAt: noti.CreatedAt,
				UpdatedAt: noti.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}

func FormatNotificationResponse(data *ent.Notification) *Notification {
	return &Notification{
		ID:           data.ID,
		SubjectId:    data.SubjectID,
		CreatorId:    data.CreatorID,
		Event:        data.Event,
		Activity:     data.Activity,
		Description:  data.Description,
		Data:         data.Data.Data,
		CustomerRead: data.CustomerReadAt,
		AgentRead:    data.AgentReadAt,
		MerchantRead: data.MerchantReadAt,
		AdminRead: func() []*AdminRead {
			if data.AdminReadAt == nil {
				return nil
			}
			var reads []*AdminRead
			for _, r := range data.AdminReadAt {
				reads = append(
					reads, &AdminRead{
						ID:     r.ID,
						ReadAt: r.ReadAt,
					},
				)
			}
			return reads
		}(),
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
func FormatNotificationsResponse(data []*ent.Notification) []*Notification {
	var response []*Notification
	for _, noti := range data {
		response = append(
			response, &Notification{
				ID:           noti.ID,
				SubjectId:    noti.SubjectID,
				CreatorId:    noti.CreatorID,
				Event:        noti.Event,
				Activity:     noti.Activity,
				Description:  noti.Description,
				Data:         noti.Data.Data,
				CustomerRead: noti.CustomerReadAt,
				AgentRead:    noti.AgentReadAt,
				MerchantRead: noti.MerchantReadAt,
				AdminRead: func() []*AdminRead {
					if noti.AdminReadAt == nil {
						return nil
					}
					var reads []*AdminRead
					for _, r := range noti.AdminReadAt {
						reads = append(
							reads, &AdminRead{
								ID:     r.ID,
								ReadAt: r.ReadAt,
							},
						)
					}
					return reads
				}(),
				CreatedAt: noti.CreatedAt,
				UpdatedAt: noti.UpdatedAt,
			},
		)
	}
	return response
}

func NotificationErrorResponse(err interface{}) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"error":  err,
	}
}
