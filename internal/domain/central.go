package domain

import "time"

type Central struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"not null" validade:"required"`
	MAC       string    `json:"mac" gorm:"unique;not null" validate:"required,mac"`
	IP        string    `json:"ip" gorm:"unique;not null" validate:"required,ipv4"`
}
