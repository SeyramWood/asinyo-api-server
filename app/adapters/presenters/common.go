package presenters

type (
	PaginationResponse struct {
		Count int `json:"count,omitempty"`
		Data  any `json:"data"`
	}
	RecordCount struct {
		Total  int `json:"total"`
		Recent int `json:"recent"`
	}
	MerchantRecordCount struct {
		Supplier *RecordCount `json:"supplier"`
		Retailer *RecordCount `json:"retailer"`
	}
	CustomerRecordCount struct {
		Business   *RecordCount `json:"business"`
		Individual *RecordCount `json:"individual"`
	}
	DashboardRecordCount struct {
		Customers *CustomerRecordCount `json:"customers"`
		Merchants *MerchantRecordCount `json:"merchants"`
		Agents    *RecordCount         `json:"agents"`
		Orders    *RecordCount         `json:"orders"`
	}
)
