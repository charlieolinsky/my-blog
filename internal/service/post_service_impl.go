package service

import (
	"context"
	"fmt"

	"github.com/charlieolinsky/my-blog/internal/model"
	"github.com/charlieolinsky/my-blog/internal/repo"
)

type postService struct {
	postRepo repo.PostRepository
	userRepo repo.UserRepository
}

func NewPostService(postRepo repo.PostRepository, userRepo repo.UserRepository) PostService {
	return &postService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (s *postService) CreatePost(ctx context.Context, newPost model.Post) error {
	user, err := s.userRepo.GetUserByID(ctx, newPost.UserID)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %w", err)
	}
	if user.Role != "admin" {
		return fmt.Errorf("user is not authorized to create a post")
	}
	if newPost.Title == "" {
		return fmt.Errorf("post title cannot be empty")
	}
	if newPost.Body == "" {
		return fmt.Errorf("post body cannot be empty")
	}

	return s.postRepo.CreatePost(ctx, newPost)
}
func (s *postService) GetPostByID(ctx context.Context, postID int) (*model.Post, error) {
	if postID <= 0 {
		return nil, fmt.Errorf("invalid post ID: %d", postID)
	}
	return s.postRepo.GetPostByID(ctx, postID)
}
func (s *postService) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	return s.postRepo.GetAllPosts(ctx)
}
func (s *postService) UpdatePost(ctx context.Context, postID int, updatedPost model.Post) error {
	if postID <= 0 {
		return fmt.Errorf("invalid post ID: %d", postID)
	}
	return s.postRepo.UpdatePost(ctx, postID, updatedPost)
}
func (s *postService) DeletePost(ctx context.Context, postID int) error {
	if postID <= 0 {
		return fmt.Errorf("invalid post ID: %d", postID)
	}
	return s.postRepo.DeletePost(ctx, postID)
}
