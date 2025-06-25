package model

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:50;" json:"username"`
	Email    string `gorm:"uniqueIndex;not null;size:255;" json:"email"`
	Password string `gorm:"not null;" json:"password"`
	Names    string `json:"names"`
}
