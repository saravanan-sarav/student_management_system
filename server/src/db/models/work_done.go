package models

import "time"

type WorkDoneEntry struct {
	ID          uint      `gorm:"primaryKey"`
	StudentID   uint      `gorm:"not null;index"`
	Date        time.Time `gorm:"not null"`
	WorkDone    string    `gorm:"size:255"`
	SubmittedAt time.Time `gorm:"default:null"`
	CreatedAt   time.Time

	// Foreign Key
	Student Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
}
