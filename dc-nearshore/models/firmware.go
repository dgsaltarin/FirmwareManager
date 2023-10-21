package models

import (
	"time"
)

type Firmware struct {
	ID           string    `json:"id"`
	Name         string    `json:"name" sql:"NOT NULL"`
	DeviceID     string    `json:"device_id" sql:"NOT NULL"`
	Version      string    `json:"version"`
	ReleaseNotes string    `json:"release_notes"`
	ReleaseDate  time.Time `json:"release_date"`
	Url          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
