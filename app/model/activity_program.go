package model

type ActivityProgram struct {
	ID               int `gorm:"primaryKey"`
	UserId           string
	Datetime         string
	StartTime        string
	EndTime          string
	PracticeSection  string
	PracticeContents string
	VenueID          int
	RouteID          int
	ContactPerson1   string
	ContactAbstract1 string
	ContactPerson2   string
	ContactAbstract2 string
	CreateTime       string
	UpdateTime       string
}