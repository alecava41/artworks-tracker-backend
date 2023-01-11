package models

type Room struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Perimeter   *string        `json:"perimeter"`
	Walls       *string        `json:"walls"`
	StarredPath *string        `json:"starredPath" gorm:"column:starredPath"`
	Artworks    []*ArtworkInfo `json:"artworks" gorm:"foreignKey:RoomId"`
}

// TableName gives table name of model
func (u Room) TableName() string {
	return "rooms"
}
