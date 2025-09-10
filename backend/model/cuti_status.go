package model

import (
	"time"
)

// StatusCuti represents the status_cuti table
type StatusCuti struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	StatusName string    `gorm:"type:varchar(50);not null" json:"status_name"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (StatusCuti) TableName() string {
	return "status_cuti"
}
