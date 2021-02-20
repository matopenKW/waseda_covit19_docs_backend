package model

type ActivityProgramID int

type ActivityProgram struct {
	ID                 ActivityProgramID `gorm:"primaryKey"`
	UserID             string
	Datetime           string
	StartTime          string
	EndTime            string
	PracticeSectionID  uint
	PracticeContentsID uint
	OutwardTrip        string
	ReturnTrip         string
	ContactPerson1     int
	ContactAbstract1   string
	ContactPerson2     int
	ContactAbstract2   string
}
