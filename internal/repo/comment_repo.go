package repo

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, newComment model.Comment) error
	GetCommentByID(ctx context.Context, CommentID int) (*model.Comment, error)
	GetAllComments(ctx context.Context) ([]model.Comment, error)
	UpdateComment(ctx context.Context, CommentID int, updatedComment model.Comment) error
	DeleteComment(ctx context.Context, CommentID int) error
}
