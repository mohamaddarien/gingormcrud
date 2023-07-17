package entities

import (
	"time"
)

type Album struct {
    ID     uint       `gorm:"primaryKey;not null" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt time.Time `json:"deleted_at"`
    Title  string       `json:"title"`
    Artist string       `json:"artist"`
    Price  float64      `json:"price"`
}