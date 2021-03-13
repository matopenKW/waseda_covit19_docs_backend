package model

type RouteSeqNo int64

type Route struct {
	UserID      string     `gorm:"primaryKey"`
	SeqNo       RouteSeqNo `gorm:"primaryKey"`
	Name        string
	OutwardTrip string
	ReturnTrip  string
}
