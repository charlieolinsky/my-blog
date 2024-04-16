package service

import (
	"context"
	"fmt"

	"github.com/charlieolinsky/my-blog/internal/model"
	"github.com/charlieolinsky/my-blog/internal/repo"
)

type commentService struct {
	commentRepo repo.CommentRepository
	userRepo    repo.UserRepository
}

func NewCommentService(commentRepo repo.CommentRepository, userRepo repo.UserRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		userRepo:    userRepo,
	}
}

func (s *commentService) CreateComment(ctx context.Context, newComment model.Comment) error {
	if newComment.Body == "" {
		return fmt.Errorf("comment body cannot be empty")
	}
	return s.commentRepo.CreateComment(ctx, newComment)
}
func (s *commentService) GetCommentByID(ctx context.Context, commentID int) (*model.Comment, error) {
	if commentID <= 0 {
		return nil, fmt.Errorf("invalid comment ID: %d", commentID)
	}
	return s.commentRepo.GetCommentByID(ctx, commentID)
}
func (s *commentService) GetAllComments(ctx context.Context) ([]model.Comment, error) {
	return s.commentRepo.GetAllComments(ctx)
}
func (s *commentService) UpdateComment(ctx context.Context, callerID int, commentID int, updatedComment model.Comment) error {
	comment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return fmt.Errorf("failed to retrieve comment: %w", err)
	}
	if comment.UserID != callerID {
		return fmt.Errorf("user is not authorized to update comment")
	}

	if commentID <= 0 {
		return fmt.Errorf("invalid comment ID: %d", commentID)
	}
	return s.commentRepo.UpdateComment(ctx, commentID, updatedComment)
}
func (s *commentService) DeleteComment(ctx context.Context, callerID int, commentID int) error {
	if commentID <= 0 {
		return fmt.Errorf("invalid comment ID: %d", commentID)
	}

	comment, err := s.commentRepo.GetCommentByID(ctx, commentID)
	if err != nil {
		return fmt.Errorf("failed to retrieve comment: %w", err)
	}
	if comment.UserID != callerID {
		return fmt.Errorf("user is not authorized to update comment")
	}

	return s.commentRepo.DeleteComment(ctx, commentID)
}
