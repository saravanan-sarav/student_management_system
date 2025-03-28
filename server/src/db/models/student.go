package models

import "time"

type Student struct {
	ID             uint      `gorm:"primaryKey"`
	Name           string    `gorm:"size:100;not null"`
	EnrollNumber   string    `gorm:"size:10;not null;uniqueIndex"`
	Email          string    `gorm:"size:100;uniqueIndex;not null"`
	Mobile         string    `gorm:"size:15;uniqueIndex"`
	Standard       string    `gorm:"size:20"`
	MotherName     string    `gorm:"size:100"`
	FatherName     string    `gorm:"size:100"`
	Address        string    `gorm:"size:255"`
	DateOfBirth    time.Time `gorm:"not null"`
	DateOfJoining  time.Time `gorm:"not null"`
	ProfilePicture string    `gorm:"size:255"`
	Gender         string    `gorm:"size:10"`
	BloodGroup     string    `gorm:"size:10"`
	Nationality    string    `gorm:"size:50"`
	CreatedAt      time.Time
	ModifiedAt     time.Time `gorm:"default:null"`

	// Relations
	Auth              Auth              `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	AttendanceEntries []AttendanceEntry `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	WorkDoneEntries   []WorkDoneEntry   `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	LeaveEntries      []LeaveEntry      `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	Notifications     []Notification    `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
}
