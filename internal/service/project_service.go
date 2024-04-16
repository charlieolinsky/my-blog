package service

import (
	"github.com/charlieolinsky/my-blog/internal/model"
)

type ProjectService interface {
	CreateProject(project model.Project) error
	GetProjectByID(id int) (model.Project, error)
	GetAllProjects() ([]model.Project, error)
	UpdateProject(project model.Project) error
	DeleteProject(id int) error
}
