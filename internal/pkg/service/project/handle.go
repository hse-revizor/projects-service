package project

import (
	"context"
	"errors"
	"fmt"

	"github.com/4kayDev/logger/log"
	"github.com/google/uuid"
	"github.com/hse-revizor/projects-service/internal/pkg/models"
	"github.com/hse-revizor/projects-service/internal/pkg/storage/sql"
	"github.com/hse-revizor/projects-service/internal/utils/json"
)

type CreateProject struct {
	Name          string
	RepositoryURL string
}

// @throws: ErrProjectNotFound, ErrProjectExists
func (s *Service) CreateProject(ctx context.Context, input *CreateProject) (*models.Project, error) {
	created, err := s.storage.CreateProject(ctx, &models.Project{
		Sources:      []string{input.RepositoryURL},
		Name:         input.Name,
		ProjectGroup: models.ProjectGroup{},
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityExists):
			return nil, ErrProjectExists
		case errors.Is(err, sql.ErrForeignKey):
			return nil, ErrForeignKeyError
		default:
			return nil, err
		}
	}
	log.Debug(fmt.Sprintf("Created Project: %s", json.ToColorJson(created)))
	return created, nil
}

// @throws: ErrProjectNotFound, ErrProjectExists
func (s *Service) UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	project, err := s.storage.UpdateProject(ctx, project)
	if err != nil {
		return nil, err
	}
	log.Debug(fmt.Sprintf("Updated Project: %s", json.ToColorJson(project)))
	return project, nil
}

// @throws: ErrProjectNotFound
func (s *Service) DeleteProject(ctx context.Context, projectId uuid.UUID) (*models.Project, error) {
	model, err := s.storage.DeleteProject(ctx, projectId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrProjectNotFound
		default:
			return nil, err
		}
	}

	log.Debug(fmt.Sprintf("Deleted Project Account: %s", json.ToColorJson(model)))

	return model, nil
}

func (s *Service) GetProjectById(ctx context.Context, projectId uuid.UUID) (*models.Project, error) {
	project, err := s.storage.FindProjectById(ctx, projectId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrProjectNotFound
		default:
			return nil, err
		}
	}

	log.Debug(fmt.Sprintf("Found Project: %s", json.ToColorJson(project)))

	return project, nil

}

type GetAllProjectsOutput struct {
	Projects []*models.Project
	Count    int32
}

func (s *Service) GetAllProjects(ctx context.Context, skip, limit int) (*GetAllProjectsOutput, error) {
	projects, err := s.storage.GetAllProjects(ctx, sql.GetAllProjectsPayload{
		PaginationInput: &sql.PaginationInput{
			Limit: limit,
			Skip:  skip,
		},
	})
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrEntityNotFound):
			return nil, ErrProjectNotFound
		default:
			return nil, err
		}
	}
	log.Debug(fmt.Sprintf("Found Project: %s", json.ToColorJson(projects)))

	return &GetAllProjectsOutput{
		Projects: projects.Projects,
		Count:    projects.Count,
	}, nil

}
