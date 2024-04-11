package model

import "time"

type Comment struct {
	CommentID int
	UserID    int
	PostID    int
	ProjectID int
	Body      string
	Likes     int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
