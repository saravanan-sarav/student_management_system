package models

import "time"

type Auth struct {
	ID                  uint      `gorm:"primaryKey"`
	StudentID           uint      `gorm:"uniqueIndex;not null"`
	Password            string    `gorm:"not null"`
	LastLogin           time.Time `gorm:"default:null"`
	FailedLoginAttempts int       `gorm:"default:0"`
	CreatedAt           time.Time
	ModifiedAt          time.Time
}
