package model

import (
	"time"
)

type User struct {
	UserID            int
	Role              string
	Email             string
	Password          string
	FirstName         string
	LastName          string
	ProfilePictureUrl string
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
}
