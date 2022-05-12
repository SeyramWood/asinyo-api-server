package models

type (
	ProductCategoryMajor struct {
		Category string `json:"category" validate:"required|string"`
	}
	ProductCategoryMinor struct {
		CategoryMajor int    `json:"categoryMajor" validate:"required"`
		Category      string `json:"category" validate:"required|string"`
	}
	Product struct {
		CategoryMajor int     `json:"categoryMajor" validate:"required"`
		CategoryMinor int     `json:"categoryMinor" validate:"required"`
		Quantity      int     `json:"quantity" validate:"required"`
		Unit          int     `json:"unit" validate:"required"`
		Name          string  `json:"name" validate:"required|string"`
		Price         float64 `json:"price" validate:"required"`
		PromoPrice    float64 `json:"promoPrice" validate:"required"`
		Description   string  `json:"description" validate:"required|string"`
		Image         []byte  `json:"image"`
	}
)
