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
)

func AdminSuccessResponse(data *ent.Admin) *fiber.Map {
	return successResponse(
		AdminResponse{
			ID:         data.ID,
			Username:   data.Username,
			LastName:   data.LastName,
			OtherName:  data.OtherName,
			Status:     string(data.Status),
			LastActive: data.LastActive,
			Roles:      formatAdminRoles(data),
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
		},
	)
}
func AdminsSuccessResponse(data *ResponseWithTotalRecords) *fiber.Map {
	var response []*AdminResponse
	for _, v := range data.Records.([]*ent.Admin) {
		response = append(
			response, &AdminResponse{
				ID:         v.ID,
				Username:   v.Username,
				LastName:   v.LastName,
				OtherName:  v.OtherName,
				Status:     string(v.Status),
				LastActive: v.LastActive,
				Roles:      formatAdminRoles(v),
				CreatedAt:  v.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			},
		)
	}
	return successResponse(
		&AdminsWithTotalRecordsResponse{
			TotalRecords: data.TotalRecords,
			Admins:       response,
		},
	)
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
