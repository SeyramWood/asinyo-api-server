package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
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
	return successResponse(
		&ProductCatMinor{
			ID:        cat.ID,
			MajorID:   cat.Edges.Major.ID,
			Major:     cat.Edges.Major.Category,
			Category:  cat.Category,
			Slug:      cat.Slug,
			Image:     cat.Image,
			CreatedAt: cat.CreatedAt,
			UpdatedAt: cat.UpdatedAt,
		},
	)
}
func ProductCatMinorsSuccessResponse(data []*ent.ProductCategoryMinor) *fiber.Map {
	var response []*ProductCatMinor
	for _, cat := range data {
		response = append(
			response, &ProductCatMinor{
				ID:        cat.ID,
				MajorID:   cat.Edges.Major.ID,
				Major:     cat.Edges.Major.Category,
				Category:  cat.Category,
				Slug:      cat.Slug,
				Image:     cat.Image,
				CreatedAt: cat.CreatedAt,
				UpdatedAt: cat.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}

func ProductCatMinorErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
