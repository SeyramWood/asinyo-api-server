package models

import "github.com/SeyramWood/app/domain/services"

type (
	TookanMultiTaskResponse struct {
		Pickups    []services.TookanMultiTaskResponseData `json:"pickups"`
		Deliveries []services.TookanMultiTaskResponseData `json:"deliveries"`
		Geofence   []services.TookanGeofenceDetails       `json:"geofence"`
	}
)
