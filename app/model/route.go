package model

type RouteID int64

type Route struct {
	ID          RouteID `gorm:"primaryKey"`
	UserID      string
	Name        string
	OutwardTrip string
	ReturnTrip  string
}
