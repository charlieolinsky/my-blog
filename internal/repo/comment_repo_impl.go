package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}

// CreateComment inserts a new comment into the database.
func (r *commentRepository) CreateComment(ctx context.Context, newComment model.Comment) error {
	query := "INSERT INTO comments (user_id, post_id, body, likes, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, NULL, NULL)"
	_, err := r.db.ExecContext(ctx, query, newComment.UserID, newComment.PostID, newComment.Body, newComment.Likes, time.Now().UTC())
	return err
}

// GetComment retrieves a comment by its ID, including soft-deleted comments.
func (r *commentRepository) GetCommentByID(ctx context.Context, CommentID int) (*model.Comment, error) {
	var comment model.Comment
	query := "SELECT comment_id, user_id, post_id, body, likes, created_at, updated_at FROM comments WHERE comment_id=?"
	err := r.db.QueryRowContext(ctx, query, CommentID).Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.Body, &comment.Likes, &comment.CreatedAt, &comment.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no comment found with ID: %d", CommentID)
		}
		return nil, err
	}
	return &comment, nil
}

// GetAllComments retrieves all comments, including soft-deleted comments.
func (r *commentRepository) GetAllComments(ctx context.Context) ([]model.Comment, error) {
	var comments []model.Comment
	query := "SELECT comment_id, user_id, post_id, body, likes, created_at, updated_at FROM comments"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve comments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment
		err = rows.Scan(&comment.CommentID, &comment.UserID, &comment.PostID, &comment.Body, &comment.Likes, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to retrieve comments: %w", err)
	}

	return comments, nil
}

// UpdateComment modifies an existing comment's details. Deleted comments may not be updated.
func (r *commentRepository) UpdateComment(ctx context.Context, CommentID int, updatedComment model.Comment) error {
	query := "UPDATE comments SET body=?, likes=?, updated_at=? WHERE comment_id=? AND deleted_at IS NULL"
	_, err := r.db.ExecContext(ctx, query, updatedComment.Body, updatedComment.Likes, time.Now().UTC(), CommentID)

	if err != nil {
		return err
	}
	return nil
}

// DeleteComment soft-deletes a comment by setting its deleted_at timestamp.
func (r *commentRepository) DeleteComment(ctx context.Context, commentID int) error {
	query := "UPDATE comments SET deleted_at=? WHERE comment_id=?"
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), commentID)

	if err != nil {
		return err
	}
	return nil
}
