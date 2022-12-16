package models

import (
	"time"
)

type Room struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	Shape     *string        `json:"shape"`
	Artworks  []*TinyArtwork `json:"artworks"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// TableName gives table name of model
func (u Room) TableName() string {
	return "rooms"
}
