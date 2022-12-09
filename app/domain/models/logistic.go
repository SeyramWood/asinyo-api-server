package models

import "github.com/SeyramWood/app/domain/services"

type (
	TookanMultiTaskResponse struct {
		Pickups    []services.TookanMultiTaskResponseData `json:"pickups"`
		Deliveries []services.TookanMultiTaskResponseData `json:"deliveries"`
		Geofence   []services.TookanGeofenceDetails       `json:"geofence"`
	}
	TookanPickupAndDeliveryTaskResponse struct {
		PickupTrackingLink  string                           `json:"pickup_tracking_link"`
		DeliveryTracingLink string                           `json:"delivery_tracing_link"`
		JobID               int                              `json:"job_id"`
		JobToken            string                           `json:"job_token"`
		PickupHash          string                           `json:"pickup_hash"`
		DeliveryHash        string                           `json:"delivery_hash"`
		CustomerName        string                           `json:"customer_name"`
		CustomerAddress     string                           `json:"customer_address"`
		GeofenceDetails     []services.TookanGeofenceDetails `json:"geofence"`
	}
)
