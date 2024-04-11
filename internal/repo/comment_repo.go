package repo

import (
	"context"
	"time"
)

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

type CommentRepository interface {
	CreateComment(ctx context.Context, newComment Comment) error
	GetComment(ctx context.Context, CommentID int) (*Comment, error)
	UpdateComment(ctx context.Context, CommentID int, updatedComment Comment) error
	DeleteComment(ctx context.Context, CommentID int) error
}
