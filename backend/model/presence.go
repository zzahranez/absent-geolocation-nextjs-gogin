package model

import "time"

type Presences struct {
	ID                   uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	UserID               uint      `gorm:"column:user_id;not null" json:"user_id"`
	PresenceDate         time.Time `gorm:"column:presence_date" json:"presence_date"`
	TimeIn               *string   `gorm:"column:time_in" json:"time_in"`
	TimeOut              *string   `gorm:"column:time_out" json:"time_out"`
	StatusID             *uint     `gorm:"column:status_id" json:"status_id"`
	LatitudeLongitudeIn  *string   `gorm:"column:latitude_longtitude_in;type:varchar(100)" json:"latitude_longtitude_in"`
	LatitudeLongitudeOut *string   `gorm:"column:latitude_longtitude_out;type:varchar(100)" json:"latitude_longtitude_out"`
	CreatedAt            time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	User          User            `gorm:"foreignKey:UserID;references:ID" json:"user"` // Relasi ke User
	PresentStatus *PresenceStatus `gorm:"foreignKey:StatusID;references:ID" json:"status"`
}

func (Presences) TableName() string {
	return "presences"
}
