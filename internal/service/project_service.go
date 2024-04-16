package service

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type ProjectService interface {
	CreateProject(ctx context.Context, newProject model.Project) error
	GetProjectByID(ctx context.Context, projectID int) (*model.Project, error)
	GetAllProjects(ctx context.Context) ([]model.Project, error)
	UpdateProject(ctx context.Context, projectID int, updatedProject model.Project) error
	DeleteProject(ctx context.Context, projectID int) error
}
