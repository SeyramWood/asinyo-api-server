package presenters

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/SeyramWood/ent"
)

type (
	PermissionsResponse struct {
		ID         int       `json:"id"`
		Permission string    `json:"permission"`
		CreatedAt  time.Time `json:"created_at,omitempty"`
		UpdatedAt  time.Time `json:"updated_at,omitempty"`
	}
	RoleResponse struct {
		ID          int                    `json:"id"`
		Role        string                 `json:"role"`
		Permissions []*PermissionsResponse `json:"permissions"`
		CreatedAt   time.Time              `json:"created_at"`
		UpdatedAt   time.Time              `json:"updated_at"`
	}
	RolesWithTotalRecordsResponse struct {
		TotalRecords int             `json:"totalRecords,omitempty"`
		Roles        []*RoleResponse `json:"roles"`
	}
)

func RoleSuccessResponse(data *ent.Role) *fiber.Map {
	return successResponse(
		&RoleResponse{
			ID:          data.ID,
			Role:        data.Role,
			Permissions: formatRolePermissions(data),
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		},
	)
}

func RoleSuccessResponses(data *ResponseWithTotalRecords) *fiber.Map {
	var response []*RoleResponse
	for _, v := range data.Records.([]*ent.Role) {
		response = append(
			response, &RoleResponse{
				ID:          v.ID,
				Role:        v.Role,
				Permissions: formatRolePermissions(v),
				CreatedAt:   v.CreatedAt,
				UpdatedAt:   v.UpdatedAt,
			},
		)
	}
	return successResponse(
		&RolesWithTotalRecordsResponse{
			TotalRecords: data.TotalRecords,
			Roles:        response,
		},
	)
}
func PermissionSuccessResponses(data []*ent.Permission) *fiber.Map {
	var response []*PermissionsResponse
	for _, v := range data {
		response = append(
			response, &PermissionsResponse{
				ID:         v.ID,
				Permission: v.Permission,
				CreatedAt:  v.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			},
		)
	}
	return successResponse(response)
}

func RolePermissionErrorResponse(err error) *fiber.Map {
	return errorResponse(err)
}

func formatRolePermissions(role *ent.Role) []*PermissionsResponse {
	result, err := role.Edges.PermissionsOrErr()
	if err != nil {
		return nil
	}
	var response []*PermissionsResponse
	for _, p := range result {
		response = append(
			response, &PermissionsResponse{
				ID:         p.ID,
				Permission: p.Permission,
				CreatedAt:  p.CreatedAt,
				UpdatedAt:  p.UpdatedAt,
			},
		)
	}

	return response
}
