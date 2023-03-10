package models

import "github.com/google/uuid"

type Artwork struct {
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Beacon      uuid.UUID `json:"beacon"`
}

// TableName gives table name of model
func (u Artwork) TableName() string {
	return "artworks"
}
