package models

type (
	LogisticConfiguration struct {
		Logistics []any  `json:"logistics"`
		Current   string `json:"current"`
	}
	PricingModelConfiguration struct {
		Models  []any  `json:"models"`
		Current string `json:"current"`
	}
)
