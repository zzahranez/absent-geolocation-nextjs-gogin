package model

type PresenceStatus struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	StatusName string `gorm:"column:status_name;type:varchar(20);not null" json:"status_name"`
}

func (PresenceStatus) TableName() string {
	return "presence_status"
}
