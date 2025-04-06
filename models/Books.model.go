package models

type Books struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}
