package models

import (
	"time"
)

type Firmware struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name" gorm:"not null"`
	DeviceID     string    `json:"device_id" gorm:"not null"`
	Version      string    `json:"version"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  time.Time `json:"release_date"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
