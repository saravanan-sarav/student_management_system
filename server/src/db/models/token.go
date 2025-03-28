package models

import "time"

type Token struct {
	ID           uint   `gorm:"primaryKey"`
	AccessToken  string `gorm:"not null"`
	RefreshToken string `gorm:"not null"`
	CreatedAt    time.Time
}
