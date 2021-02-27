package model

type UserID string

type User struct {
	ID    UserID `gorm:"primaryKey"`
	Email string
}
