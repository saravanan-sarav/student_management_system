package models

import "time"

type AttendanceEntry struct {
	ID        uint      `gorm:"primaryKey"`
	StudentID uint      `gorm:"not null;index"`
	Date      time.Time `gorm:"not null"`
	CheckIn   time.Time `gorm:"default:null"`
	CheckOut  time.Time `gorm:"default:null"`
	Status    string    `gorm:"size:20;not null;default:'Present'"`
	CreatedAt time.Time

	// Foreign Key
	Student Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
}
