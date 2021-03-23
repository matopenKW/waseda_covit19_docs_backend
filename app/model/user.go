package model

type UserID string

type User struct {
	ID             UserID `gorm:"primaryKey"`
	Email          string
	Name           string
	UniversityName string
	FacultyName    string
	StudentNo      string
	CellPhonNo     string
	Ki             int
	PartID         int
}
