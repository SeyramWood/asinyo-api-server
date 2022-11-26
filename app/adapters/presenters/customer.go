package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type (
	IndividualCustomer struct {
		ID         int       `json:"id"`
		LastName   string    `json:"lastName"`
		OtherName  string    `json:"otherName"`
		Phone      string    `json:"phone"`
		OtherPhone string    `json:"otherPhone"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
	BusinessCustomer struct {
		ID         int                             `json:"id"`
		Logo       string                          `json:"logo"`
		Name       string                          `json:"name"`
		Phone      string                          `json:"phone"`
		OtherPhone string                          `json:"OtherPhone"`
		Contact    *models.BusinessCustomerContact `json:"contact"`
		CreatedAt  time.Time                       `json:"created_at"`
		UpdatedAt  time.Time                       `json:"updated_at"`
	}
)

func CustomerSuccessResponse(data *ent.Customer) *fiber.Map {
	if c, err := data.Edges.IndividualOrErr(); err == nil {
		return successResponse(
			&IndividualCustomer{
				ID:         data.ID,
				LastName:   c.LastName,
				OtherName:  c.OtherName,
				Phone:      c.Phone,
				OtherPhone: c.OtherPhone,
				CreatedAt:  data.CreatedAt,
				UpdatedAt:  data.UpdatedAt,
			},
		)
	}
	if c, err := data.Edges.BusinessOrErr(); err == nil {
		return successResponse(
			&BusinessCustomer{
				ID:         data.ID,
				Logo:       c.Logo,
				Name:       c.Name,
				Phone:      c.Phone,
				OtherPhone: c.OtherPhone,
				Contact:    c.Contact,
				CreatedAt:  data.CreatedAt,
				UpdatedAt:  data.UpdatedAt,
			},
		)
	}

	return successResponse(nil)
}
func CustomersSuccessResponse(data []*ent.Customer) *fiber.Map {
	// var response []*Customer
	// for _, v := range data {
	// 	response = append(
	// 		response, &Customer{
	// 			ID:           v.ID,
	// 			Username:     v.Username,
	// 			CustomerType: v.Type,
	// 			CreatedAt:    v.CreatedAt,
	// 			UpdatedAt:    v.UpdatedAt,
	// 		},
	// 	)
	// }
	return successResponse(nil)
}

func CustomerErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
