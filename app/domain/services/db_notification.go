package services

type (
	DBNotificationMessage struct {
		AdminIDs    []int
		MerchantIDs []int
		AgentIDs    []int
		CustomerIDs []int
		SubjectType string
		SubjectId   int
		CreatorType string
		CreatorId   int
		Event       string
		Activity    string
		Description string
		Data        any
	}
	NotificationFilters struct {
		MainFilter, Filter, SortBy, OrderBy, UserType string
		UserId, Limit, Offset                         int
	}
	NotificationClient struct {
		ID      string
		Params  *NotificationFilters
		Message chan []byte
	}
)
