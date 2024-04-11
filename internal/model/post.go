package model

import "time"

type Post struct {
	PostID    int
	UserID    int
	ProjectID int
	Title     string
	Body      string
	Likes     int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
