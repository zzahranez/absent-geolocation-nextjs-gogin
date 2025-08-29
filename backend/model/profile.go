package model

import (
	"time"
)

type Profile struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int       `gorm:"unique;not null" json:"user_id"`
	Bio       string    `gorm:"type:text" json:"bio,omitempty"`
	Avatar    string    `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone,omitempty"`
	Address   string    `gorm:"type:text" json:"address,omitempty"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
}
