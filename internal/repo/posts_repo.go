package repo

import (
	"context"
	"time"
)

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

type PostRepository interface {
	CreatePost(ctx context.Context, newPost Post) error
	GetPost(ctx context.Context, PostID int) (*Post, error)
	UpdatePost(ctx context.Context, PostID int, updatedPost Post) error
	DeletePost(ctx context.Context, PostID int) error
}
