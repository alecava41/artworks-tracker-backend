package models

type Artwork struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TableName gives table name of model
func (u Artwork) TableName() string {
	return "artworks"
}
