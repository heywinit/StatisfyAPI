package models

import (
	"gorm.io/gorm"
)

type Followers struct {
	Total int `json:"total"`
}

// User represents a user in the Statisfy application.
type User struct {
	gorm.Model
	DisplayName  string            `gorm:"not null"`
	ID           string            `gorm:"primaryKey"`
	ExternalURLs map[string]string `gorm:"type:json"`
	Followers    Followers         `gorm:"embedded;embeddedPrefix:followers_"`
	Email        string            `gorm:"uniqueIndex;not null"`
	Images       []string          `gorm:"type:json"`
	URI          string            `gorm:"not null"`
}

type PrivateUser struct {
	User
	Country   string `json:"country"`
	Email     string `json:"email"`
	Product   string `json:"product"`
	Birthdate string `json:"birthdate"`
}

// TableName sets the table name for User model in the database.
func (u *User) TableName() string {
	return "users"
}
