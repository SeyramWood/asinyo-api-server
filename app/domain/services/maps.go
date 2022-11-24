package services

type (
	GeocodingData struct {
		Street   string `json:"street,omitempty"`
		Number   int    `json:"number,omitempty"`
		City     string `json:"city,omitempty"`
		District string `json:"district,omitempty"`
		State    string `json:"state,omitempty"`
		Country  string `json:"country,omitempty"`
	}
	Coordinate struct {
		Latitude  float64 `json:"latitude,omitempty"`
		Longitude float64 `json:"longitude,omitempty"`
	}
)
