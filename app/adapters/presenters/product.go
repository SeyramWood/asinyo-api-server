package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	ProductStore struct {
		ID           int    `json:"id"`
		BusinessName string `json:"businessName"`
	}
	MerchantResponse struct {
		Username  string        `json:"username"`
		LastName  string        `json:"lastName"`
		OtherName string        `json:"otherName"`
		Store     *ProductStore `json:"store"`
	}
	ProductWithMerchant struct {
		ID          int               `json:"id"`
		MajorID     int               `json:"categoryMajorId"`
		MinorID     int               `json:"categoryMinorId"`
		Name        string            `json:"name"`
		Unit        string            `json:"unit"`
		Weight      int               `json:"weight"`
		Quantity    int               `json:"quantity"`
		Price       float64           `json:"price"`
		PromoPrice  float64           `json:"promoPrice"`
		Description string            `json:"description"`
		Image       string            `json:"image"`
		Minor       string            `json:"categoryMinor"`
		Major       string            `json:"categoryMajor"`
		CreatedAt   time.Time         `json:"created_at"`
		UpdatedAt   time.Time         `json:"updated_at"`
		Merchant    *MerchantResponse `json:"merchant"`
	}

	ProductWithStore struct {
		ID          int           `json:"id"`
		MajorID     int           `json:"categoryMajorId"`
		MinorID     int           `json:"categoryMinorId"`
		Name        string        `json:"name"`
		Unit        string        `json:"unit"`
		Weight      int           `json:"weight"`
		Quantity    int           `json:"quantity"`
		Price       float64       `json:"price"`
		PromoPrice  float64       `json:"promoPrice"`
		Description string        `json:"description"`
		Image       string        `json:"image"`
		Minor       string        `json:"categoryMinor"`
		Major       string        `json:"categoryMajor"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
		Store       *ProductStore `json:"store"`
	}

	ProductCategoryWithMerchant struct {
		ID        int                    `json:"id"`
		Category  string                 `json:"category"`
		Slug      string                 `json:"slug"`
		CreatedAt time.Time              `json:"created_at"`
		UpdatedAt time.Time              `json:"updated_at"`
		Product   []*ProductWithMerchant `json:"products"`
	}
)

func ProductWithMerchantResponse(data *ent.Product) *fiber.Map {
	return successResponse(formatProductWithMerchant(data))
}

func ProductsWithStoreResponse(data []*ent.Product) *fiber.Map {
	var response []*ProductWithStore
	for _, v := range data {
		response = append(
			response, &ProductWithStore{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Weight:      int(v.Weight),
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  v.PromoPrice,
				Description: v.Description,
				Image:       v.Image,
				Major:       v.Edges.Major.Category,
				Minor:       v.Edges.Minor.Category,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Store: func() *ProductStore {
					if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err == nil {
						return &ProductStore{
							ID:           s.ID,
							BusinessName: s.Name,
						}
					}
					return nil
				}(),
			},
		)
	}
	return successResponse(response)

}

func ProductsWithMerchantResponse(data []*ent.Product) *fiber.Map {
	return successResponse(formatProductsWithMerchant(data))
}

func ProductsBestSellerResponse(data []*ent.Product) *fiber.Map {
	var response []*ProductWithMerchant

	for _, v := range data {

		response = append(
			response, &ProductWithMerchant{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Weight:      int(v.Weight),
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  v.PromoPrice,
				Description: v.Description,
				Image:       v.Image,
				Major:       v.Edges.Major.Category,
				Minor:       v.Edges.Minor.Category,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Merchant: &MerchantResponse{
					Username:  v.Edges.Merchant.Username,
					LastName:  v.Edges.Merchant.Edges.Retailer.LastName,
					OtherName: v.Edges.Merchant.Edges.Retailer.OtherName,
					Store: func() *ProductStore {
						if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err != nil {
							return &ProductStore{
								ID:           s.ID,
								BusinessName: s.Name,
							}
						}
						return nil
					}(),
				},
			},
		)

	}

	return successResponse(response)

}

func ProductsCategoryMajorWithMerchantResponse(data []*ent.ProductCategoryMajor) *fiber.Map {
	var response []*ProductCategoryWithMerchant
	for _, cat := range data {
		response = append(
			response, &ProductCategoryWithMerchant{
				ID:        cat.ID,
				Category:  cat.Category,
				Slug:      cat.Slug,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
				Product:   formatProductsWithMerchant(cat.Edges.Products),
			},
		)
	}
	return successResponse(response)
}

func ProductsCategoryMinorWithMerchantResponse(data []*ent.ProductCategoryMinor) *fiber.Map {
	var response []*ProductCategoryWithMerchant
	for _, cat := range data {
		response = append(
			response, &ProductCategoryWithMerchant{
				ID:        cat.ID,
				Category:  cat.Category,
				Slug:      cat.Slug,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
				Product:   formatProductsWithMerchant(cat.Edges.Products),
			},
		)
	}
	return successResponse(response)
}

func formatProductsWithMerchant(data []*ent.Product) []*ProductWithMerchant {
	var response []*ProductWithMerchant
	for _, v := range data {
		response = append(
			response, &ProductWithMerchant{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Weight:      int(v.Weight),
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  v.PromoPrice,
				Description: v.Description,
				Image:       v.Image,
				Major:       v.Edges.Major.Category,
				Minor:       v.Edges.Minor.Category,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Merchant: func() *MerchantResponse {
					if s, err := v.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
						return &MerchantResponse{
							Username:  v.Edges.Merchant.Username,
							LastName:  s.LastName,
							OtherName: s.OtherName,
							Store: func() *ProductStore {
								if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err == nil {
									return &ProductStore{
										ID:           s.ID,
										BusinessName: s.Name,
									}
								}
								return nil
							}(),
						}
					}
					if r, err := v.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
						return &MerchantResponse{
							Username:  v.Edges.Merchant.Username,
							LastName:  r.LastName,
							OtherName: r.OtherName,
							Store: func() *ProductStore {
								if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err == nil {
									return &ProductStore{
										ID:           s.ID,
										BusinessName: s.Name,
									}
								}
								return nil
							}(),
						}
					}
					return nil
				}(),
			},
		)
	}
	return response
}

func formatProductWithMerchant(data *ent.Product) *ProductWithMerchant {

	return &ProductWithMerchant{
		ID:          data.ID,
		MajorID:     data.Edges.Major.ID,
		MinorID:     data.Edges.Minor.ID,
		Name:        data.Name,
		Unit:        data.Unit,
		Weight:      int(data.Weight),
		Quantity:    int(data.Quantity),
		Price:       data.Price,
		PromoPrice:  data.PromoPrice,
		Description: data.Description,
		Image:       data.Image,
		Major:       data.Edges.Major.Category,
		Minor:       data.Edges.Minor.Category,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Merchant: func() *MerchantResponse {
			if s, err := data.Edges.Merchant.Edges.SupplierOrErr(); err == nil {
				return &MerchantResponse{
					Username:  data.Edges.Merchant.Username,
					LastName:  s.LastName,
					OtherName: s.OtherName,
					Store: func() *ProductStore {
						if s, err := data.Edges.Merchant.Edges.StoreOrErr(); err == nil {
							return &ProductStore{
								ID:           s.ID,
								BusinessName: s.Name,
							}
						}
						return nil
					}(),
				}
			}
			if r, err := data.Edges.Merchant.Edges.RetailerOrErr(); err == nil {
				return &MerchantResponse{
					Username:  data.Edges.Merchant.Username,
					LastName:  r.LastName,
					OtherName: r.OtherName,
					Store: func() *ProductStore {
						if s, err := data.Edges.Merchant.Edges.StoreOrErr(); err == nil {
							return &ProductStore{
								ID:           s.ID,
								BusinessName: s.Name,
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}

}

func ProductErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
