package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Books struct {
	ID            string       `json:"id"`
	Title         string       `gorm:"column:title" json:"title"`
	Author        string       `gorm:"column:author" json:"author"`
	PublishedYear int          `gorm:"column:published_year" json:"publishedYear"`
	CreatedAt     time.Time    `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt     sql.NullTime `gorm:"column:updated_at" json:"updatedAt"`
}

func (Books) TableName() string {
	return "books.books"
}

func (b Books) MarshalJSON() ([]byte, error) {
	type Alias Books
	var updatedAt string
	if b.UpdatedAt.Valid {
		updatedAt = b.UpdatedAt.Time.Format("2006-01-02")
	} else {
		updatedAt = ""
	}
	return json.Marshal(&struct {
		*Alias
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}{
		Alias:     (*Alias)(&b),
		CreatedAt: b.CreatedAt.Format("2006-01-02"),
		UpdatedAt: updatedAt,
	})
}
