package model

import "time"

type ActivityProgram struct {
	ID               int `gorm:"primaryKey"`
	UserID           string
	Datetime         *time.Time
	StartTime        string
	EndTime          string
	PracticeSection  string
	PracticeContents string
	VenueID          int
	OutwardTrip      string
	ReturnTrip       string
	ContactPerson1   bool
	ContactAbstract1 string
	ContactPerson2   bool
	ContactAbstract2 string
	CreateTime       time.Time
	UpdateTime       time.Time
}
