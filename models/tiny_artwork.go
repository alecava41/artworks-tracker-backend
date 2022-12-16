package models

type TinyArtwork struct {
	ID      uint   `json:"id"`
	PosX    uint   `json:"posX"`
	PosY    uint   `json:"posY"`
	Starred bool   `json:"starred"`
	Type    string `json:"type"`
}

// TableName gives table name of model
func (u TinyArtwork) TableName() string {
	return "artworks"
}
