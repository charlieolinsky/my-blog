package service

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type PostService interface {
	CreatePost(ctx context.Context, newPost model.Post) error
	GetPostByID(ctx context.Context, postID int) (*model.Post, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
	UpdatePost(ctx context.Context, postID int, updatedPost model.Post) error
	DeletePost(ctx context.Context, postID int) error
}
