package models

import (
	"time"
)

type Device struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" sql:"NOT NULL"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
