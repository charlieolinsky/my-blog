package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

// CreatePost inserts a new post into the database.
func (r *postRepository) CreatePost(ctx context.Context, newPost Post) error {
	query := "INSERT INTO posts (user_id, project_id, title, body, likes, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, NULL, NULL)"
	_, err := r.db.ExecContext(ctx, query, newPost.UserID, newPost.ProjectID, newPost.Title, newPost.Body, newPost.Likes, time.Now().UTC())
	return err
}

// GetPost retrieves a post by its ID, excluding soft-deleted posts.
func (r *postRepository) GetPost(ctx context.Context, PostID int) (*Post, error) {
	var post Post
	query := "SELECT post_id, user_id, project_id, title, body, likes, created_at, updated_at FROM posts WHERE post_id=? AND deleted_at IS NULL"
	err := r.db.QueryRowContext(ctx, query, PostID).Scan(&post.PostID, &post.UserID, &post.ProjectID, &post.Title, &post.Body, &post.Likes, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no post found with ID: %d", PostID)
		}
		return nil, err
	}
	return &post, nil
}

// UpdatePost modifies an existing post's details.
func (r *postRepository) UpdatePost(ctx context.Context, PostID int, updatedPost Post) error {
	query := "UPDATE posts SET title=?, body=?, likes=?, updated_at=? WHERE post_id=? AND deleted_at IS NULL"
	_, err := r.db.ExecContext(ctx, query, updatedPost.Title, updatedPost.Body, updatedPost.Likes, time.Now().UTC(), PostID)

	if err != nil {
		return err
	}
	return nil
}

// DeletePost soft-deletes a post by setting its deleted_at timestamp.
func (r *postRepository) DeletePost(ctx context.Context, PostID int) error {
	query := "UPDATE posts SET deleted_at=? WHERE post_id=?"
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC(), PostID)

	if err != nil {
		return err
	}
	return nil
}
