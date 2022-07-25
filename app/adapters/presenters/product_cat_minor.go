package presenters

import (
	"sync"
	"time"

	"github.com/SeyramWood/ent"
	"github.com/gofiber/fiber/v2"
)

type (
	ProductCatMinor struct {
		ID        int       `json:"id"`
		MajorID   int       `json:"majorId"`
		Major     string    `json:"major"`
		Image     string    `json:"image"`
		Category  string    `json:"category"`
		Slug      string    `json:"slug"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

func ProductCatMinorSuccessResponse(cat *ent.ProductCategoryMinor) *fiber.Map {
	return successResponse(&ProductCatMinor{
		ID:        cat.ID,
		MajorID:   cat.Edges.Major.ID,
		Major:     cat.Edges.Major.Category,
		Category:  cat.Category,
		Slug:      cat.Slug,
		Image:     cat.Image,
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
	})
}
func ProductCatMinorsSuccessResponse(data []*ent.ProductCategoryMinor) *fiber.Map {
	var response []*ProductCatMinor
	wg := sync.WaitGroup{}
	for _, cat := range data {
		wg.Add(1)
		go func(cat *ent.ProductCategoryMinor) {
			defer wg.Done()
			response = append(response, &ProductCatMinor{
				ID:        cat.ID,
				MajorID:   cat.Edges.Major.ID,
				Major:     cat.Edges.Major.Category,
				Category:  cat.Category,
				Slug:      cat.Slug,
				Image:     cat.Image,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
			})
		}(cat)
	}
	wg.Wait()
	return successResponse(response)
}

// UserErrorResponse is the ErrorResponse that will be passed in the response by Handler
func ProductCatMinorErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
