package model

import "time"

// Cuti represents the cuti table
type Cuti struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uint      `gorm:"not null" json:"user_id"`
	SubmissionDate time.Time `gorm:"type:date;not null" json:"submission_date"`
	StartDate      time.Time `gorm:"type:date;not null" json:"start_date"`
	EndDate        time.Time `gorm:"type:date;not null" json:"end_date"`
	Information    string    `gorm:"type:text" json:"information"`
	StatusCutiID   uint      `gorm:"not null" json:"status_cuti_id"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relations
	StatusCuti StatusCuti `gorm:"foreignKey:StatusCutiID" json:"status_cuti"`
}

func (Cuti) TableName() string {
	return "cuti"
}
