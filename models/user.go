package models

// User struct that contains information about users
// Sets column definitions for GORM migration to database
type User struct {
	UserID   int `gorm:"primary_key"`
	Username string
	Password []byte
	IsAdmin  bool
}
