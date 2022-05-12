package presenters

import (
	"sync"
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	ProductCatMajor struct {
		ID        int       `json:"id"`
		Category  string    `json:"category"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	ProductCatMajorChildren struct {
		Key  int             `json:"key"`
		Data ProductCatMajor `json:"data"`
	}
	ProductCatMajors struct {
		Key      int                       `json:"key"`
		Data     ProductCatMajor           `json:"data"`
		Children []ProductCatMajorChildren `json:"children"`
	}
)

func ProductCatMajorSuccessResponse(data *ent.ProductCategoryMajor) *fiber.Map {
	return successResponse(ProductCatMajor{
		ID:        data.ID,
		Category:  data.Category,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	})
}
func ProductCatMajorsSuccessResponse(data []*ent.ProductCategoryMajor) *fiber.Map {

	var response []ProductCatMajors
	wg := sync.WaitGroup{}
	for _, v := range data {
		wg.Add(1)
		go func(v *ent.ProductCategoryMajor) {
			defer wg.Done()
			major := ProductCatMajors{
				Key: v.ID,
				Data: ProductCatMajor{
					ID:        v.ID,
					Category:  v.Category,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
			}
			for _, m := range v.Edges.Minors {
				major.Children = append(major.Children, ProductCatMajorChildren{
					Key: m.ID,
					Data: ProductCatMajor{
						ID:        m.ID,
						Category:  m.Category,
						CreatedAt: m.CreatedAt,
						UpdatedAt: m.UpdatedAt,
					}})
			}
			response = append(response, major)
		}(v)
	}
	wg.Wait()

	return successResponse(response)
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ProductCatMajorErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
