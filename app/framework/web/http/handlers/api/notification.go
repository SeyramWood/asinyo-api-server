package api

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/adapters/gateways"
	"github.com/SeyramWood/app/adapters/presenters"
	"github.com/SeyramWood/app/domain/services"
)

type notificationHandler struct {
	service gateways.DBNotificationService
}

func NewNotificationHandler(service gateways.DBNotificationService) *notificationHandler {
	return &notificationHandler{
		service: service,
	}
}

func (h *notificationHandler) FetchNewNotifications() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")
		
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		requestParams := &services.NotificationFilters{
			MainFilter: c.Query("mainFilter", ""),
			Filter:     c.Query("filter", ""),
			SortBy:     c.Query("sortBy", ""),
			OrderBy:    c.Query("orderBy", ""),
			UserType:   userType,
			UserId:     userId,
			Limit:      limit,
			Offset:     offset,
		}
		clientID := fmt.Sprintf("%s%s%d", strings.Split(c.Get("User-Agent"), " ")[0], c.Get("X-Real-IP"), userId)
		c.Context().SetBodyStreamWriter(
			func(w *bufio.Writer) {
				messageChan := h.service.Broker(
					"get_client", clientID,
					requestParams,
				)
				if messageChan == nil {
					messageChan = h.service.Broker("add_client", clientID, requestParams)
				}
				defer func() {
					_ = h.service.Broker("close_client", clientID, requestParams)
				}()
				for {
					select {
					case mg := <-messageChan.Message:
						fmt.Fprintf(w, "data: %v\nretry:5000\n\n", string(mg))
						err := w.Flush()
						if err != nil {
							break
						}
						// case <-time.After(20 * time.Second):
						// 	fmt.Fprintf(w, "retry:100000\n\n")
						// 	err := w.Flush()
						// 	if err != nil {
						// 		break
						// 	}

					}
				}
			},
		)

		return nil
	}
}
func (h *notificationHandler) FetchAllByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		limit, _ := strconv.Atoi(c.Query("limit", "0"))
		offset, _ := strconv.Atoi(c.Query("offset", "0"))
		requestParams := &services.NotificationFilters{
			MainFilter: c.Query("mainFilter", ""),
			Filter:     c.Query("filter", ""),
			SortBy:     c.Query("sortBy", ""),
			OrderBy:    c.Query("orderBy", ""),
			UserType:   userType,
			UserId:     userId,
			Limit:      limit,
			Offset:     offset,
		}
		results, err := h.service.FetchUserNotifications(requestParams)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.NotificationErrorResponse(err))
		}
		return c.JSON(presenters.NotificationsSuccessResponse(results))

	}

}
func (h *notificationHandler) MarkAsRead() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		notificationId, _ := c.ParamsInt("notificationId")
		timestamp := c.Query("readAt")
		_, err := h.service.MarkAsRead(userId, notificationId, userType, timestamp)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.NotificationErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"status":  true,
				"message": "Updated successfully.",
			},
		)

	}
}
func (h *notificationHandler) MarkSelectedAsRead() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userType := c.Params("userType")
		userId, _ := c.ParamsInt("userId")
		timestamp := c.Query("readAt")
		type requestType struct {
			Notifications []int
		}
		var request requestType
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.NotificationErrorResponse(err))
		}
		err = h.service.MarkSelectedAsRead(userId, request.Notifications, userType, timestamp)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.NotificationErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"status":  true,
				"message": "Updated successfully.",
			},
		)

	}
}
func (h *notificationHandler) RemoveByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, _ := c.ParamsInt("id")
		err := h.service.Remove(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.NotificationErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"data": fiber.Map{
					"status":  true,
					"message": "Deleted successfully.",
				},
			},
		)

	}

}
func (h *notificationHandler) RemoveSelected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		type requestType struct {
			Notifications []int
		}
		var request requestType
		err := c.BodyParser(&request)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenters.NotificationErrorResponse(err))
		}
		err = h.service.RemoveSelected(request.Notifications)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenters.NotificationErrorResponse(err))
		}
		return c.JSON(
			fiber.Map{
				"data": fiber.Map{
					"status":  true,
					"message": "Deleted successfully.",
				},
			},
		)

	}

}
func (h *notificationHandler) Create() fiber.Handler {
	// return func(c *fiber.Ctx) error {
	// 	var request models.Address
	// 	userType := c.Params("userType")
	// 	userId, _ := c.ParamsInt("userId")
	//
	// 	err := c.BodyParser(&request)
	//
	// 	if err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(presenters.MerchantErrorResponse(err))
	// 	}
	//
	// 	result, err := h.service.Create(&request, userId, userType)
	//
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(presenters.AddressErrorResponse(errors.New("error creating merchant")))
	// 	}
	// 	h.maps.ExecuteTask(result, "geocoding", "address")
	//
	// 	return c.JSON(presenters.AddressSuccessResponse(result))
	//
	// }
	return nil
}
