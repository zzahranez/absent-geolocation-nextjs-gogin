package dto

import "time"

type MontlyPresenceSummary struct {
	TotalHadir     int `json:"total_hadir"`
	TotalIzin      int `json:"total_izin"`
	TotalTerlambat int `json:"total_terlambat"`
	TotalSakit     int `json:"total_sakit"`
}

type PresenceToDay struct {
	Morning  string `json:"morning_absen"`
	AfterNon string `json:"after_noon_absen"`
}

type PresencesAll struct {
	PresenceDate           time.Time `gorm:"column:presence_date" json:"presence_date"`
	TimeIn                 *string   `gorm:"column:time_in" json:"time_in"`
	TimeOut                *string   `gorm:"column:time_out" json:"time_out"`
	LatitudeLongTitudeIn   string    `json:"latitudelongtitude_in"`
	LatitudeLongTitudeeOut string    `json:"latitudelongtitude_out"`
}

type PresenceResponse struct {
	Name                  string                `json:"user_name"`
	MontlyPresenceSummary MontlyPresenceSummary `json:"montly_presence_summary"`
	PresenceToday         PresenceToDay         `json:"presence_today"`
	Presence              []PresencesAll        `json:"presence"`
}
