package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

// CreatePost inserts a new post into the database.
func (r *postRepository) CreatePost(ctx context.Context, newPost model.Post) error {
	query := "INSERT INTO posts (user_id, project_id, title, body, likes, created_at, updated_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?, NULL, NULL)"
	_, err := r.db.ExecContext(ctx, query, newPost.UserID, newPost.ProjectID, newPost.Title, newPost.Body, newPost.Likes, time.Now().UTC())
	return err
}

// GetPost retrieves a post by its ID, excluding soft-deleted posts.
func (r *postRepository) GetPostByID(ctx context.Context, PostID int) (*model.Post, error) {
	var post model.Post
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

// GetAllPosts retrieves all posts, excluding soft-deleted posts.
func (r *postRepository) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	var posts []model.Post
	query := "SELECT post_id, user_id, project_id, title, body, likes, created_at, updated_at FROM posts WHERE deleted_at IS NULL"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve posts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		err = rows.Scan(&post.PostID, &post.UserID, &post.ProjectID, &post.Title, &post.Body, &post.Likes, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error reading posts: %w", err)
	}
	return posts, nil
}

// UpdatePost modifies an existing post's details.
func (r *postRepository) UpdatePost(ctx context.Context, PostID int, updatedPost model.Post) error {
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
