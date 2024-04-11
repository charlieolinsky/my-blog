package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/charlieolinsky/my-blog/internal/model"
)

// projectRepository manages project-related database operations.
type projectRepository struct {
	db *sql.DB
}

// NewProjectRepository creates a new project repository.
func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// CreateProject adds a new project to the database.
func (r *projectRepository) CreateProject(ctx context.Context, newProject model.Project) error {
	query := "INSERT INTO projects (user_id, title, body, created_at, updated_at) VALUES (?, ?, ?, ?, NULL)"
	_, err := r.db.ExecContext(ctx, query, newProject.UserId, newProject.Title, newProject.Body, time.Now().UTC())
	return err
}

// GetProject fetches a project by ID, excluding deleted projects.
func (r *projectRepository) GetProject(ctx context.Context, projectID int) (*model.Project, error) {
	var project model.Project
	query := "SELECT project_id, user_id, title, body, created_at, updated_at FROM projects WHERE project_id=? AND deleted_at IS NULL"
	err := r.db.QueryRowContext(ctx, query, projectID).Scan(&project.ProjectId, &project.UserId, &project.Title, &project.Body, &project.CreatedAt, &project.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no project found with ID: %d", projectID)
		}
		return nil, err
	}
	return &project, nil
}

// UpdateProject modifies an existing project, excluding deleted ones.
func (r *projectRepository) UpdateProject(ctx context.Context, projectID int, updatedProject model.Project) error {
	query := "UPDATE projects SET title=?, body=?, updated_at=? WHERE project_id=? AND deleted_at IS NULL"
	res, err := r.db.ExecContext(ctx, query, updatedProject.Title, updatedProject.Body, time.Now().UTC(), projectID)

	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("no project found with ID: %d or it has been deleted", projectID)
	}

	return nil
}

// DeleteProject soft-deletes a project by setting deleted_at to the current time.
func (r *projectRepository) DeleteProject(ctx context.Context, projectID int) error {
	query := "UPDATE projects SET deleted_at=? WHERE project_id=?"
	res, err := r.db.ExecContext(ctx, query, time.Now().UTC(), projectID)

	if err != nil {
		return err
	}

	if rowsAffected, err := res.RowsAffected(); err != nil || rowsAffected == 0 {
		return fmt.Errorf("no project found with ID: %d", projectID)
	}

	return nil
}
