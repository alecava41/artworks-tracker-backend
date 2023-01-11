package models

import "github.com/google/uuid"

type ArtworkInfo struct {
	ID        uint      `json:"id"`
	PosX      uint      `json:"posX" gorm:"column:posX"`
	PosY      uint      `json:"posY" gorm:"column:posY"`
	Starred   bool      `json:"starred"`
	Side      string    `json:"side"`
	Type      string    `json:"type"`
	Direction string    `json:"direction"`
	Beacon    uuid.UUID `json:"beacon"`
	RoomId    uint      `json:"-" gorm:"column:room"`
}

// TableName gives table name of model
func (u ArtworkInfo) TableName() string {
	return "artworks"
}
