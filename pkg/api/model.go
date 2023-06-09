package api

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null;" json:"updated_at"`
}

type URL struct {
	Model
	LongURL  string `gorm:"column:long_url;" json:"long_url"`
	ShortURL string `gorm:"column:short_url;" json:"short_url"`
}
