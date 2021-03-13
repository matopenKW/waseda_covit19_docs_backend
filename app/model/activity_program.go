package model

type ActivityProgramSeqNo int

type ActivityProgram struct {
	UserID             UserID               `gorm:"primaryKey"`
	SeqNo              ActivityProgramSeqNo `gorm:"primaryKey"`
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
