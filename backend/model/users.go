package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	Role      string    `gorm:"type:varchar(50);not null" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// One-to-One relationship
	// Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"profile"`

	// // One-to-Many relationship
	// Presences []Presence `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"presences"`
}


func (User) TableName() string{
	return  "users"
}