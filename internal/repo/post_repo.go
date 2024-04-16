package repo

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, newPost model.Post) error
	GetPostByID(ctx context.Context, PostID int) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
	UpdatePost(ctx context.Context, PostID int, updatedPost model.Post) error
	DeletePost(ctx context.Context, PostID int) error
}
