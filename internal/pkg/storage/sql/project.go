package sql

import (
	"errors"

	"github.com/google/uuid"
	"github.com/hse-revizor/projects-service/internal/pkg/models"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PaginationPayload struct {
	Limit int
	Token int64
}

func (s *Storage) CreateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(&model).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrDuplicatedKey):
			return nil, ErrEntityExists
		case errors.Is(err, gorm.ErrForeignKeyViolated):
			return nil, ErrForeignKey
		default:
			return nil, err
		}
	}

	return model, nil
}
func (s *Storage) FindProjectById(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	intent := new(models.Project)
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.
		Model(&models.Project{}).Where("id = ?", id).
		First(intent).
		Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrEntityNotFound
		default:
			return nil, err
		}
	}

	return intent, nil
}

func (s *Storage) UpdateProject(ctx context.Context, model *models.Project) (*models.Project, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Clauses(clause.Returning{}).Model(&model).Where("id = ?", model.Id).Updates(model)
	if result.Error != nil {
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return nil, ErrEntityNotFound
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return nil, ErrEntityExists
		default:
			return nil, result.Error
		}
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}

	return model, nil
}

func (s *Storage) DeleteProject(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)

	deletedIntent := new(models.Project)
	result := tr.Clauses(clause.Returning{}).Where("id = ?", id).Delete(deletedIntent)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrEntityNotFound
		}

		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotFound
	}

	return deletedIntent, nil
}
func (s *Storage) GetProjectsById(ctx context.Context, ids []uuid.UUID) ([]*models.Project, error) {
	res := make([]*models.Project, 0)
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Model(&models.Project{}).Where("id  in ?", ids).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PaginationInput struct {
	Limit int
	Skip  int
}

type GetAllProjectsPayload struct {
	*PaginationInput
}
type GetAllProjectsOutput struct {
	Projects []*models.Project
	Count    int32
}

func (s *Storage) GetAllProjects(ctx context.Context, input GetAllProjectsPayload) (*GetAllProjectsOutput, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var out GetAllProjectsOutput
	var count int64
	err := tr.
		Model(&models.Project{}).
		Count(&count).
		Limit(input.Limit).
		Offset(input.Skip).
		Find(&out.Projects).
		Error
	if err != nil {
		return nil, err
	}
	out.Count = int32(count)
	return &out, nil
}
