package dto

import "time"

type PresenceRequest struct {
	LatitudeLongitude *string `json:"latitudeLongitude,omitempty"` // gabungkan lat & long
}

type CreatePresenceResponse struct {
	Name                 string    `json:"name"`
	PresenceDate         time.Time `json:"presence_date"`
	TimeIn               string    `json:"time_in"`
	TimeOut              *string   `json:"time_out,omitempty"`
	IsLate               bool      `json:"is_late"`
	LatitudeLongitudeIn  string    `json:"latitude_longitude_in"`
	LatitudeLongitudeOut *string   `json:"latitude_longitude_out,omitempty"`
	Status               string    `json:"status"`
}
