package models

type (
	ProductCategoryMajor struct {
		Category string `json:"category" validate:"required|string"`
	}
	ProductCategoryMinor struct {
		CategoryMajor int    `json:"categoryMajor" validate:"required"`
		Category      string `json:"category" validate:"required|string"`
		Image         []byte `json:"image" form:"image"`
	}
	Product struct {
		CategoryMajor int     `json:"categoryMajor" form:"categoryMajor" validate:"required"`
		CategoryMinor int     `json:"categoryMinor" form:"categoryMinor" validate:"required"`
		Merchant      int     `json:"merchant" form:"merchant" validate:"required"`
		Quantity      int     `json:"quantity" form:"quantity" validate:"required"`
		Weight        int     `json:"weight" form:"weight" validate:"required"`
		Unit          string  `json:"unit" form:"unit" validate:"required|string"`
		Name          string  `json:"name" form:"name" validate:"required|string"`
		Price         float64 `json:"price" form:"price" validate:"required"`
		PromoPrice    float64 `json:"promoPrice" form:"promoPrice" validate:"required"`
		Description   string  `json:"description" form:"description" validate:"required"`
		Image         []byte  `json:"image" form:"image"`
	}
)
