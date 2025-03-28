package models

import "time"

type LeaveEntry struct {
	ID         uint      `gorm:"primaryKey"`
	StudentID  uint      `gorm:"not null;index"`
	Date       time.Time `gorm:"not null"`
	Reason     string    `gorm:"size:255"`
	ApprovedBy string    `gorm:"size:100"`
	CreatedAt  time.Time

	// Foreign Key
	Student Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
}
