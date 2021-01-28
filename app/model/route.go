package model

type Route struct {
	ID          int64 `gorm:"primaryKey"`
	UserID      string
	OutwardTrip string
	ReturnTrip  string
}
