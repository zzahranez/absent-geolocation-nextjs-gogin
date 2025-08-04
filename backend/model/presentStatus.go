package model

type PresentStatus struct {
	Id         uint   `gorm:"primaryKey" json:"id"`
	StatusName string `gorm:"column:status_name;type:varchar(20);not null" json:"status_name"`
}

func (PresentStatus) TableName() string {
	return "present_status"
}
