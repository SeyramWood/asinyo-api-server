package presenters

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	ProductCatMajor struct {
		ID        int       `json:"id"`
		Category  string    `json:"category"`
		Slug      string    `json:"slug"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	ProductCatMinors struct {
		ID        int       `json:"id"`
		Category  string    `json:"category"`
		Slug      string    `json:"slug"`
		Image     string    `json:"image"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	ProductCatMajorChildren struct {
		Key  string           `json:"key"`
		Data ProductCatMinors `json:"data"`
	}
	ProductCatMajors struct {
		Key          string                     `json:"key"`
		Data         ProductCatMajor            `json:"data"`
		ProductCount int                        `json:"productCount,omitempty"`
		Children     []*ProductCatMajorChildren `json:"children"`
	}
)

func ProductCatMajorSuccessResponse(data *ent.ProductCategoryMajor) *fiber.Map {
	return successResponse(
		&ProductCatMajor{
			ID:        data.ID,
			Category:  data.Category,
			Slug:      data.Slug,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		},
	)
}
func ProductCatMajorsSuccessResponse(data []*ent.ProductCategoryMajor) *fiber.Map {

	var response []*ProductCatMajors

	for _, v := range data {
		major := &ProductCatMajors{
			Key: strconv.Itoa(v.ID),
			Data: ProductCatMajor{
				ID:        v.ID,
				Category:  v.Category,
				Slug:      v.Slug,
				CreatedAt: v.CreatedAt,
				UpdatedAt: v.UpdatedAt,
			},
			ProductCount: len(v.Edges.Products),
			Children: func() []*ProductCatMajorChildren {
				var children []*ProductCatMajorChildren
				for _, m := range v.Edges.Minors {
					children = append(
						children, &ProductCatMajorChildren{
							Key: fmt.Sprintf("%s-%s", strconv.Itoa(v.ID), strconv.Itoa(m.ID)),
							Data: ProductCatMinors{
								ID:        m.ID,
								Category:  m.Category,
								Slug:      m.Slug,
								Image:     m.Image,
								CreatedAt: m.CreatedAt,
								UpdatedAt: m.UpdatedAt,
							},
						},
					)
				}
				return children
			}(),
		}

		response = append(response, major)

	}

	return successResponse(response)
}

func ProductCatMajorErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
