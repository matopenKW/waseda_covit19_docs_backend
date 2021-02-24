package model

type LastUpload struct {
	WeekID  int `gorm:"primaryKey"`
	DriveID string
}
