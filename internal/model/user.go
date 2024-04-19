package model

import (
	"time"
)

type User struct {
	UserID            int
	Role              string `json:"role"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	ProfilePictureUrl string `json:"profilePictureUrl"`
	CreatedAt         time.Time
	UpdatedAt         *time.Time
	DeletedAt         *time.Time
}
