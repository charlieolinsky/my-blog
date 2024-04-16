package service

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type CommentService interface {
	CreateComment(ctx context.Context, newComment model.Comment) error
	GetCommentByID(ctx context.Context, commentID int) (*model.Comment, error)
	GetAllComments(ctx context.Context) ([]model.Comment, error)
	UpdateComment(ctx context.Context, callerID int, commentID int, updatedComment model.Comment) error
	DeleteComment(ctx context.Context, callerID int, commentID int) error
}
