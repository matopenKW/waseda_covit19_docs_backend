package model

import "time"

type ActivityProgram struct {
	ID               int `gorm:"primaryKey"`
	UserId           string
	Datetime         string
	StartTime        string
	EndTime          string
	PracticeSection  string
	PracticeContents string
	VenueID          int
	Outbound         string
	ReturnTrip       string
	ContactPerson1   string
	ContactAbstract1 string
	ContactPerson2   string
	ContactAbstract2 string
	CreateTime       time.Time
	UpdateTime       time.Time
}
