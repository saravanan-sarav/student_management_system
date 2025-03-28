package models

import "time"

type Notification struct {
	ID        uint   `gorm:"primaryKey"`
	StudentID uint   `gorm:"not null;index"`
	Message   string `gorm:"size:255;not null"`
	IsRead    bool   `gorm:"default:false"`
	CreatedAt time.Time

	// Foreign Key
	Student Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
}
