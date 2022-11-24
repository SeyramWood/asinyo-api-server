package services

type (
	TookanTaskResponse struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
		Data    struct {
			TrackingLink    string `json:"tracking_link"`
			JobID           int    `json:"job_id"`
			JobHash         string `json:"job_hash"`
			JobToken        string `json:"job_token"`
			CustomerName    string `json:"customer_name"`
			CustomerAddress string `json:"customer_address"`
			GeofenceDetails []struct {
				RegionID   int    `json:"region_id"`
				RegionName string `json:"region_name"`
			} `json:"geofence_details"`
		} `json:"data"`
	}
	TookanMultiTaskResponseData struct {
		JobID                     int           `json:"job_id,omitempty"`
		JobHash                   string        `json:"job_hash,omitempty"`
		JobToken                  string        `json:"job_token,omitempty"`
		Status                    bool          `json:"status,omitempty"`
		AutoAssignmentData        int           `json:"auto_assignment_data,omitempty"`
		OrderID                   string        `json:"order_id,omitempty"`
		ResultTrackingLink        string        `json:"result_tracking_link,omitempty"`
		NewAgentTaskInsertionData []interface{} `json:"new_agent_task_insertion_data,omitempty"`
	}
	TookanGeofenceDetails struct {
		RegionID   int    `json:"region_id"`
		RegionName string `json:"region_name"`
	}
	TookanPickupDeliveryResponse struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
		Data    struct {
			Pickups         []TookanMultiTaskResponseData `json:"pickups"`
			Deliveries      []TookanMultiTaskResponseData `json:"deliveries"`
			GeofenceDetails []TookanGeofenceDetails       `json:"geofence_details"`
		} `json:"data"`
	}
	TookanMetadata struct {
		Label string `json:"label"`
		Data  any    `json:"data"`
	}
	TookanDeliveryTask struct {
		APIKey              string           `json:"api_key"`
		OrderID             string           `json:"order_id"`
		JobDescription      string           `json:"job_description"`
		CustomerEmail       string           `json:"customer_email"`
		CustomerUsername    string           `json:"customer_username"`
		CustomerPhone       string           `json:"customer_phone"`
		CustomerAddress     string           `json:"customer_address"`
		Latitude            string           `json:"latitude"`
		Longitude           string           `json:"longitude"`
		JobDeliveryDatetime string           `json:"job_delivery_datetime"`
		CustomFieldTemplate string           `json:"custom_field_template"`
		MetaData            []TookanMetadata `json:"meta_data"`
		TeamID              string           `json:"team_id"`
		AutoAssignment      string           `json:"auto_assignment"`
		HasPickup           string           `json:"has_pickup"`
		HasDelivery         string           `json:"has_delivery"`
		LayoutType          string           `json:"layout_type"`
		TrackingLink        int              `json:"tracking_link"`
		Timezone            string           `json:"timezone"`
		FleetID             string           `json:"fleet_id"`
		RefImages           []string         `json:"ref_images"`
		Notify              int              `json:"notify"`
		Tags                string           `json:"tags"`
		Geofence            int              `json:"geofence"`
	}

	TookanPickupDelivery struct {
		JobID          string           `json:"job_id,omitempty"`
		Address        string           `json:"address"`
		Latitude       float64          `json:"latitude"`
		Longitude      float64          `json:"longitude"`
		Time           string           `json:"time"`
		Phone          string           `json:"phone"`
		JobDescription string           `json:"job_description"`
		TemplateName   string           `json:"template_name"`
		TemplateData   []TookanMetadata `json:"template_data"`
		RefImages      []string         `json:"ref_images"`
		Name           string           `json:"name"`
		Email          string           `json:"email"`
		OrderID        string           `json:"order_id"`
	}

	TookanPickupDeliveryTask struct {
		APIKey         string                 `json:"api_key"`
		FleetID        int                    `json:"fleet_id"`
		Timezone       int                    `json:"timezone"`
		HasPickup      int                    `json:"has_pickup"`
		HasDelivery    int                    `json:"has_delivery"`
		LayoutType     int                    `json:"layout_type"`
		Geofence       int                    `json:"geofence"`
		TeamID         string                 `json:"team_id"`
		AutoAssignment int                    `json:"auto_assignment"`
		Tags           string                 `json:"tags"`
		Pickups        []TookanPickupDelivery `json:"pickups"`
		Deliveries     []TookanPickupDelivery `json:"deliveries"`
	}
	TookanPickupDeliveryUpdateTask struct {
		APIKey     string                 `json:"api_key"`
		Pickups    []TookanPickupDelivery `json:"pickups"`
		Deliveries []TookanPickupDelivery `json:"deliveries"`
	}

	FareEstimateResponseFormula struct {
		DisplayName      string  `json:"display_name,omitempty"`
		Key              string  `json:"key,omitempty"`
		Type             float64 `json:"type,omitempty"`
		Surge            string  `json:"surge,omitempty"`
		MultiplyingValue string  `json:"multiplying_value,omitempty"`
		Expression       string  `json:"expression,omitempty"`
		Sum              float64 `json:"sum,omitempty"`
	}
	FareEstimateResponseData struct {
		Distance      float64                        `json:"distance,omitempty"`
		Time          float64                        `json:"time,omitempty"`
		Formula       []*FareEstimateResponseFormula `json:"formula,omitempty"`
		EstimatedFare float64                        `json:"estimated_fare,omitempty"`
	}
	FareEstimateResponse struct {
		Message string                   `json:"message,omitempty"`
		Status  int                      `json:"status,omitempty"`
		Data    FareEstimateResponseData `json:"data,omitempty"`
	}
	FareEstimateRequest struct {
		TemplateName      string `json:"template_name,omitempty"`
		PickupLongitude   string `json:"pickup_longitude,omitempty"`
		PickupLatitude    string `json:"pickup_latitude,omitempty"`
		APIKey            string `json:"api_key,omitempty"`
		DeliveryLatitude  string `json:"delivery_latitude,omitempty"`
		DeliveryLongitude string `json:"delivery_longitude,omitempty"`
		FormulaType       int    `json:"formula_type,omitempty"`
		MapKeys           struct {
			MapPlanType  int    `json:"map_plan_type,omitempty"`
			GoogleAPIKey string `json:"google_api_key,omitempty"`
		} `json:"map_keys,omitempty"`
	}
)
