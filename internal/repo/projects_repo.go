package repo

import (
	"context"
	"time"
)

type Project struct {
	ProjectId int
	UserId    int
	Title     string
	Body      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type ProjectRepository interface {
	CreateProject(ctx context.Context, newProject Project) error
	GetProject(ctx context.Context, projectID int) (*Project, error)
	UpdateProject(ctx context.Context, projectID int, updatedProject Project) error
	DeleteProject(ctx context.Context, projectID int) error
}
