package presenters

import (
	"sync"
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
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
	Product struct {
		ID          int               `json:"id"`
		MajorID     int               `json:"categoryMajorId"`
		MinorID     int               `json:"categoryMinorId"`
		Name        string            `json:"name"`
		Unit        string            `json:"unit"`
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

	AllProducts struct {
		ID          int           `json:"id"`
		MajorID     int           `json:"categoryMajorId"`
		MinorID     int           `json:"categoryMinorId"`
		Name        string        `json:"name"`
		Unit        string        `json:"unit"`
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

	CategoryMajor struct {
		ID        int        `json:"id"`
		Category  string     `json:"category"`
		Slug      string     `json:"slug"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Product   []*Product `json:"products"`
	}

	CategoryMinor struct {
		ID        int        `json:"id"`
		Category  string     `json:"category"`
		Slug      string     `json:"slug"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		Product   []*Product `json:"products"`
	}
)

func ProductSuccessResponse(data *ent.Product) *fiber.Map {

	return successResponse(&Product{
		ID:          data.ID,
		MajorID:     data.Edges.Major.ID,
		MinorID:     data.Edges.Minor.ID,
		Name:        data.Name,
		Unit:        data.Unit,
		Quantity:    int(data.Quantity),
		Price:       data.Price,
		PromoPrice:  *data.PromoPrice,
		Description: data.Description,
		Image:       data.Image,
		Major:       data.Edges.Major.Category,
		Minor:       data.Edges.Minor.Category,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Merchant: &MerchantResponse{
			Username:  data.Edges.Merchant.Username,
			LastName:  data.Edges.Merchant.Edges.Supplier.LastName,
			OtherName: data.Edges.Merchant.Edges.Supplier.OtherName,
			Store: func() *ProductStore {
				if s, err := data.Edges.Merchant.Edges.StoreOrErr(); err == nil {
					return &ProductStore{
						ID:           s.ID,
						BusinessName: s.Name,
					}
				}
				return nil
			}(),
		},
	})
}

func ProductRetailerMerchantSuccessResponse(data *ent.Product) *fiber.Map {

	return successResponse(&Product{
		ID:          data.ID,
		MajorID:     data.Edges.Major.ID,
		MinorID:     data.Edges.Minor.ID,
		Name:        data.Name,
		Unit:        data.Unit,
		Quantity:    int(data.Quantity),
		Price:       data.Price,
		PromoPrice:  *data.PromoPrice,
		Description: data.Description,
		Image:       data.Image,
		Major:       data.Edges.Major.Category,
		Minor:       data.Edges.Minor.Category,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Merchant: &MerchantResponse{
			Username:  data.Edges.Merchant.Username,
			LastName:  data.Edges.Merchant.Edges.Retailer.LastName,
			OtherName: data.Edges.Merchant.Edges.Retailer.OtherName,
			Store: func() *ProductStore {
				if s, err := data.Edges.Merchant.Edges.StoreOrErr(); err == nil {
					return &ProductStore{
						ID:           s.ID,
						BusinessName: s.Name,
					}
				}
				return nil
			}(),
		},
	})
}

func ProductSupplierResponse(data *ent.Product) *fiber.Map {
	return successResponse(&Product{
		ID:          data.ID,
		MajorID:     data.Edges.Major.ID,
		MinorID:     data.Edges.Minor.ID,
		Name:        data.Name,
		Unit:        data.Unit,
		Quantity:    int(data.Quantity),
		Price:       data.Price,
		PromoPrice:  *data.PromoPrice,
		Description: data.Description,
		Image:       data.Image,
		Major:       data.Edges.Major.Category,
		Minor:       data.Edges.Minor.Category,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Merchant: &MerchantResponse{
			Username:  data.Edges.Merchant.Username,
			LastName:  data.Edges.Merchant.Edges.Supplier.LastName,
			OtherName: data.Edges.Merchant.Edges.Supplier.OtherName,
			Store: func() *ProductStore {
				if s, err := data.Edges.Merchant.Edges.StoreOrErr(); err == nil {
					return &ProductStore{
						ID:           s.ID,
						BusinessName: s.Name,
					}
				}
				return nil
			}(),
		},
	})

}

func ProductRetailerResponse(data *ent.Product) *fiber.Map {
	return successResponse(&Product{
		ID:          data.ID,
		MajorID:     data.Edges.Major.ID,
		MinorID:     data.Edges.Minor.ID,
		Name:        data.Name,
		Unit:        data.Unit,
		Quantity:    int(data.Quantity),
		Price:       data.Price,
		PromoPrice:  *data.PromoPrice,
		Description: data.Description,
		Image:       data.Image,
		Major:       data.Edges.Major.Category,
		Minor:       data.Edges.Minor.Category,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		Merchant: &MerchantResponse{
			Username:  data.Edges.Merchant.Username,
			LastName:  data.Edges.Merchant.Edges.Retailer.LastName,
			OtherName: data.Edges.Merchant.Edges.Retailer.OtherName,
			Store: func() *ProductStore {
				if s, err := data.Edges.Merchant.Edges.StoreOrErr(); err == nil {
					return &ProductStore{
						ID:           s.ID,
						BusinessName: s.Name,
					}
				}
				return nil
			}(),
		},
	})

}

func ProductsSuccessResponse(data []*ent.Product) *fiber.Map {
	var response []*AllProducts
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Product) {
			defer wg.Done()
			response = append(response, &AllProducts{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  *v.PromoPrice,
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
			})
		}(v)
	}
	wg.Wait()
	return successResponse(response)

}

func ProductsSupplierResponse(data []*ent.Product) *fiber.Map {
	var response []*Product
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Product) {
			defer wg.Done()
			response = append(response, &Product{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  *v.PromoPrice,
				Description: v.Description,
				Image:       v.Image,
				Major:       v.Edges.Major.Category,
				Minor:       v.Edges.Minor.Category,
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
				Merchant: &MerchantResponse{
					Username:  v.Edges.Merchant.Username,
					LastName:  v.Edges.Merchant.Edges.Supplier.LastName,
					OtherName: v.Edges.Merchant.Edges.Supplier.OtherName,
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
			})
		}(v)
	}
	wg.Wait()
	return successResponse(response)

}
func ProductsRetailerResponse(data []*ent.Product) *fiber.Map {
	var response []*Product
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Product) {
			defer wg.Done()
			response = append(response, &Product{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  *v.PromoPrice,
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
						if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err == nil {
							return &ProductStore{
								ID:           s.ID,
								BusinessName: s.Name,
							}
						}
						return nil
					}(),
				},
			})
		}(v)
	}
	wg.Wait()
	return successResponse(response)

}

func ProductsBestSellerResponse(data []*ent.Product) *fiber.Map {
	var response []*Product
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.Product) {
			defer wg.Done()
			response = append(response, &Product{
				ID:          v.ID,
				MajorID:     v.Edges.Major.ID,
				MinorID:     v.Edges.Minor.ID,
				Name:        v.Name,
				Unit:        v.Unit,
				Quantity:    int(v.Quantity),
				Price:       v.Price,
				PromoPrice:  *v.PromoPrice,
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
			})
		}(v)
	}
	wg.Wait()
	return successResponse(response)

}

func ProductsRetailerCategoryMajorResponse(data []*ent.ProductCategoryMajor) *fiber.Map {
	var response []*CategoryMajor
	wg := sync.WaitGroup{}
	for _, cat := range data {
		wg.Add(1)
		go func(cat *ent.ProductCategoryMajor) {
			defer wg.Done()
			response = append(response, &CategoryMajor{
				ID:        cat.ID,
				Category:  cat.Category,
				Slug:      cat.Slug,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
				Product:   formatCategoryProducts(cat.Edges.Products, "retailer"),
			})
		}(cat)
	}
	wg.Wait()
	return successResponse(response)
}

func ProductsRetailerCategoryMinorResponse(data []*ent.ProductCategoryMinor) *fiber.Map {
	var response []*CategoryMajor
	wg := sync.WaitGroup{}
	for _, cat := range data {
		wg.Add(1)
		go func(cat *ent.ProductCategoryMinor) {
			defer wg.Done()
			response = append(response, &CategoryMajor{
				ID:        cat.ID,
				Category:  cat.Category,
				Slug:      cat.Slug,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
				Product:   formatCategoryProducts(cat.Edges.Products, "retailer"),
			})
		}(cat)
	}
	wg.Wait()
	return successResponse(response)
}

func ProductsSupplierCategoryMajorResponse(data []*ent.ProductCategoryMajor) *fiber.Map {
	var response []*CategoryMajor
	wg := sync.WaitGroup{}
	for _, cat := range data {
		wg.Add(1)
		go func(cat *ent.ProductCategoryMajor) {
			defer wg.Done()
			response = append(response, &CategoryMajor{
				ID:        cat.ID,
				Category:  cat.Category,
				Slug:      cat.Slug,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
				Product:   formatCategoryProducts(cat.Edges.Products, "supplier"),
			})
		}(cat)
	}
	wg.Wait()
	return successResponse(response)
}
func ProductsSupplierCategoryMinorResponse(data []*ent.ProductCategoryMinor) *fiber.Map {
	var response []*CategoryMajor
	wg := sync.WaitGroup{}
	for _, cat := range data {
		wg.Add(1)
		go func(cat *ent.ProductCategoryMinor) {
			defer wg.Done()
			response = append(response, &CategoryMajor{
				ID:        cat.ID,
				Category:  cat.Category,
				Slug:      cat.Slug,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
				Product:   formatCategoryProducts(cat.Edges.Products, "supplier"),
			})
		}(cat)
	}
	wg.Wait()
	return successResponse(response)
}

func formatCategoryProducts(data []*ent.Product, merchant string) []*Product {
	var response []*Product
	wg := sync.WaitGroup{}
	if merchant == "retailer" {
		for _, v := range data {
			wg.Add(1)
			go func(v *ent.Product) {
				defer wg.Done()
				response = append(response, &Product{
					ID:          v.ID,
					MajorID:     v.Edges.Major.ID,
					MinorID:     v.Edges.Minor.ID,
					Name:        v.Name,
					Unit:        v.Unit,
					Quantity:    int(v.Quantity),
					Price:       v.Price,
					PromoPrice:  *v.PromoPrice,
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
							if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err == nil {
								return &ProductStore{
									ID:           s.ID,
									BusinessName: s.Name,
								}
							}
							return nil
						}(),
					},
				})
			}(v)
		}

	} else {
		for _, v := range data {
			wg.Add(1)
			go func(v *ent.Product) {
				defer wg.Done()
				response = append(response, &Product{
					ID:          v.ID,
					MajorID:     v.Edges.Major.ID,
					MinorID:     v.Edges.Minor.ID,
					Name:        v.Name,
					Unit:        v.Unit,
					Quantity:    int(v.Quantity),
					Price:       v.Price,
					PromoPrice:  *v.PromoPrice,
					Description: v.Description,
					Image:       v.Image,
					Major:       v.Edges.Major.Category,
					Minor:       v.Edges.Minor.Category,
					CreatedAt:   v.CreatedAt,
					UpdatedAt:   v.UpdatedAt,
					Merchant: &MerchantResponse{
						Username:  v.Edges.Merchant.Username,
						LastName:  v.Edges.Merchant.Edges.Supplier.LastName,
						OtherName: v.Edges.Merchant.Edges.Supplier.OtherName,
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
				})
			}(v)
		}
	}

	wg.Wait()

	return response
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ProductErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
