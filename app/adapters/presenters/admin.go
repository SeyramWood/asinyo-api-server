package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	AdminResponse struct {
		ID         int          `json:"id"`
		Username   string       `json:"username"`
		LastName   string       `json:"lastName"`
		OtherName  string       `json:"otherName"`
		Phone      string       `json:"phone"`
		OtherPhone string       `json:"otherPhone"`
		Status     string       `json:"status"`
		LastActive string       `json:"lastActive"`
		Roles      []*AdminRole `json:"roles,omitempty"`
		CreatedAt  time.Time    `json:"created_at"`
		UpdatedAt  time.Time    `json:"updated_at"`
	}
	AdminRole struct {
		ID   int    `json:"id"`
		Role string `json:"role"`
	}
	AdminsWithTotalRecordsResponse struct {
		TotalRecords int              `json:"totalRecords,omitempty"`
		Admins       []*AdminResponse `json:"admins"`
	}
	MyClientsWithPurchaseOrder struct {
		ID       int    `json:"id"`
		Business string `json:"business"`
	}
)

func AdminSuccessResponse(data *ent.Admin) *fiber.Map {
	return successResponse(
		AdminResponse{
			ID:         data.ID,
			Username:   data.Username,
			LastName:   data.LastName,
			OtherName:  data.OtherName,
			Phone:      data.Phone,
			OtherPhone: data.OtherPhone,
			Status:     string(data.Status),
			LastActive: data.LastActive,
			Roles:      formatAdminRoles(data),
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
		},
	)
}

func DashboardCountSuccessResponse(data *DashboardRecordCount) *fiber.Map {
	return successResponse(data)
}

func MyClientSuccessResponse(data *ent.Customer) *fiber.Map {
	if c, err := data.Edges.IndividualOrErr(); err == nil {
		return successResponse(
			&Customer{
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
		return successResponse(
			&Customer{
				ID:           data.ID,
				Logo:         c.Logo,
				Username:     data.Username,
				Company:      c.Name,
				Phone:        c.Phone,
				OtherPhone:   c.OtherPhone,
				CustomerType: "business",
				Contact:      c.Contact,
				CreatedAt:    data.CreatedAt,
				UpdatedAt:    data.UpdatedAt,
			},
		)
	}
	return successResponse(nil)
}

func MyClientsSuccessResponse(res *PaginationResponse) *fiber.Map {
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
					CreatedAt:    data.CreatedAt,
					UpdatedAt:    data.UpdatedAt,
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
func MyClientsPurchaseRequestSuccessResponse(res []*ent.Customer) *fiber.Map {
	var response []*Customer
	for _, data := range res {
		if c, err := data.Edges.IndividualOrErr(); err == nil {
			response = append(
				response, &Customer{
					ID:              data.ID,
					LastName:        c.LastName,
					OtherName:       c.OtherName,
					Phone:           c.Phone,
					OtherPhone:      c.OtherPhone,
					PurchaseRequest: formatManagerPR(data),
					CustomerType:    "individual",
					CreatedAt:       data.CreatedAt,
					UpdatedAt:       data.UpdatedAt,
				},
			)
		}
		if c, err := data.Edges.BusinessOrErr(); err == nil {
			response = append(
				response, &Customer{
					ID:              data.ID,
					Company:         c.Name,
					Phone:           c.Phone,
					OtherPhone:      c.OtherPhone,
					PurchaseRequest: formatManagerPR(data),
					CustomerType:    "business",
					CreatedAt:       data.CreatedAt,
					UpdatedAt:       data.UpdatedAt,
				},
			)
		}
	}
	return successResponse(&response)
}
func PaginateOrdersSuccessResponse(res *PaginationResponse) *fiber.Map {
	var response []*Order
	for _, data := range res.Data.([]*ent.Order) {
		response = append(
			response, &Order{
				ID:            data.ID,
				OrderNumber:   data.OrderNumber,
				Amount:        data.Amount,
				Currency:      data.Currency,
				Channel:       data.Channel,
				PaymentMethod: string(data.PaymentMethod),
				PaidAt:        data.PaidAt,
				Status:        string(data.Status),
				Approval:      string(data.CustomerApproval),
				CreatedAt:     data.CreatedAt,
				UpdatedAt:     data.UpdatedAt,
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

func AdminProductsResponse(res *PaginationResponse) *fiber.Map {
	var response []*ProductWithStore
	for _, v := range res.Data.([]*ent.Product) {
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
				PriceModel: func() *PriceModel {
					if m, err := v.Edges.PriceModelOrErr(); err == nil {
						return &PriceModel{
							ID:            m.ID,
							Name:          m.Name,
							Initials:      m.Initials,
							Formula:       m.Formula,
							AsinyoFormula: m.AsinyoFormula,
						}
					}
					return nil
				}(),
				Store: func() *ProductStore {
					if s, err := v.Edges.Merchant.Edges.StoreOrErr(); err == nil {
						return &ProductStore{
							ID:           s.ID,
							BusinessName: s.Name,
							Coordinate:   s.Coordinate,
						}
					}
					return nil
				}(),
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

func AdminsSuccessResponse(data *PaginationResponse) *fiber.Map {
	var response []*AdminResponse
	for _, v := range data.Data.([]*ent.Admin) {
		response = append(
			response, &AdminResponse{
				ID:         v.ID,
				Username:   v.Username,
				LastName:   v.LastName,
				OtherName:  v.OtherName,
				Phone:      v.Phone,
				OtherPhone: v.OtherPhone,
				Status:     string(v.Status),
				LastActive: v.LastActive,
				Roles:      formatAdminRoles(v),
				CreatedAt:  v.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			},
		)
	}
	return successResponse(
		&PaginationResponse{
			Count: data.Count,
			Data:  response,
		},
	)
}
func AccountManagersSuccessResponse(data []*ent.Admin) *fiber.Map {
	var response []*AdminResponse
	for _, v := range data {
		response = append(
			response, &AdminResponse{
				ID:         v.ID,
				Username:   v.Username,
				LastName:   v.LastName,
				OtherName:  v.OtherName,
				Status:     string(v.Status),
				LastActive: v.LastActive,
				CreatedAt:  v.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}

func AdminErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}

func formatAdminRoles(data *ent.Admin) []*AdminRole {
	roles, err := data.Edges.RolesOrErr()
	if err != nil {
		return nil
	}
	var response []*AdminRole
	for _, r := range roles {
		response = append(
			response, &AdminRole{
				ID:   r.ID,
				Role: r.Role,
			},
		)
	}
	return response
}
func formatManagerPR(c *ent.Customer) []*PurchaseRequest {
	pos, err := c.Edges.PurchaseRequestOrErr()
	if err != nil {
		return nil
	}
	var response []*PurchaseRequest
	for _, data := range pos {
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
	return response
}
