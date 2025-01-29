package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	Id           uuid.UUID    `gorm:"primaryKey;column:id"`
	Name         string       `gorm:"column:name"`
	Sources      Sources      `gorm:"column:sources"`
	ProjectGroup ProjectGroup `gorm:"many2many:project_groups_projects"`

	CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type ProjectGroup struct {
	Id          uuid.UUID `gorm:"primaryKey;column:id"`
	Files       Sources   `gorm:"column:files"`
	Projects    []Project `gorm:"many2many:project_groups_projects"`
	WorkspaceID uuid.UUID `gorm:"column:workspace_id"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
type Sources []string

func (b *Sources) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	return json.Unmarshal(bytes, b)
}

func (b Sources) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (g *Project) BeforeCreate(tx *gorm.DB) error {
	if g.Id == uuid.Nil {
		g.Id = uuid.New()
	}
	return nil
}
