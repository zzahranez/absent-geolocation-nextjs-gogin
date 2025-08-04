package dto

import "time"

type MontlyPresenceSummary struct {
}

type PresencesAll struct {
	PresenceDate           time.Time `gorm:"column:presence_date" json:"presence_date"`
	TimeIn                 *string   `gorm:"column:time_in" json:"time_in"`
	TimeOut                *string   `gorm:"column:time_out" json:"time_out"`
	LatitudeLongTitudeIn   string    `json:"latitudelongtitude_in"`
	LatitudeLongTitudeeOut string    `json:"latitudelongtitude_out"`
}

type PresenceResponse struct {
	Name     string         `json:"user_name"`
	Presence []PresencesAll `json:"presence"`
}
