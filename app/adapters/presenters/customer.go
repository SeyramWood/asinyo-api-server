package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent"
)

type (
	CustomerAccountManager struct {
		ID         int    `json:"id"`
		Username   string `json:"username,omitempty"`
		LastName   string `json:"lastName,omitempty"`
		OtherName  string `json:"otherName,omitempty"`
		Phone      string `json:"phone,omitempty"`
		OtherPhone string `json:"otherPhone,omitempty"`
	}
	IndividualCustomer struct {
		ID           int       `json:"id"`
		Username     string    `json:"username,omitempty"`
		LastName     string    `json:"lastName,omitempty"`
		OtherName    string    `json:"otherName,omitempty"`
		Phone        string    `json:"phone,omitempty"`
		OtherPhone   string    `json:"otherPhone,omitempty"`
		CustomerType string    `json:"customerType,omitempty"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at,omitempty"`
	}
	BusinessCustomer struct {
		ID           int                             `json:"id"`
		Username     string                          `json:"username,omitempty"`
		Logo         string                          `json:"logo,omitempty"`
		Name         string                          `json:"name,omitempty"`
		Phone        string                          `json:"phone,omitempty"`
		OtherPhone   string                          `json:"OtherPhone,omitempty"`
		Contact      *models.BusinessCustomerContact `json:"contact,omitempty"`
		Manager      *CustomerAccountManager         `json:"manager,omitempty"`
		CustomerType string                          `json:"customerType,omitempty"`
		CreatedAt    time.Time                       `json:"created_at"`
		UpdatedAt    time.Time                       `json:"updated_at,omitempty"`
	}
	Customer struct {
		ID              int                             `json:"id"`
		Username        string                          `json:"username,omitempty"`
		LastName        string                          `json:"lastName,omitempty"`
		OtherName       string                          `json:"otherName,omitempty"`
		Logo            string                          `json:"logo,omitempty"`
		Company         string                          `json:"company,omitempty"`
		Phone           string                          `json:"phone,omitempty"`
		OtherPhone      string                          `json:"OtherPhone,omitempty"`
		Contact         *models.BusinessCustomerContact `json:"contact,omitempty"`
		Manager         *CustomerAccountManager         `json:"manager,omitempty"`
		PurchaseRequest []*PurchaseRequest              `json:"purchaseRequest,omitempty"`
		CustomerType    string                          `json:"customerType,omitempty"`
		CreatedAt       time.Time                       `json:"created_at"`
		UpdatedAt       time.Time                       `json:"updated_at,omitempty"`
	}
	PurchaseRequest struct {
		ID          int       `json:"id"`
		Name        string    `json:"name,omitempty"`
		Description string    `json:"description,omitempty"`
		File        string    `json:"file,omitempty"`
		Signed      string    `json:"signed,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at,omitempty"`
	}
)

func CustomerSuccessResponse(data *ent.Customer) *fiber.Map {
	if c, err := data.Edges.IndividualOrErr(); err == nil {
		return successResponse(
			&IndividualCustomer{
				ID:           data.ID,
				LastName:     c.LastName,
				OtherName:    c.OtherName,
				Phone:        c.Phone,
				OtherPhone:   c.OtherPhone,
				CustomerType: "individual",
				CreatedAt:    data.CreatedAt,
				UpdatedAt:    data.UpdatedAt,
			},
		)
	}
	if c, err := data.Edges.BusinessOrErr(); err == nil {
		return successResponse(
			&BusinessCustomer{
				ID:           data.ID,
				Logo:         c.Logo,
				Name:         c.Name,
				Phone:        c.Phone,
				OtherPhone:   c.OtherPhone,
				Contact:      c.Contact,
				CustomerType: "business",
				Manager: func() *CustomerAccountManager {
					if manager, err := data.Edges.AdminOrErr(); err == nil {
						return &CustomerAccountManager{
							ID:         manager.ID,
							Username:   manager.Username,
							LastName:   manager.LastName,
							OtherName:  manager.OtherName,
							Phone:      manager.Phone,
							OtherPhone: manager.OtherPhone,
						}
					}
					return nil
				}(),
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
		)
	}

	return successResponse(nil)
}

func CustomersSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*Customer
	for _, data := range res.Data.([]*ent.Customer) {
		if c, err := data.Edges.IndividualOrErr(); err == nil {
			response = append(
				response, &Customer{
					ID:           data.ID,
					Username:     data.Username,
					LastName:     c.LastName,
					OtherName:    c.OtherName,
					Phone:        c.Phone,
					OtherPhone:   c.OtherPhone,
					CustomerType: "individual",
					CreatedAt:    data.CreatedAt,
					UpdatedAt:    data.UpdatedAt,
				},
			)
		}
		if c, err := data.Edges.BusinessOrErr(); err == nil {
			response = append(
				response, &Customer{
					ID:           data.ID,
					Logo:         c.Logo,
					Username:     data.Username,
					Company:      c.Name,
					Phone:        c.Phone,
					OtherPhone:   c.OtherPhone,
					CustomerType: "business",
					Contact:      c.Contact,
					Manager: func() *CustomerAccountManager {
						if manager, err := data.Edges.AdminOrErr(); err == nil {
							return &CustomerAccountManager{
								ID:         manager.ID,
								Username:   manager.Username,
								LastName:   manager.LastName,
								OtherName:  manager.OtherName,
								Phone:      manager.Phone,
								OtherPhone: manager.OtherPhone,
							}
						}
						return nil
					}(),
					CreatedAt: data.CreatedAt,
					UpdatedAt: data.UpdatedAt,
				},
			)
		}
	}
	return successResponse(
		&PaginationResponse{
			Count: res.Count,
			Data:  response,
		},
	)
}

func PurchaseOrdersSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*PurchaseRequest
	for _, data := range res.Data.([]*ent.PurchaseRequest) {
		response = append(
			response, &PurchaseRequest{
				ID:          data.ID,
				Name:        data.Name,
				Description: data.Description,
				File:        data.File,
				Signed:      data.Signed,
				CreatedAt:   data.CreatedAt,
				UpdatedAt:   data.UpdatedAt,
			},
		)
	}
	return successResponse(
		&PaginationResponse{
			Count: res.Count,
			Data:  response,
		},
	)
}

func PurchaseRequestSuccessResponse(data *ent.PurchaseRequest) *fiber.Map {
	return successResponse(
		&PurchaseRequest{
			ID:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			File:        data.File,
			Signed:      data.Signed,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		},
	)
}

func CustomerErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}
