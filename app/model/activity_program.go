package model

type ActivityProgramID int

type ActivityProgram struct {
	ID               ActivityProgramID `gorm:"primaryKey"`
	UserID           string
	Datetime         string
	StartTime        string
	EndTime          string
	PracticeSection  string
	PracticeContents string
	VenueID          int
	RouteID          int
	OutwardTrip      string
	ReturnTrip       string
	ContactPerson1   bool
	ContactAbstract1 string
	ContactPerson2   bool
	ContactAbstract2 string
}
