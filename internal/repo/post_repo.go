package repo

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, newPost model.Post) error
	GetPost(ctx context.Context, PostID int) (*model.Post, error)
	UpdatePost(ctx context.Context, PostID int, updatedPost model.Post) error
	DeletePost(ctx context.Context, PostID int) error
}
