package models

import "time"

type Blog struct {
	ID        string   `json:"id" gorm:"primaryKey"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Category  string   `json:"category"`
	Tags      []string `json:"tags" gorm:"type:text[]"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
