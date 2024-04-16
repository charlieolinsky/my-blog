package service

import (
	"context"
	"fmt"

	"github.com/charlieolinsky/my-blog/internal/model"
	"github.com/charlieolinsky/my-blog/internal/repo"
)

type projectService struct {
	projectRepo repo.ProjectRepository
	userRepo    repo.UserRepository
}

func NewProjectService(projectRepo repo.ProjectRepository, userRepo repo.UserRepository) ProjectService {
	return &projectService{
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (s *projectService) CreateProject(ctx context.Context, newProject model.Project) error {
	user, err := s.userRepo.GetUserByID(ctx, newProject.UserId)
	if err != nil {
		return fmt.Errorf("failed to retrieve user: %w", err)
	}
	if user.Role != "admin" {
		return fmt.Errorf("insufficient permissions to create a project")
	}

	if newProject.Title == "" {
		return fmt.Errorf("title is required")
	}

	return s.projectRepo.CreateProject(ctx, newProject)
}

func (s *projectService) GetProjectByID(ctx context.Context, projectID int) (*model.Project, error) {
	if projectID <= 0 {
		return nil, fmt.Errorf("invalid project ID: %d", projectID)
	}
	return s.projectRepo.GetProject(ctx, projectID)
}

func (s *projectService) GetAllProjects(ctx context.Context) ([]model.Project, error) {
	return s.projectRepo.GetAllProjects(ctx)
}

func (s *projectService) UpdateProject(ctx context.Context, projectID int, updatedProject model.Project) error {
	if projectID <= 0 {
		return fmt.Errorf("invalid project ID: %d", projectID)
	}
	if updatedProject.Title == "" {
		return fmt.Errorf("title is required")
	}

	return s.projectRepo.UpdateProject(ctx, projectID, updatedProject)
}

func (s *projectService) DeleteProject(ctx context.Context, projectID int) error {
	if projectID <= 0 {
		return fmt.Errorf("invalid project ID: %d", projectID)
	}
	return s.projectRepo.DeleteProject(ctx, projectID)
}
