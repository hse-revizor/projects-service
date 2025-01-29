package project

import (
	"context"

	"github.com/google/uuid"
	"github.com/hse-revizor/projects-service/internal/pkg/models"
	"github.com/hse-revizor/projects-service/internal/pkg/storage/sql"
)

type storage interface {
	CreateProject(context.Context, *models.Project) (*models.Project, error)
	FindProjectById(context.Context, uuid.UUID) (*models.Project, error)
	UpdateProject(context.Context, *models.Project) (*models.Project, error)
	DeleteProject(context.Context, uuid.UUID) (*models.Project, error)
	GetAllProjects(ctx context.Context, input sql.GetAllProjectsPayload) (*sql.GetAllProjectsOutput, error)
}
type Service struct {
	storage storage
}

func New(storage storage) *Service {
	return &Service{storage: storage}
}
