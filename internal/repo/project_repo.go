package repo

import (
	"context"

	"github.com/charlieolinsky/my-blog/internal/model"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, newProject model.Project) error
	GetProject(ctx context.Context, projectID int) (*model.Project, error)
	UpdateProject(ctx context.Context, projectID int, updatedProject model.Project) error
	DeleteProject(ctx context.Context, projectID int) error
}
