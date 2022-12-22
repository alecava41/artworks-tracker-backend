package models

type TinyArtwork struct {
	ID      uint   `json:"id"`
	PosX    uint   `json:"posX" gorm:"column:posX"`
	PosY    uint   `json:"posY" gorm:"column:posY"`
	Starred bool   `json:"starred"`
	Type    string `json:"type"`
	RoomId  uint   `json:"-" gorm:"column:room"`
}

// TableName gives table name of model
func (u TinyArtwork) TableName() string {
	return "artworks"
}
