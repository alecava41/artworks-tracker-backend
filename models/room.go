package models

type Room struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Shape    *string        `json:"shape"`
	Artworks []*TinyArtwork `json:"artworks" gorm:"foreignKey:RoomId"`
}

// TableName gives table name of model
func (u Room) TableName() string {
	return "rooms"
}
