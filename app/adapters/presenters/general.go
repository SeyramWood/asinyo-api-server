package presenters

type (
	ResponseWithTotalRecords struct {
		TotalRecords int
		Records      any
	}
)
