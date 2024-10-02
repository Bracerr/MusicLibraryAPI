package domain

import (
	"time"
)

type Song struct {
	ID          uint      `gorm:"primaryKey"`
	GroupName   string    `gorm:"not null"`
	Song        string    `gorm:"not null"`
	ReleaseDate time.Time `gorm:""`
	Text        string    `gorm:"type:text"`
	Link        string    `gorm:"type:varchar(255)"`
}
